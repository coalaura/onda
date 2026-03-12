package zstd

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

var sig = [4]byte{0x28, 0xb5, 0x2f, 0xfd}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var s [4]byte

	_, err := io.ReadFull(r, s[:])
	if err != nil {
		return nil, err
	}

	if s != sig {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "Zstandard compressed data",
	}, nil
}
