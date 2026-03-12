package pcx

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
	var (
		magic byte
		ver   byte
		enc   byte
		bpp   byte
		xmin  uint16
		ymin  uint16
		xmax  uint16
		ymax  uint16
	)

	err := binary.Read(r, binary.LittleEndian, &magic)
	if err != nil {
		return nil, err
	}

	if magic != 0x0a {
		return nil, types.ErrUnknownFormat
	}

	err = binary.Read(r, binary.LittleEndian, &ver)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &enc)
	if err != nil {
		return nil, err
	}

	if enc != 1 {
		return nil, types.ErrUnknownFormat
	}

	err = binary.Read(r, binary.LittleEndian, &bpp)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &xmin)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &ymin)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &xmax)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &ymax)
	if err != nil {
		return nil, err
	}

	width := xmax - xmin + 1
	height := ymax - ymin + 1

	return &types.Metadata{
		Name: "PCX image",
		Info: fmt.Sprintf("%dx%d %dbpp", width, height, bpp),
	}, nil
}
