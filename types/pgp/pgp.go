package pgp

import (
	"bytes"
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

var sig = []byte("-----BEGIN PGP")

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var buf [32]byte

	_, err := io.ReadFull(r, buf[:])
	if err != nil {
		return nil, err
	}

	if !bytes.HasPrefix(buf[:], sig) {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "PGP armored data",
	}, nil
}
