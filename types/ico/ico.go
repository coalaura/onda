package ico

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
	var (
		reserved uint16
		typ      uint16
		count    uint16
	)

	err := binary.Read(r, binary.LittleEndian, &reserved)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &typ)
	if err != nil {
		return nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &count)
	if err != nil {
		return nil, err
	}

	if reserved != 0 || typ != 1 || count == 0 {
		return nil, types.ErrUnknownFormat
	}

	var entry [16]byte

	_, err = io.ReadFull(r, entry[:])
	if err != nil {
		return nil, err
	}

	w := int(entry[0])
	h := int(entry[1])

	if w == 0 {
		w = 256
	}

	if h == 0 {
		h = 256
	}

	return &types.Metadata{
		Name: "Windows icon",
		Info: fmt.Sprintf("%dx%d, %d images", w, h, count),
	}, nil
}
