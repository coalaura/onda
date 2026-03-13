package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/coalaura/onda/types"
	"github.com/coalaura/plain"

	_ "github.com/coalaura/onda/types/all"
)

var log = plain.New()

func main() {
	if len(os.Args) < 2 {
		log.Errorln("Usage: onda <file>")

		os.Exit(1)
	}

	path, err := filepath.Abs(os.Args[1])
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
		log.Println(meta.Format())

		return
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
		log.Errorln(err)

		os.Exit(1)
	}

	log.Println(meta.Format())
}

func detectPath(name string, info os.FileInfo) *types.Metadata {
	mode := info.Mode()

	if mode&os.ModeSymlink != 0 {
		return &types.Metadata{
			File: name,
			Name: "Filesystem Entry",
			Type: "Symbolic Link",
		}
	}

	if mode.IsDir() {
		return &types.Metadata{
			File: name,
			Name: "Filesystem Entry",
			Type: "Directory",
		}
	}

	if mode&os.ModeNamedPipe != 0 {
		return &types.Metadata{
			File: name,
			Name: "Filesystem Entry",
			Type: "Named Pipe",
		}
	}

	if mode&os.ModeSocket != 0 {
		return &types.Metadata{
			File: name,
			Name: "Filesystem Entry",
			Type: "Socket",
		}
	}

	if mode&os.ModeDevice != 0 {
		if mode&os.ModeCharDevice != 0 {
			return &types.Metadata{
				File: name,
				Name: "Filesystem Entry",
				Type: "Character Device",
			}
		}

		return &types.Metadata{
			File: name,
			Name: "Filesystem Entry",
			Type: "Block Device",
		}
	}

	if !mode.IsRegular() {
		return &types.Metadata{
			File: name,
			Name: "Filesystem Entry",
			Type: "Special",
		}
	}

	return nil
}
