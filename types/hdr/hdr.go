package hdr

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

var sig = []byte("#?RADIANCE")

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var buf [10]byte

	_, err := io.ReadFull(r, buf[:])
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(sig); i++ {
		if buf[i] != sig[i] {
			return nil, types.ErrUnknownFormat
		}
	}

	return &types.Metadata{
		Name: "Radiance HDR image",
	}, nil
}
