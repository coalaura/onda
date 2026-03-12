package elf

import (
	"fmt"
	"io"
	"unsafe"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var ident [16]byte

	_, err := io.ReadFull(r, ident[:])
	if err != nil {
		return nil, err
	}

	if ident[0] != 0x7f || ident[1] != 'E' || ident[2] != 'L' || ident[3] != 'F' {
		return nil, types.ErrUnknownFormat
	}

	class := ident[4]
	machineOff := int64(18)

	_, err = r.Seek(machineOff, io.SeekStart)
	if err != nil {
		return nil, err
	}

	var machine uint16

	_, err = io.ReadFull(r, (*[2]byte)(unsafe.Pointer(&machine))[:])
	if err != nil {
		return nil, err
	}

	var bits string

	switch class {
	case 1:
		bits = "ELF32"
	case 2:
		bits = "ELF64"
	default:
		bits = "ELF"
	}

	return &types.Metadata{
		Name: "ELF Executable",
		Type: bits,
		Info: machineName(machine),
	}, nil
}

func machineName(m uint16) string {
	switch m {
	case 0x03:
		return "x86"
	case 0x3e:
		return "x86-64"
	case 0xb7:
		return "ARM64"
	case 0x28:
		return "ARM"
	default:
		return fmt.Sprintf("machine 0x%04x", m)
	}
}
