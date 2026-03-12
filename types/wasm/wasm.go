package wasm

import (
	"encoding/binary"
	"fmt"
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

	if sig != [4]byte{0x00, 'a', 's', 'm'} {
		return nil, types.ErrUnknownFormat
	}

	var version uint32

	err = binary.Read(r, binary.LittleEndian, &version)
	if err != nil {
		return nil, err
	}

	return &types.Metadata{
		Name: "WebAssembly module",
		Type: fmt.Sprintf("v%d", version),
	}, nil
}
