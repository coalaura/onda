package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/coalaura/plain"
	"github.com/coalaura/wtf/types"
)

var (
	Version = "dev"
	log     = plain.New()
)

const MaxReadSize = 128 * 1024

const bashCompletion = `_wtf() {
    local cur prev
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"
    case "$prev" in
        --completion)
            COMPREPLY=($(compgen -W "bash" -- "$cur"))
            return
            ;;
    esac
    if [[ "$cur" == -* ]]; then
        COMPREPLY=($(compgen -W "-l -p -t -d -v -h --help --version --completion" -- "$cur"))
    else
        compopt -o filenames
        COMPREPLY=($(compgen -f -- "$cur"))
    fi
}
complete -F _wtf wtf
`

type kindGroup struct {
	kind  string
	types []string
}

func main() {
	var (
		targetFile   string
		porcelain    bool
		timing       bool
		deepAnalysis bool
	)

	for _, arg := range os.Args[1:] {
		switch arg {
		case "-h", "--help":
			log.Println("wtf - hardware-accelerated file sniffer")
			log.Println("\nUsage: wtf [flags] <file>")
			log.Println("\nFlags:")
			log.Println("  -l, --list       List all supported formats and sub-formats")
			log.Println("  -p, --porcelain  Print easily parseable output (tab-separated: Kind\\tType\\tConfidence)")
			log.Println("  -t, --time       Print time taken (read / sniff; disabled by -p)")
			log.Println("  -d, --deep       Enable deep analysis (entropy, byte distribution) for unknown files")
			log.Println("      --completion Print shell completion script (bash)")
			log.Println("  -v, --version    Print version information")
			log.Println("  -h, --help       Print this help message")

			os.Exit(0)
		case "-v", "--version":
			log.Printf("wtf %s\n", Version)

			os.Exit(0)
		case "-l", "--list":
			listFormats()

			os.Exit(0)
		case "--completion":
			if len(os.Args) < 3 || os.Args[2] != "bash" {
				log.Errorln("Usage: wtf --completion bash")

				os.Exit(1)
			}

			fmt.Print(bashCompletion)

			os.Exit(0)
		case "-p", "--porcelain":
			porcelain = true
		case "-t", "--time":
			timing = true
		case "-d", "--deep":
			deepAnalysis = true
		default:
			if len(arg) == 0 || arg[0] != '-' {
				if targetFile != "" {
					log.Errorln("Usage: wtf [flags] <file>")

					os.Exit(1)
				}

				targetFile = arg
			} else {
				log.Errorln("Unknown argument: ", arg)

				os.Exit(1)
			}
		}
	}

	processFile(targetFile, porcelain, timing, deepAnalysis)
}

func processFile(path string, porcelain bool, timing bool, deepAnalysis bool) {
	name := filepath.Base(path)

	t0 := time.Now()

	info, err := os.Lstat(path)
	if err != nil {
		log.Errorln(err)

		os.Exit(1)
	}

	meta := detectPath(name, info)
	if meta != nil {
		d0 := time.Since(t0)

		printMeta(meta, formatTimings(d0, 0, timing), porcelain)

		return
	}

	readSize := int(min(info.Size(), int64(MaxReadSize)))

	if readSize == 0 {
		log.Errorln("empty file")

		os.Exit(2)
	}

	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		log.Errorln(err)

		os.Exit(1)
	}

	defer file.Close()

	buf := make([]byte, readSize)

	n, err := io.ReadFull(file, buf)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		log.Errorln(err)

		os.Exit(1)
	}

	d0 := time.Since(t0)
	t1 := time.Now()

	meta, err = types.DetectReaderAt(name, buf[:n], file, info.Size())
	if err != nil && !errors.Is(err, types.ErrUnknownFormat) {
		if porcelain {
			log.Println("Unknown\t")
		} else {
			log.Errorln(err)
		}

		os.Exit(1)
	}

	if deepAnalysis && errors.Is(err, types.ErrUnknownFormat) {
		deep := types.DetectDeepAnalysis(types.Buffer(buf[:n]))
		if deep != nil {
			deep.File = name
			meta = deep
		}
	}

	d1 := time.Since(t1)

	if meta == nil {
		if porcelain {
			log.Println("Unknown\t")
		} else {
			log.Errorln(types.ErrUnknownFormat)
		}

		os.Exit(1)
	}

	printMeta(meta, formatTimings(d0, d1, timing), porcelain)
}

