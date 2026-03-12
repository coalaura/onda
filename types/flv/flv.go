package flv

import (
	"fmt"
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var hdr [9]byte

	_, err := io.ReadFull(r, hdr[:])
	if err != nil {
		return nil, err
	}

	if hdr[0] != 'F' || hdr[1] != 'L' || hdr[2] != 'V' {
		return nil, types.ErrUnknownFormat
	}

	flags := hdr[4]

	var streams string

	switch flags & 0x05 {
	case 0x01:
		streams = "video"
	case 0x04:
		streams = "audio"
	case 0x05:
		streams = "audio+video"
	default:
		streams = "unknown streams"
	}

	return &types.Metadata{
		Name: "Flash Video",
		Type: streams,
		Info: fmt.Sprintf("version %d", hdr[3]),
	}, nil
}
