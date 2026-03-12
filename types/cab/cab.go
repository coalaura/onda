package cab

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

	if magic != [4]byte{'M', 'S', 'C', 'F'} {
		return nil, types.ErrUnknownFormat
	}

	var (
		reserved uint32
		size     uint32
	)

	err = binary.Read(r, binary.LittleEndian, &reserved)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &size)
	if err != nil {
		return nil, err
	}

	return &types.Metadata{
		Name: "Microsoft Cabinet archive",
		Info: fmt.Sprintf("%d bytes", size),
	}, nil
}
