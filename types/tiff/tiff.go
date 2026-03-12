package tiff

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
	var hdr [4]byte

	_, err := io.ReadFull(r, hdr[:])
	if err != nil {
		return nil, err
	}

	var order binary.ByteOrder

	switch {
	case hdr[0] == 'I' && hdr[1] == 'I':
		order = binary.LittleEndian
	case hdr[0] == 'M' && hdr[1] == 'M':
		order = binary.BigEndian
	default:
		return nil, types.ErrUnknownFormat
	}

	version := order.Uint16(hdr[2:4])

	if version != 42 {
		return nil, types.ErrUnknownFormat
	}

	var ifd uint32

	err = binary.Read(r, order, &ifd)
	if err != nil {
		return nil, err
	}

	_, err = r.Seek(int64(ifd), io.SeekStart)
	if err != nil {
		return nil, err
	}

	var count uint16

	err = binary.Read(r, order, &count)
	if err != nil {
		return nil, err
	}

	var (
		width  uint32
		height uint32
	)

	for i := 0; i < int(count); i++ {
		var (
			tag uint16
			typ uint16
			num uint32
			val uint32
		)

		err = binary.Read(r, order, &tag)
		if err != nil {
			return nil, err
		}

		err = binary.Read(r, order, &typ)
		if err != nil {
			return nil, err
		}

		err = binary.Read(r, order, &num)
		if err != nil {
			return nil, err
		}

		err = binary.Read(r, order, &val)
		if err != nil {
			return nil, err
		}

		if tag == 256 {
			width = val
		}

		if tag == 257 {
			height = val
		}

		if width != 0 && height != 0 {
			break
		}
	}

	var info string

	if width != 0 && height != 0 {
		info = fmt.Sprintf("%dx%d", width, height)
	}

	return &types.Metadata{
		Name: "TIFF image",
		Info: info,
	}, nil
}
