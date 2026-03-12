package au

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

	if magic != [4]byte{'.', 's', 'n', 'd'} {
		return nil, types.ErrUnknownFormat
	}

	var (
		offset uint32
		size   uint32
	)

	err = binary.Read(r, binary.BigEndian, &offset)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &size)
	if err != nil {
		return nil, err
	}

	return &types.Metadata{
		Name: "Sun AU audio",
		Info: fmt.Sprintf("data %d bytes", size),
	}, nil
}
