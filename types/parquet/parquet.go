package parquet

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

var sig = [4]byte{'P', 'A', 'R', '1'}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var s [4]byte

	_, err := io.ReadFull(r, s[:])
	if err != nil {
		return nil, err
	}

	if s != sig {
		return nil, types.ErrUnknownFormat
	}

	return &types.Metadata{
		Name: "Apache Parquet",
	}, nil
}
