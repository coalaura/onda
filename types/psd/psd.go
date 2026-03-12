package psd

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

	if sig != [4]byte{'8', 'B', 'P', 'S'} {
		return nil, types.ErrUnknownFormat
	}

	var version uint16

	err = binary.Read(r, binary.BigEndian, &version)
	if err != nil {
		return nil, err
	}

	if version != 1 {
		return nil, types.ErrUnknownFormat
	}

	_, err = r.Seek(6, io.SeekCurrent)
	if err != nil {
		return nil, err
	}

	var (
		channels uint16
		height   uint32
		width    uint32
		depth    uint16
	)

	err = binary.Read(r, binary.BigEndian, &channels)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &height)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &width)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &depth)
	if err != nil {
		return nil, err
	}

	return &types.Metadata{
		Name: "Photoshop document",
		Info: fmt.Sprintf("%dx%d, %d channels, %d-bit", width, height, channels, depth),
	}, nil
}
