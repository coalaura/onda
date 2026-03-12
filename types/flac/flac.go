package flac

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var sig [4]byte

	_, err := io.ReadFull(r, sig[:])
	if err != nil {
		return nil, err
	}

	if sig != [4]byte{'f', 'L', 'a', 'C'} {
		return nil, types.ErrUnknownFormat
	}

	var hdr [4]byte

	_, err = io.ReadFull(r, hdr[:])
	if err != nil {
		return nil, err
	}

	if hdr[0]&0x7f != 0 {
		return nil, types.ErrUnknownFormat
	}

	length := int(hdr[1])<<16 | int(hdr[2])<<8 | int(hdr[3])

	if length < 34 {
		return nil, types.ErrUnknownFormat
	}

	var buf [34]byte

	_, err = io.ReadFull(r, buf[:])
	if err != nil {
		return nil, err
	}

	x := binary.BigEndian.Uint64(buf[10:18])

	sampleRate := int((x >> 44) & 0xfffff)
	channels := int((x>>41)&7) + 1
	bits := int((x>>36)&31) + 1

	return &types.Metadata{
		Name: "FLAC audio",
		Info: fmt.Sprintf("%d Hz, %d-bit, %d ch", sampleRate, bits, channels),
	}, nil
}
