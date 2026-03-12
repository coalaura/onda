package mp4

import (
	"fmt"
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var size [4]byte

	_, err := io.ReadFull(r, size[:])
	if err != nil {
		return nil, err
	}

	var name [4]byte

	_, err = io.ReadFull(r, name[:])
	if err != nil {
		return nil, err
	}

	if name != [4]byte{'f', 't', 'y', 'p'} {
		return nil, types.ErrUnknownFormat
	}

	var brand [4]byte

	_, err = io.ReadFull(r, brand[:])
	if err != nil {
		return nil, err
	}

	return &types.Metadata{
		Name: "MP4 video",
		Type: fmt.Sprintf("brand %.4s", brand[:]),
	}, nil
}
