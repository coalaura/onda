package webp

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
	var hdr [12]byte

	_, err := io.ReadFull(r, hdr[:])
	if err != nil {
		return nil, err
	}

	if hdr[0] != 'R' || hdr[1] != 'I' || hdr[2] != 'F' || hdr[3] != 'F' {
		return nil, types.ErrUnknownFormat
	}

	if hdr[8] != 'W' || hdr[9] != 'E' || hdr[10] != 'B' || hdr[11] != 'P' {
		return nil, types.ErrUnknownFormat
	}

	var chunk [4]byte

	_, err = io.ReadFull(r, chunk[:])
	if err != nil {
		return nil, err
	}

	var size uint32

	err = binary.Read(r, binary.LittleEndian, &size)
	if err != nil {
		return nil, err
	}

	switch string(chunk[:]) {
	case "VP8 ":
		_, err = r.Seek(7, io.SeekCurrent)
		if err != nil {
			return nil, err
		}

		var dims [4]byte

		_, err = io.ReadFull(r, dims[:])
		if err != nil {
			return nil, err
		}

		w := int(binary.LittleEndian.Uint16(dims[0:2]) & 0x3fff)
		h := int(binary.LittleEndian.Uint16(dims[2:4]) & 0x3fff)

		return &types.Metadata{
			Name: "WebP image",
			Type: "VP8",
			Info: fmt.Sprintf("%dx%d", w, h),
		}, nil

	case "VP8L":
		var b [5]byte

		_, err = io.ReadFull(r, b[:])
		if err != nil {
			return nil, err
		}

		if b[0] != 0x2f {
			return nil, types.ErrUnknownFormat
		}

		bits := binary.LittleEndian.Uint32(b[1:5])

		w := int((bits & 0x3fff) + 1)
		h := int(((bits >> 14) & 0x3fff) + 1)

		return &types.Metadata{
			Name: "WebP image",
			Type: "VP8L",
			Info: fmt.Sprintf("%dx%d", w, h),
		}, nil

	case "VP8X":
		_, err = r.Seek(4, io.SeekCurrent)
		if err != nil {
			return nil, err
		}

		var b [6]byte

		_, err = io.ReadFull(r, b[:])
		if err != nil {
			return nil, err
		}

		w := int(uint32(b[0])|uint32(b[1])<<8|uint32(b[2])<<16) + 1
		h := int(uint32(b[3])|uint32(b[4])<<8|uint32(b[5])<<16) + 1

		return &types.Metadata{
			Name: "WebP image",
			Type: "VP8X",
			Info: fmt.Sprintf("%dx%d", w, h),
		}, nil
	}

	return &types.Metadata{
		Name: "WebP image",
	}, nil
}
