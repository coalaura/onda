package mkv

import (
	"bytes"
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

var ebml = []byte{0x1A, 0x45, 0xDF, 0xA3}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var hdr [4]byte

	_, err := io.ReadFull(r, hdr[:])
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(hdr[:], ebml) {
		return nil, types.ErrUnknownFormat
	}

	var buf [256]byte

	n, err := io.ReadFull(r, buf[:])
	if err != nil && err != io.ErrUnexpectedEOF {
		return nil, err
	}

	data := buf[:n]

	if bytes.Contains(data, []byte("webm")) {
		return &types.Metadata{
			Name: "WebM video",
		}, nil
	}

	if bytes.Contains(data, []byte("matroska")) {
		return &types.Metadata{
			Name: "Matroska video",
		}, nil
	}

	return &types.Metadata{
		Name: "EBML container",
	}, nil
}
