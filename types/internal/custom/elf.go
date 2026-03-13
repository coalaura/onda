package custom

import "github.com/coalaura/onda/types"

func DetectELF(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte{0x7f, 'E', 'L', 'F'}) {
		return nil
	}

	class := byte(0)
	data := byte(0)

	if b.Len() > 4 {
		class = b[4]
	}

	if b.Len() > 5 {
		data = b[5]
	}

	typ := elfType(class, data)

	return &types.Metadata{
		Kind: types.KindExecutableAndLinkableFormat,
		Type: typ,
	}
}

func elfType(class byte, data byte) types.TypeID {
	switch class {
	case 1:
		switch data {
		case 1:
			return types.TypeELF32LittleEndian
		case 2:
			return types.TypeELF32BigEndian
		default:
			return types.TypeELF32
		}
	case 2:
		switch data {
		case 1:
			return types.TypeELF64LittleEndian
		case 2:
			return types.TypeELF64BigEndian
		default:
			return types.TypeELF64
		}
	default:
		return types.TypeELF
	}
}
