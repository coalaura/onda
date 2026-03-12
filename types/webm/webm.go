package webm

import (
	"bytes"
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

var ebml = []byte{0x1a, 0x45, 0xdf, 0xa3}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var buf [64]byte

	_, err := io.ReadFull(r, buf[:])
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(buf[:4], ebml) {
		return nil, types.ErrUnknownFormat
	}

	if bytes.Contains(buf[:], []byte("webm")) {
		return &types.Metadata{
			Name: "WebM video",
		}, nil
	}

	return nil, types.ErrUnknownFormat
}
