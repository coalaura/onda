package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"

	"github.com/coalaura/onda/types"
	_ "github.com/coalaura/onda/types/all"
	"github.com/coalaura/plain"
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

	chunk := make([]byte, 4096)

	n, err := io.ReadFull(file, chunk)
	if err != nil {
		log.Errorln(err)

		os.Exit(1)
	}

	rd := bytes.NewReader(chunk[:n])

	meta, err := types.Detect(name, rd)
	if err != nil {
		log.Errorln(err)

		os.Exit(1)
	}

	log.Println(meta.Format())
}
