package custom

import "github.com/coalaura/onda/types"

func DetectPE(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte{'M', 'Z'}) {
		return nil
	}

	peOff, ok := b.U32LE(0x3c)
	if !ok {
		return nil
	}

	off := int(peOff)

	if !b.Has(off, []byte{'P', 'E', 0, 0}) {
		return nil
	}

	machine, ok := b.U16LE(off + 4)
	if !ok {
		return nil
	}

	magic, ok := b.U16LE(off + 24)
	if !ok {
		return nil
	}

	var typ string

	switch magic {
	case 0x10b:
		typ = "PE32"
	case 0x20b:
		typ = "PE32+"
	default:
		return nil
	}

	return &types.Metadata{
		Name: "Portable Executable",
		Type: typ + " " + peMachine(machine),
	}
}

func peMachine(machine uint16) string {
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
		return "Itanium"
	default:
		return "unknown"
	}
}
