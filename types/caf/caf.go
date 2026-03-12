package caf

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var sig [4]byte

	_, err := io.ReadFull(r, sig[:])
	if err != nil {
		return nil, err
	}

	if sig != [4]byte{'c', 'a', 'f', 'f'} {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "CoreAudio container",
	}, nil
}
