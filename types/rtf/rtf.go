package rtf

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var hdr [5]byte

	_, err := io.ReadFull(r, hdr[:])
	if err != nil {
		return nil, err
	}

	if hdr != [5]byte{'{', '\\', 'r', 't', 'f'} {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "RTF document",
	}, nil
}
