package bzip2

import (
	"fmt"
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var hdr [4]byte

	_, err := io.ReadFull(r, hdr[:])
	if err != nil {
		return nil, err
	}

	if hdr[0] != 'B' || hdr[1] != 'Z' || hdr[2] != 'h' {
		return nil, types.ErrUnknownFormat
	}

	level := hdr[3]

	if level < '1' || level > '9' {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "bzip2 archive",
		Info: fmt.Sprintf("block size %ck", level),
	}, nil
}