func detectPath(name string, info os.FileInfo) *types.Metadata {
	mode := info.Mode()

	if mode&os.ModeSymlink != 0 {
		return &types.Metadata{
			File: name,
			Kind: types.KindFilesystemEntry,
			Type: types.TypeSymbolicLink,
		}
	}

	if mode.IsDir() {
		return &types.Metadata{
			File: name,
			Kind: types.KindFilesystemEntry,
			Type: types.TypeDirectory,
		}
	}

	if mode&os.ModeNamedPipe != 0 {
		return &types.Metadata{
			File: name,
			Kind: types.KindFilesystemEntry,
			Type: types.TypeNamedPipe,
		}
	}

	if mode&os.ModeSocket != 0 {
		return &types.Metadata{
			File: name,
			Kind: types.KindFilesystemEntry,
			Type: types.TypeSocket,
		}
	}

	if mode&os.ModeDevice != 0 {
		if mode&os.ModeCharDevice != 0 {
			return &types.Metadata{
				File: name,
				Kind: types.KindFilesystemEntry,
				Type: types.TypeCharacterDevice,
			}
		}

		return &types.Metadata{
			File: name,
			Kind: types.KindFilesystemEntry,
			Type: types.TypeBlockDevice,
		}
	}

	if !mode.IsRegular() {
		return &types.Metadata{
			File: name,
			Kind: types.KindFilesystemEntry,
			Type: types.TypeSpecial,
		}
	}

	return nil
}

func printMeta(meta *types.Metadata, timings string, porcelain bool) {
	if porcelain {
		log.Printf("%s\t%s\t%s\n", meta.Kind.String(), meta.Type.String(), meta.Confidence.String())
	} else {
		log.Println(meta.Format(timings))
	}
}

func formatTimings(d0, d1 time.Duration, enabled bool) string {
	if !enabled {
		return ""
	}

	return fmt.Sprintf("%s / %s", formatTime(d0), formatTime(d1))
}

func formatTime(d time.Duration) string {
	if d < 0 {
		d = 0
	}

	if d == 0 {
		return "0s"
	}

	type unit struct {
		value  time.Duration
		suffix string
	}

	units := []unit{
		{time.Hour, "h"},
		{time.Minute, "m"},
		{time.Second, "s"},
		{time.Millisecond, "ms"},
		{time.Microsecond, "µs"},
	}

	parts := make([]string, 0, len(units))

	for _, unit := range units {
		if d < unit.value {
			continue
		}

		count := d / unit.value
		d -= count * unit.value

		parts = append(parts, fmt.Sprintf("%d%s", count, unit.suffix))
	}

	if len(parts) == 0 {
		return fmt.Sprintf("%dns", d/time.Nanosecond)
	}

	return strings.Join(parts, " ")
}

func listFormats() {
	formats := types.ListFormats()

	if len(formats) == 0 {
		return
	}

	var groups []kindGroup

	current := &kindGroup{kind: formats[0].Kind}

	for _, f := range formats[1:] {
		if f.Kind != current.kind {
			groups = append(groups, *current)

			current = &kindGroup{kind: f.Kind}
		}

		if f.Type != "" {
			current.types = append(current.types, f.Type)
		}
	}

	groups = append(groups, *current)

	for _, g := range groups {
		log.Printf("%s\n", g.kind)

		for i, t := range g.types {
			if i == len(g.types)-1 {
				log.Printf("  └─ %s\n", t)
			} else {
				log.Printf("  ├─ %s\n", t)
			}
		}
	}
}
