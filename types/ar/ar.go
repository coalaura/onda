package ar

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

var sig = [8]byte{'!', '<', 'a', 'r', 'c', 'h', '>', '\n'}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var s [8]byte

	_, err := io.ReadFull(r, s[:])
	if err != nil {
		return nil, err
	}

	if s != sig {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "Unix archive",
	}, nil
}
