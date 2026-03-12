package lz4

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

var sig = [4]byte{0x04, 0x22, 0x4d, 0x18}

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
		Name: "LZ4 compressed data",
	}, nil
}
