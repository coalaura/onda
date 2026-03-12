package nes

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
	var hdr [16]byte

	_, err := io.ReadFull(r, hdr[:])
	if err != nil {
		return nil, err
	}

	if hdr[0] != 'N' || hdr[1] != 'E' || hdr[2] != 'S' || hdr[3] != 0x1A {
		return nil, types.ErrUnknownFormat
	}

	prg := int(hdr[4]) * 16
	chr := int(hdr[5]) * 8

	return &types.Metadata{
		Name: "NES ROM",
		Info: fmt.Sprintf("PRG %dKB, CHR %dKB", prg, chr),
	}, nil
}
