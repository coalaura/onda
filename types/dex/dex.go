package dex

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var s [8]byte

	_, err := io.ReadFull(r, s[:])
	if err != nil {
		return nil, err
	}

	if s[0] != 'd' || s[1] != 'e' || s[2] != 'x' || s[3] != '\n' {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "Dalvik bytecode",
		Type: string(s[4:7]),
	}, nil
}
