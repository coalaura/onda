package pe

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
	var mz [2]byte

	_, err := io.ReadFull(r, mz[:])
	if err != nil {
		return nil, err
	}

	if mz != [2]byte{'M', 'Z'} {
		return nil, types.ErrUnknownFormat
	}

	_, err = r.Seek(0x3c, io.SeekStart)
	if err != nil {
		return nil, err
	}

	var peOff uint32

	err = binary.Read(r, binary.LittleEndian, &peOff)
	if err != nil {
		return nil, err
	}

	_, err = r.Seek(int64(peOff), io.SeekStart)
	if err != nil {
		return nil, err
	}

	var sig [4]byte

	_, err = io.ReadFull(r, sig[:])
	if err != nil {
		return nil, err
	}

	if sig != [4]byte{'P', 'E', 0, 0} {
		return nil, types.ErrUnknownFormat
	}

	var machine uint16

	err = binary.Read(r, binary.LittleEndian, &machine)
	if err != nil {
		return nil, err
	}

	_, err = r.Seek(2+4+4+4, io.SeekCurrent)
	if err != nil {
		return nil, err
	}

	var sizeOfOptionalHeader uint16

	err = binary.Read(r, binary.LittleEndian, &sizeOfOptionalHeader)
	if err != nil {
		return nil, err
	}

	var characteristics uint16

	err = binary.Read(r, binary.LittleEndian, &characteristics)
	if err != nil {
		return nil, err
	}

	var magic uint16

	err = binary.Read(r, binary.LittleEndian, &magic)
	if err != nil {
		return nil, err
	}

	var (
		peType          string
		subsystemOffset int64
	)

	switch magic {
	case 0x10b:
		peType = "PE32"
		subsystemOffset = 68
	case 0x20b:
		peType = "PE32+"
		subsystemOffset = 88
	default:
		return nil, types.ErrUnknownFormat
	}

	role := "executable"

	if characteristics&0x2000 != 0 {
		role = "DLL"
	} else {
		var subsystem uint16

		if sizeOfOptionalHeader >= uint16(subsystemOffset+2) {
			_, err := r.Seek(subsystemOffset-2, io.SeekCurrent)
			if err == nil {
				binary.Read(r, binary.LittleEndian, &subsystem)

				if subsystem == 1 {
					role = "system driver"
				}
			}
		}
	}

	return &types.Metadata{
		Name: "Portable Executable",
		Type: fmt.Sprintf("%s %s", peType, role),
		Info: machineName(machine),
	}, nil
}

func machineName(machine uint16) string {
	switch machine {
	case 0x014c:
		return "x86"
	case 0x8664:
		return "x86-64"
	case 0x01c0:
		return "ARM"
	case 0x01c4:
		return "ARMv7"
	case 0xaa64:
		return "ARM64"
	case 0x0200:
		return "Intel Itanium"
	default:
		return fmt.Sprintf("machine 0x%04x", machine)
	}
}
