package sevenz

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

var sig = [6]byte{0x37, 0x7a, 0xbc, 0xaf, 0x27, 0x1c}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var s [6]byte

	_, err := io.ReadFull(r, s[:])
	if err != nil {
		return nil, err
	}

	if s != sig {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "7-Zip archive",
	}, nil
}
