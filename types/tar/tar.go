package tar

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	_, err := r.Seek(257, io.SeekStart)
	if err != nil {
		return nil, err
	}

	var magic [5]byte

	_, err = io.ReadFull(r, magic[:])
	if err != nil {
		return nil, err
	}

	if magic != [5]byte{'u', 's', 't', 'a', 'r'} {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "TAR archive",
	}, nil
}
