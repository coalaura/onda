package dds

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var magic [4]byte

	_, err := io.ReadFull(r, magic[:])
	if err != nil {
		return nil, err
	}

	if magic != [4]byte{'D', 'D', 'S', ' '} {
		return nil, types.ErrUnknownFormat
	}

	var size uint32

	err = binary.Read(r, binary.LittleEndian, &size)
	if err != nil {
		return nil, err
	}

	if size != 124 {
		return nil, types.ErrUnknownFormat
	}

	var (
		flags  uint32
		height uint32
		width  uint32
	)

	err = binary.Read(r, binary.LittleEndian, &flags)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &height)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &width)
	if err != nil {
		return nil, err
	}

	return &types.Metadata{
		Name: "DirectDraw Surface",
		Info: fmt.Sprintf("%dx%d", width, height),
	}, nil
}
