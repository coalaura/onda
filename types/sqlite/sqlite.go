package sqlite

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

var sig = [16]byte{
	'S', 'Q', 'L', 'i', 't', 'e', ' ', 'f', 'o', 'r', 'm', 'a', 't', ' ', '3', 0,
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var s [16]byte

	_, err := io.ReadFull(r, s[:])
	if err != nil {
		return nil, err
	}

	if s != sig {
		return nil, types.ErrUnknownFormat
	}

	var pageSize uint16

	err = binary.Read(r, binary.BigEndian, &pageSize)
	if err != nil {
		return nil, err
	}

	var info string

	if pageSize == 1 {
		info = "page size 65536"
	} else {
		info = fmt.Sprintf("page size %d", pageSize)
	}

	return &types.Metadata{
		Name: "SQLite database",
		Info: info,
	}, nil
}
