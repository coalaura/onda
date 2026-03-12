package pnm

import (
	"fmt"
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var m [2]byte

	_, err := io.ReadFull(r, m[:])
	if err != nil {
		return nil, err
	}

	if m[0] != 'P' {
		return nil, types.ErrUnknownFormat
	}

	var name string

	switch m[1] {
	case '1', '4':
		name = "PBM image"
	case '2', '5':
		name = "PGM image"
	case '3', '6':
		name = "PPM image"
	case '7':
		name = "PAM image"
	default:
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: name,
		Type: fmt.Sprintf("P%c", m[1]),
	}, nil
}
