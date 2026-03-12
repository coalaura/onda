package avif

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
		size  uint32
		ftyp  [4]byte
		brand [4]byte
	)

	err := binary.Read(r, binary.BigEndian, &size)
	if err != nil {
		return nil, err
	}

	_, err = io.ReadFull(r, ftyp[:])
	if err != nil {
		return nil, err
	}

	if ftyp != [4]byte{'f', 't', 'y', 'p'} {
		return nil, types.ErrUnknownFormat
	}

	_, err = io.ReadFull(r, brand[:])
	if err != nil {
		return nil, err
	}

	if brand != [4]byte{'a', 'v', 'i', 'f'} {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "AVIF image",
	}, nil
}
