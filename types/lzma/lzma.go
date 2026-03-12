package lzma

import (
	"encoding/binary"
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var (
		props byte
		dict  uint32
	)

	err := binary.Read(r, binary.LittleEndian, &props)
	if err != nil {
		return nil, err
	}

	if props > (4*5+4)*9+8 {
		return nil, types.ErrUnknownFormat
	}

	err = binary.Read(r, binary.LittleEndian, &dict)
	if err != nil {
		return nil, err
	}

	if dict == 0 {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "LZMA compressed data",
	}, nil
}
