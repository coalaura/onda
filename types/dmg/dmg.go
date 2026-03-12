package dmg

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	_, err := r.Seek(-512, io.SeekEnd)
	if err != nil {
		return nil, err
	}

	var sig [4]byte

	_, err = io.ReadFull(r, sig[:])
	if err != nil {
		return nil, err
	}

	if sig != [4]byte{'k', 'o', 'l', 'y'} {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "Apple disk image",
		Type: "DMG",
	}, nil
}
