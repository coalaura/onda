package deb

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

var arMagic = [8]byte{'!', '<', 'a', 'r', 'c', 'h', '>', '\n'}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var hdr [8]byte

	_, err := io.ReadFull(r, hdr[:])
	if err != nil {
		return nil, err
	}

	if hdr != arMagic {
		return nil, types.ErrUnknownFormat
	}

	var name [16]byte

	_, err = io.ReadFull(r, name[:])
	if err != nil {
		return nil, err
	}

	if name[0] != 'd' || name[1] != 'e' || name[2] != 'b' || name[3] != 'i' {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "Debian package",
	}, nil
}
