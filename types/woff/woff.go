package woff

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
	var sig [4]byte

	_, err := io.ReadFull(r, sig[:])
	if err != nil {
		return nil, err
	}

	if sig != [4]byte{'w', 'O', 'F', 'F'} {
		return nil, types.ErrUnknownFormat
	}

	var flavor uint32

	err = binary.Read(r, binary.BigEndian, &flavor)
	if err != nil {
		return nil, err
	}

	return &types.Metadata{
		Name: "WOFF font",
		Type: fmt.Sprintf("flavor 0x%08x", flavor),
	}, nil
}
