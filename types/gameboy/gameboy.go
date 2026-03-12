package gameboy

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
	_, err := r.Seek(0x104, io.SeekStart)
	if err != nil {
		return nil, err
	}

	var logo [48]byte

	_, err = io.ReadFull(r, logo[:])
	if err != nil {
		return nil, err
	}

	if logo[0] != 0xCE || logo[1] != 0xED {
		return nil, types.ErrUnknownFormat
	}

	_, err = r.Seek(0x134, io.SeekStart)
	if err != nil {
		return nil, err
	}

	var title [16]byte

	_, err = io.ReadFull(r, title[:])
	if err != nil {
		return nil, err
	}

	return &types.Metadata{
		Name: "Game Boy ROM",
		Info: fmt.Sprintf("title %.16s", title),
	}, nil
}
