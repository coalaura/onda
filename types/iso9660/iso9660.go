package iso9660

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	_, err := r.Seek(16*2048, io.SeekStart)
	if err != nil {
		return nil, err
	}

	var hdr [7]byte

	_, err = io.ReadFull(r, hdr[:])
	if err != nil {
		return nil, err
	}

	if hdr[0] != 1 {
		return nil, types.ErrUnknownFormat
	}

	if hdr[1] != 'C' || hdr[2] != 'D' || hdr[3] != '0' || hdr[4] != '0' || hdr[5] != '1' {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "ISO9660 disk image",
	}, nil
}
