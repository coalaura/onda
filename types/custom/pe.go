package custom

import "github.com/coalaura/wtf/types"

func DetectPE(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte{'M', 'Z'}) {
		return nil
	}

	if !isValidDOSHeader(b) {
		return nil
	}

	peOff, ok := b.U32LE(0x3c)
	if !ok {
		return &types.Metadata{Kind: types.KindDOSExecutable}
	}

	off := int(peOff)

	if b.Has(off, []byte{'P', 'E', 0, 0}) {
		machine, ok := b.U16LE(off + 4)
		if !ok {
			return &types.Metadata{Kind: types.KindPortableExecutable}
		}

		magic, ok := b.U16LE(off + 24)
		if !ok {
			return &types.Metadata{Kind: types.KindPortableExecutable}
		}

		var typ types.TypeID

		switch magic {
		case 0x10b:
			typ = pe32MachineType(machine)
		case 0x20b:
			typ = pe32PlusMachineType(machine)
		default:
			return &types.Metadata{Kind: types.KindPortableExecutable}
		}

		return &types.Metadata{
			Kind: types.KindPortableExecutable,
			Type: typ,
		}
	}

	if b.Has(off, []byte{'N', 'E'}) {
		return &types.Metadata{Kind: types.KindPortableExecutable, Type: types.TypeWindowsNE}
	}

	if b.Has(off, []byte{'L', 'E'}) {
		return &types.Metadata{Kind: types.KindPortableExecutable, Type: types.TypeWindowsLE}
	}

	if b.Has(off, []byte{'L', 'X'}) {
		return &types.Metadata{Kind: types.KindPortableExecutable, Type: types.TypeWindowsLX}
	}

	return &types.Metadata{Kind: types.KindDOSExecutable}
}

func isValidDOSHeader(b types.Buffer) bool {
	if b.Len() < 28 {
		return false
	}

	lastPage, ok := b.U16LE(2)
	if !ok || lastPage > 512 {
		return false
	}

	pages, ok := b.U16LE(4)
	if !ok || pages == 0 {
		return false
	}

	headerParagraphs, ok := b.U16LE(8)
	if !ok || headerParagraphs < 2 {
		return false
	}

	relocationOffset, ok := b.U16LE(24)
	return ok && relocationOffset >= 0x1c
}

func pe32MachineType(machine uint16) types.TypeID {
	switch machine {
	case 0x014c:
		return types.TypePE32X86
	case 0x8664:
		return types.TypePE32X8664
	case 0x01c0:
		return types.TypePE32ARM
	case 0x01c4:
		return types.TypePE32ARMv7
	case 0xaa64:
		return types.TypePE32ARM64
	case 0x0200:
		return types.TypePE32Itanium
	default:
		return types.TypePE32Unknown
	}
}

func pe32PlusMachineType(machine uint16) types.TypeID {
	switch machine {
	case 0x8664:
		return types.TypePE32PlusX8664
	case 0xaa64:
		return types.TypePE32PlusARM64
	default:
		return types.TypePE32PlusUnknown
	}
}
