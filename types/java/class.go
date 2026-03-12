package javaclass

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
	var magic uint32

	err := binary.Read(r, binary.BigEndian, &magic)
	if err != nil {
		return nil, err
	}

	if magic != 0xCAFEBABE {
		return nil, types.ErrUnknownFormat
	}

	var minor uint16

	err = binary.Read(r, binary.BigEndian, &minor)
	if err != nil {
		return nil, err
	}

	var major uint16

	err = binary.Read(r, binary.BigEndian, &major)
	if err != nil {
		return nil, err
	}

	return &types.Metadata{
		Name: "Java class",
		Info: fmt.Sprintf("version %d.%d", major, minor),
	}, nil
}
