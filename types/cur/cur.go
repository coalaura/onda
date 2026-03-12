package cur

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
		reserved uint16
		typ      uint16
		count    uint16
	)

	err := binary.Read(r, binary.LittleEndian, &reserved)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &typ)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &count)
	if err != nil {
		return nil, err
	}

	if reserved != 0 || typ != 2 {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "Windows cursor",
		Info: "images present",
	}, nil
}
