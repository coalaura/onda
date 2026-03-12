package swf

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
	var hdr [8]byte

	_, err := io.ReadFull(r, hdr[:])
	if err != nil {
		return nil, err
	}

	if hdr[1] != 'W' || hdr[2] != 'S' {
		return nil, types.ErrUnknownFormat
	}

	var typ string

	switch hdr[0] {
	case 'F':
		typ = "uncompressed"
	case 'C':
		typ = "zlib compressed"
	case 'Z':
		typ = "lzma compressed"
	default:
		return nil, types.ErrUnknownFormat
	}

	version := hdr[3]

	return &types.Metadata{
		Name: "Shockwave Flash",
		Type: typ,
		Info: fmt.Sprintf("version %d", version),
	}, nil
}
