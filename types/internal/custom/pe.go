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

	var typ types.TypeID

	switch magic {
	case 0x10b:
		typ = pe32MachineType(machine)
	case 0x20b:
		typ = pe32PlusMachineType(machine)
	default:
		return nil
	}

	return &types.Metadata{
		Kind: types.KindPortableExecutable,
		Type: typ,
	}
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
