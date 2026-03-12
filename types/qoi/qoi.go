package qoi

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
	var magic [4]byte

	_, err := io.ReadFull(r, magic[:])
	if err != nil {
		return nil, err
	}

	if magic != [4]byte{'q', 'o', 'i', 'f'} {
		return nil, types.ErrUnknownFormat
	}

	var (
		width    uint32
		height   uint32
		channels uint8
		colors   string
	)

	err = binary.Read(r, binary.BigEndian, &width)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &height)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &channels)
	if err != nil {
		return nil, err
	}

	switch channels {
	case 3:
		colors = "RGB"
	case 4:
		colors = "RGBA"
	default:
		colors = fmt.Sprintf("%d channels", channels)
	}

	return &types.Metadata{
		Name: "QOI image",
		Type: colors,
		Info: fmt.Sprintf("%dx%d", width, height),
	}, nil
}
