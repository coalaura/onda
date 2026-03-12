package bmp

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
	var sig [2]byte

	_, err := io.ReadFull(r, sig[:])
	if err != nil {
		return nil, err
	}

	if sig != [2]byte{'B', 'M'} {
		return nil, types.ErrUnknownFormat
	}

	_, err = r.Seek(16, io.SeekCurrent)
	if err != nil {
		return nil, err
	}

	var width int32

	err = binary.Read(r, binary.LittleEndian, &width)
	if err != nil {
		return nil, err
	}

	var height int32

	err = binary.Read(r, binary.LittleEndian, &height)
	if err != nil {
		return nil, err
	}

	_, err = r.Seek(2, io.SeekCurrent)
	if err != nil {
		return nil, err
	}

	var bpp uint16

	err = binary.Read(r, binary.LittleEndian, &bpp)
	if err != nil {
		return nil, err
	}

	return &types.Metadata{
		Name: "Bitmap image",
		Type: fmt.Sprintf("%d-bit", bpp),
		Info: fmt.Sprintf("%dx%d", width, height),
	}, nil
}
