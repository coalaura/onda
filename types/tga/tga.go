package tga

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
	var header [18]byte

	_, err := io.ReadFull(r, header[:])
	if err != nil {
		return nil, err
	}

	imgType := header[2]

	if imgType == 0 {
		return nil, types.ErrUnknownFormat
	}

	width := binary.LittleEndian.Uint16(header[12:14])
	height := binary.LittleEndian.Uint16(header[14:16])

	var kind string

	switch imgType {
	case 1:
		kind = "color-mapped"
	case 2:
		kind = "truecolor"
	case 3:
		kind = "grayscale"
	case 9:
		kind = "RLE color-mapped"
	case 10:
		kind = "RLE truecolor"
	case 11:
		kind = "RLE grayscale"
	default:
		kind = fmt.Sprintf("type %d", imgType)
	}

	return &types.Metadata{
		Name: "TGA image",
		Type: kind,
		Info: fmt.Sprintf("%dx%d", width, height),
	}, nil
}
