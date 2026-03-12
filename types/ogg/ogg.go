package ogg

import (
	"bytes"
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

	if sig != [4]byte{'O', 'g', 'g', 'S'} {
		return nil, types.ErrUnknownFormat
	}

	var buf [256]byte

	n, err := io.ReadFull(r, buf[:])
	if err != nil && err != io.ErrUnexpectedEOF {
		return nil, err
	}

	data := buf[:n]

	if bytes.Contains(data, []byte("vorbis")) {
		return &types.Metadata{Name: "Ogg Vorbis audio"}, nil
	}

	if bytes.Contains(data, []byte("OpusHead")) {
		return &types.Metadata{Name: "Ogg Opus audio"}, nil
	}

	if bytes.Contains(data, []byte("theora")) {
		return &types.Metadata{Name: "Ogg Theora video"}, nil
	}

	return &types.Metadata{
		Name: "Ogg container",
	}, nil
}
