package macho

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
	var magic uint32

	err := binary.Read(r, binary.BigEndian, &magic)
	if err != nil {
		return nil, err
	}

	var (
		order binary.ByteOrder
		bits  string
	)

	switch magic {
	case 0xfeedface:
		order = binary.BigEndian
		bits = "Mach-O 32"
	case 0xfeedfacf:
		order = binary.BigEndian
		bits = "Mach-O 64"
	case 0xcefaedfe:
		order = binary.LittleEndian
		bits = "Mach-O 32"
	case 0xcffaedfe:
		order = binary.LittleEndian
		bits = "Mach-O 64"
	default:
		return nil, types.ErrUnknownFormat
	}

	var cpu uint32

	err = binary.Read(r, order, &cpu)
	if err != nil {
		return nil, err
	}

	return &types.Metadata{
		Name: "Mach-O",
		Type: bits,
		Info: cpuName(cpu),
	}, nil
}

func cpuName(c uint32) string {
	switch c {
	case 7:
		return "x86"
	case 0x01000007:
		return "x86-64"
	case 12:
		return "ARM"
	case 0x0100000c:
		return "ARM64"
	default:
		return fmt.Sprintf("cpu 0x%x", c)
	}
}
