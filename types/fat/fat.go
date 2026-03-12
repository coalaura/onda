package fat

import (
	"bytes"
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var hdr [90]byte

	_, err := io.ReadFull(r, hdr[:])
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(hdr[54:59], []byte("FAT12")) && !bytes.Equal(hdr[54:59], []byte("FAT16")) && !bytes.Equal(hdr[82:87], []byte("FAT32")) {
		return nil, types.ErrUnknownFormat
	}

	var typ string

	switch {
	case bytes.Equal(hdr[54:59], []byte("FAT12")):
		typ = "FAT12"
	case bytes.Equal(hdr[54:59], []byte("FAT16")):
		typ = "FAT16"
	case bytes.Equal(hdr[82:87], []byte("FAT32")):
		typ = "FAT32"
	}

	return &types.Metadata{
		Name: "FAT filesystem",
		Type: typ,
	}, nil
}
