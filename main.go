package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/coalaura/onda/types"
	"github.com/coalaura/plain"
)

var (
	Version = "dev"
	log     = plain.New()
)

func main() {
	var (
		targetFile string
		porcelain  bool
	)

	for _, arg := range os.Args[1:] {
		switch arg {
		case "-h", "--help":
			fmt.Println("onda - hardware-accelerated file sniffer")
			fmt.Println("\nUsage: onda [flags] <file>")
			fmt.Println("\nFlags:")
			fmt.Println("  -p, --porcelain  Print easily parseable output (tab-separated: Kind\\tType)")
			fmt.Println("  -v, --version    Print version information")
			fmt.Println("  -h, --help       Print this help message")

			os.Exit(0)
		case "-v", "--version":
			fmt.Printf("onda %s\n", Version)

			os.Exit(0)
		case "-p", "--porcelain":
			porcelain = true
		default:
			if targetFile == "" && (len(arg) == 0 || arg[0] != '-') {
				targetFile = arg
			} else {
				log.Errorln("Unknown argument or multiple files specified: ", arg)

				os.Exit(1)
			}
		}
	}

	if targetFile == "" {
		log.Errorln("Usage: onda [flags] <file>")

		os.Exit(1)
	}

	path, err := filepath.Abs(targetFile)
	if err != nil {
		log.Errorln(err)

		os.Exit(1)
	}

	name := filepath.Base(path)

	info, err := os.Lstat(path)
	if err != nil {
		log.Errorln(err)

		os.Exit(1)
	}

	meta := detectPath(name, info)
	if meta != nil {
		printMeta(meta, porcelain)

		return
	}

	if info.Size() == 0 {
		log.Errorln("empty file")

		os.Exit(2)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Errorln(err)

		os.Exit(1)
	}

	defer file.Close()

	var buf [65536]byte

	n, err := file.Read(buf[:])
	if err != nil && err != io.EOF {
		log.Errorln(err)

		os.Exit(1)
	}

	meta, err = types.Detect(name, buf[:n])
	if err != nil {
		if porcelain {
			fmt.Println("Unknown\t")
		} else {
			log.Errorln(err)
		}

		os.Exit(1)
	}

	printMeta(meta, porcelain)
}

func printMeta(meta *types.Metadata, porcelain bool) {
	if porcelain {
		fmt.Printf("%s\t%s\n", meta.Kind.String(), meta.Type.String())
	} else {
		log.Println(meta.Format())
	}
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
