package aac

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var hdr [2]byte

	_, err := io.ReadFull(r, hdr[:])
	if err != nil {
		return nil, err
	}

	if hdr[0] != 0xff || (hdr[1]&0xf0) != 0xf0 {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "AAC audio (ADTS)",
	}, nil
}
