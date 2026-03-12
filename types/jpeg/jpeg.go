package jpeg

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
	var soi [2]byte

	_, err := io.ReadFull(r, soi[:])
	if err != nil {
		return nil, err
	}

	if soi != [2]byte{0xff, 0xd8} {
		return nil, types.ErrUnknownFormat
	}

	for {
		var marker [2]byte

		_, err = io.ReadFull(r, marker[:])
		if err != nil {
			return nil, err
		}

		if marker[0] != 0xff {
			return nil, types.ErrUnknownFormat
		}

		if marker[1] == 0xd9 || marker[1] == 0xda {
			break
		}

		var length uint16

		err = binary.Read(r, binary.BigEndian, &length)
		if err != nil {
			return nil, err
		}

		if marker[1] >= 0xc0 && marker[1] <= 0xc3 {
			_, err = r.Seek(1, io.SeekCurrent)
			if err != nil {
				return nil, err
			}

			var height uint16

			err = binary.Read(r, binary.BigEndian, &height)
			if err != nil {
				return nil, err
			}

			var width uint16

			err = binary.Read(r, binary.BigEndian, &width)
			if err != nil {
				return nil, err
			}

			return &types.Metadata{
				Name: "JPEG image",
				Info: fmt.Sprintf("%dx%d", width, height),
			}, nil
		}

		_, err = r.Seek(int64(length-2), io.SeekCurrent)
		if err != nil {
			return nil, err
		}
	}

	return &types.Metadata{
		Name: "JPEG image",
	}, nil
}
