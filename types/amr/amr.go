package amr

import (
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

var (
	sigAMR   = []byte("#!AMR\n")
	sigAMRWB = []byte("#!AMR-WB\n")
)

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var buf [9]byte

	_, err := io.ReadFull(r, buf[:])
	if err != nil {
		return nil, err
	}

	if string(buf[:6]) == string(sigAMR) {
		return &types.Metadata{
			Name: "AMR audio",
		}, nil
	}

	if string(buf[:9]) == string(sigAMRWB) {
		return &types.Metadata{
			Name: "AMR audio",
			Type: "wideband",
		}, nil
	}

	return nil, types.ErrUnknownFormat
}
