package gzip

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var hdr [3]byte

	_, err := io.ReadFull(r, hdr[:])
	if err != nil {
		return nil, err
	}

	if hdr[0] != 0x1f || hdr[1] != 0x8b {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "GZIP compressed data",
	}, nil
}
