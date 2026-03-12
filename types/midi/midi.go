package midi

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

	if sig != [4]byte{'M', 'T', 'h', 'd'} {
		return nil, types.ErrUnknownFormat
	}

	var size uint32

	err = binary.Read(r, binary.BigEndian, &size)
	if err != nil {
		return nil, err
	}

	if size < 6 {
		return nil, types.ErrUnknownFormat
	}

	var (
		format uint16
		tracks uint16
	)

	err = binary.Read(r, binary.BigEndian, &format)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &tracks)
	if err != nil {
		return nil, err
	}

	return &types.Metadata{
		Name: "MIDI sequence",
		Type: fmt.Sprintf("format %d", format),
		Info: fmt.Sprintf("%d tracks", tracks),
	}, nil
}
