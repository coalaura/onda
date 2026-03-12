package otf

import (
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

	if sig != [4]byte{'O', 'T', 'T', 'O'} {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "OpenType font",
	}, nil
}
