package mpeg

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var s [4]byte

	_, err := io.ReadFull(r, s[:])
	if err != nil {
		return nil, err
	}

	if s != [4]byte{0x00, 0x00, 0x01, 0xba} && s != [4]byte{0x00, 0x00, 0x01, 0xb3} {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "MPEG video",
		Type: "Program stream",
	}, nil
}
