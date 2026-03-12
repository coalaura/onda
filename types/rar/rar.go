package rar

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

var (
	rar4 = [7]byte{'R', 'a', 'r', '!', 0x1a, 0x07, 0x00}
	rar5 = [8]byte{'R', 'a', 'r', '!', 0x1a, 0x07, 0x01, 0x00}
)

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var b [8]byte

	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return nil, err
	}

	if b == rar5 {
		return &types.Metadata{
			Name: "RAR archive",
			Type: "RAR5",
		}, nil
	}

	if b[0] == rar4[0] && b[1] == rar4[1] && b[2] == rar4[2] && b[3] == rar4[3] && b[4] == rar4[4] && b[5] == rar4[5] && b[6] == rar4[6] {
		return &types.Metadata{
			Name: "RAR archive",
			Type: "RAR4",
		}, nil
	}

	return nil, types.ErrUnknownFormat
}
