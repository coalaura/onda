package ext

import (
	"encoding/binary"
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	_, err := r.Seek(1024+56, io.SeekStart)
	if err != nil {
		return nil, err
	}

	var magic uint16

	err = binary.Read(r, binary.LittleEndian, &magic)
	if err != nil {
		return nil, err
	}

	if magic != 0xEF53 {
		return nil, types.ErrUnknownFormat
	}

	_, err = r.Seek(1024+24, io.SeekStart)
	if err != nil {
		return nil, err
	}

	var rev uint32

	err = binary.Read(r, binary.LittleEndian, &rev)
	if err != nil {
		return nil, err
	}

	var name string

	if rev == 0 {
		name = "ext2 filesystem"
	} else {
		name = "ext filesystem"
	}

	return &types.Metadata{
		Name: name,
		Info: "superblock present",
	}, nil
}
