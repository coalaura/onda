package qcow2

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
		magic   uint32
		version uint32
	)

	err := binary.Read(r, binary.BigEndian, &magic)
	if err != nil {
		return nil, err
	}

	if magic != 0x514649fb {
		return nil, types.ErrUnknownFormat
	}

	err = binary.Read(r, binary.BigEndian, &version)
	if err != nil {
		return nil, err
	}

	return &types.Metadata{
		Name: "QCOW disk image",
		Type: fmt.Sprintf("version %d", version),
	}, nil
}
