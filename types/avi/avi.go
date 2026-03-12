package avi

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
	var (
		riff [4]byte
		size uint32
		avi  [4]byte
	)

	_, err := io.ReadFull(r, riff[:])
	if err != nil {
		return nil, err
	}

	if riff != [4]byte{'R', 'I', 'F', 'F'} {
		return nil, types.ErrUnknownFormat
	}

	err = binary.Read(r, binary.LittleEndian, &size)
	if err != nil {
		return nil, err
	}

	_, err = io.ReadFull(r, avi[:])
	if err != nil {
		return nil, err
	}

	if avi != [4]byte{'A', 'V', 'I', ' '} {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "AVI video",
		Info: fmt.Sprintf("%d bytes container", size+8),
	}, nil
}
