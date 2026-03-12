package pdf

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var hdr [8]byte

	_, err := io.ReadFull(r, hdr[:])
	if err != nil {
		return nil, err
	}

	if string(hdr[:5]) != "%PDF-" {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "PDF document",
		Type: string(hdr[5:8]),
	}, nil
}
