package png

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

var sig = [8]byte{137, 'P', 'N', 'G', 13, 10, 26, 10}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var s [8]byte

	_, err := io.ReadFull(r, s[:])
	if err != nil {
		return nil, err
	}

	if s != sig {
		return nil, types.ErrUnknownFormat
	}

	var length uint32

	err = binary.Read(r, binary.BigEndian, &length)
	if err != nil {
		return nil, err
	}

	var name [4]byte

	_, err = io.ReadFull(r, name[:])
	if err != nil {
		return nil, err
	}

	if name != [4]byte{'I', 'H', 'D', 'R'} {
		return nil, types.ErrUnknownFormat
	}

	var width uint32

	err = binary.Read(r, binary.BigEndian, &width)
	if err != nil {
		return nil, err
	}

	var height uint32

	err = binary.Read(r, binary.BigEndian, &height)
	if err != nil {
		return nil, err
	}

	return &types.Metadata{
		Name: "PNG image",
		Info: fmt.Sprintf("%dx%d", width, height),
	}, nil
}
