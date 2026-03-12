package gpt

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	_, err := r.Seek(512, io.SeekStart)
	if err != nil {
		return nil, err
	}

	var sig [8]byte

	_, err = io.ReadFull(r, sig[:])
	if err != nil {
		return nil, err
	}

	if sig != [8]byte{'E', 'F', 'I', ' ', 'P', 'A', 'R', 'T'} {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "GUID partition table",
		Type: "GPT",
	}, nil
}
