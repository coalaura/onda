package zip

import (
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

	if sig != [4]byte{'P', 'K', 3, 4} && sig != [4]byte{'P', 'K', 5, 6} && sig != [4]byte{'P', 'K', 7, 8} {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "ZIP archive",
	}, nil
}
