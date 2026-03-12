package jp2

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

var sig = [12]byte{
	0x00, 0x00, 0x00, 0x0c,
	0x6a, 0x50, 0x20, 0x20,
	0x0d, 0x0a, 0x87, 0x0a,
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var s [12]byte

	_, err := io.ReadFull(r, s[:])
	if err != nil {
		return nil, err
	}

	if s != sig {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "JPEG 2000 image",
		Type: "JP2",
	}, nil
}
