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

	file, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0)
	if err != nil {
		log.Errorln(err)

		os.Exit(1)
	}

	defer file.Close()

	var buf [4096]byte

	n, err := file.Read(buf[:])
	if err != nil && err != io.EOF {
		log.Errorln(err)

		os.Exit(1)
	}

	meta, err := types.Detect(name, buf[:n])
	if err != nil {
		log.Errorln(err)

		os.Exit(1)
	}

	log.Println(meta.Format())
}
