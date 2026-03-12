package gif

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
	var hdr [6]byte

	_, err := io.ReadFull(r, hdr[:])
	if err != nil {
		return nil, err
	}

	if string(hdr[:3]) != "GIF" {
		return nil, types.ErrUnknownFormat
	}

	var width uint16

	err = binary.Read(r, binary.LittleEndian, &width)
	if err != nil {
		return nil, err
	}

	var height uint16

	err = binary.Read(r, binary.LittleEndian, &height)
	if err != nil {
		return nil, err
	}

	return &types.Metadata{
		Name: "GIF image",
		Type: string(hdr[3:]),
		Info: fmt.Sprintf("%dx%d", width, height),
	}, nil
}
