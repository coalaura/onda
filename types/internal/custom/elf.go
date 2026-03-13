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

	var typ string

	switch class {
	case 1:
		typ = "ELF32"
	case 2:
		typ = "ELF64"
	default:
		typ = "ELF"
	}

	switch data {
	case 1:
		typ += " Little-Endian"
	case 2:
		typ += " Big-Endian"
	}

	return &types.Metadata{
		Name: "Executable and Linkable Format",
		Type: typ,
	}
}
