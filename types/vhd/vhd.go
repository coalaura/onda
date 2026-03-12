package vhd

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	_, err := r.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	var buf [8]byte

	_, err = r.Seek(-512, io.SeekEnd)
	if err != nil {
		return nil, types.ErrUnknownFormat
	}

	_, err = io.ReadFull(r, buf[:])
	if err != nil {
		return nil, err
	}

	if string(buf[:]) != "conectix" {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "Virtual Hard Disk",
		Type: "VHD",
	}, nil
}
