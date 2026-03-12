package vdi

import (
	"bytes"
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

var sig = []byte("<<< Oracle VM VirtualBox Disk Image >>>")

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var buf [64]byte

	_, err := io.ReadFull(r, buf[:])
	if err != nil {
		return nil, err
	}

	if !bytes.Contains(buf[:], sig) {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "VirtualBox disk image",
		Type: "VDI",
	}, nil
}
