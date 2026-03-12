package aiff

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
		form [4]byte
		size uint32
		aiff [4]byte
	)

	_, err := io.ReadFull(r, form[:])
	if err != nil {
		return nil, err
	}

	if form != [4]byte{'F', 'O', 'R', 'M'} {
		return nil, types.ErrUnknownFormat
	}

	err = binary.Read(r, binary.BigEndian, &size)
	if err != nil {
		return nil, err
	}

	_, err = io.ReadFull(r, aiff[:])
	if err != nil {
		return nil, err
	}

	if aiff != [4]byte{'A', 'I', 'F', 'F'} && aiff != [4]byte{'A', 'I', 'F', 'C'} {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "AIFF audio",
		Type: string(aiff[:]),
		Info: fmt.Sprintf("%d bytes container", size+8),
	}, nil
}
