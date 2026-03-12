package icns

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

	if magic != [4]byte{'i', 'c', 'n', 's'} {
		return nil, types.ErrUnknownFormat
	}

	var size uint32

	err = binary.Read(r, binary.BigEndian, &size)
	if err != nil {
		return nil, err
	}

	return &types.Metadata{
		Name: "Apple icon",
		Info: fmt.Sprintf("%d bytes", size),
	}, nil
}
