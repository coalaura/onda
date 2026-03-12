package wave64

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

var sig = [16]byte{
	0x72, 0x69, 0x66, 0x66,
	0x2E, 0x91, 0xCF, 0x11,
	0xA5, 0xD6, 0x28, 0xDB,
	0x04, 0xC1, 0x00, 0x00,
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var s [16]byte

	_, err := io.ReadFull(r, s[:])
	if err != nil {
		return nil, err
	}

	if s != sig {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "Sony Wave64 audio",
	}, nil
}
