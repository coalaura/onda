package rpm

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

var magic = [4]byte{0xed, 0xab, 0xee, 0xdb}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var m [4]byte

	_, err := io.ReadFull(r, m[:])
	if err != nil {
		return nil, err
	}

	if m != magic {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "RPM package",
	}, nil
}
