package pcapng

import (
	"encoding/binary"
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var magic uint32

	err := binary.Read(r, binary.LittleEndian, &magic)
	if err != nil {
		return nil, err
	}

	if magic != 0x0a0d0d0a {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "PCAP Next Generation capture",
	}, nil
}
