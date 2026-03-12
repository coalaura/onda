package dicom

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	_, err := r.Seek(128, io.SeekStart)
	if err != nil {
		return nil, err
	}

	var sig [4]byte

	_, err = io.ReadFull(r, sig[:])
	if err != nil {
		return nil, err
	}

	if sig != [4]byte{'D', 'I', 'C', 'M'} {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "DICOM medical image",
	}, nil
}
