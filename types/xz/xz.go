package xz

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

var sig = [6]byte{0xfd, '7', 'z', 'X', 'Z', 0x00}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var s [6]byte

	_, err := io.ReadFull(r, s[:])
	if err != nil {
		return nil, err
	}

	if s != sig {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "XZ compressed archive",
	}, nil
}
