package fits

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
	var hdr [30]byte

	_, err := io.ReadFull(r, hdr[:])
	if err != nil {
		return nil, err
	}

	if !bytes.HasPrefix(hdr[:], []byte("SIMPLE  =")) {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "FITS astronomical image",
	}, nil
}
