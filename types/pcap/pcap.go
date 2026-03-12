package pcap

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
	var magic uint32

	err := binary.Read(r, binary.LittleEndian, &magic)
	if err != nil {
		return nil, err
	}

	var order binary.ByteOrder

	switch magic {
	case 0xa1b2c3d4:
		order = binary.BigEndian
	case 0xd4c3b2a1:
		order = binary.LittleEndian
	default:
		return nil, types.ErrUnknownFormat
	}

	var (
		major uint16
		minor uint16
	)

	err = binary.Read(r, order, &major)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, order, &minor)
	if err != nil {
		return nil, err
	}

	_, err = r.Seek(8, io.SeekCurrent)
	if err != nil {
		return nil, err
	}

	var link uint32

	err = binary.Read(r, order, &link)
	if err != nil {
		return nil, err
	}

	return &types.Metadata{
		Name: "PCAP capture",
		Type: fmt.Sprintf("v%d.%d", major, minor),
		Info: fmt.Sprintf("linktype %d", link),
	}, nil
}
