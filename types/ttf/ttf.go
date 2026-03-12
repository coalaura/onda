package ttf

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
	var version uint32

	err := binary.Read(r, binary.BigEndian, &version)
	if err != nil {
		return nil, err
	}

	if version != 0x00010000 {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "TrueType font",
	}, nil
}
