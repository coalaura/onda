package ole

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

var sig = [8]byte{
	0xd0, 0xcf, 0x11, 0xe0,
	0xa1, 0xb1, 0x1a, 0xe1,
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var s [8]byte

	_, err := io.ReadFull(r, s[:])
	if err != nil {
		return nil, err
	}

	if s != sig {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "OLE compound document",
	}, nil
}
