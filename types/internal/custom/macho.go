package custom

import "github.com/coalaura/onda/types"

func DetectMachO(b types.Buffer) *types.Metadata {
	if b.Has(0, []byte{0xfe, 0xed, 0xfa, 0xce}) {
		return &types.Metadata{
			Name: "Mach-O binary",
			Type: "32-bit big-endian",
		}
	}

	if b.Has(0, []byte{0xce, 0xfa, 0xed, 0xfe}) {
		return &types.Metadata{
			Name: "Mach-O binary",
			Type: "32-bit little-endian",
		}
	}

	if b.Has(0, []byte{0xfe, 0xed, 0xfa, 0xcf}) {
		return &types.Metadata{
			Name: "Mach-O binary",
			Type: "64-bit big-endian",
		}
	}

	if b.Has(0, []byte{0xcf, 0xfa, 0xed, 0xfe}) {
		return &types.Metadata{
			Name: "Mach-O binary",
			Type: "64-bit little-endian",
		}
	}

	if b.Has(0, []byte{0xca, 0xfe, 0xba, 0xbe}) {
		nfatArch, ok := b.U32BE(4)
		if !ok {
			return nil
		}

		if nfatArch == 0 || nfatArch > 64 {
			return nil
		}

		return &types.Metadata{
			Name: "Mach-O universal binary",
		}
	}

	if b.Has(0, []byte{0xca, 0xfe, 0xba, 0xbf}) {
		nfatArch, ok := b.U32BE(4)
		if !ok {
			return nil
		}

		if nfatArch == 0 || nfatArch > 64 {
			return nil
		}

		return &types.Metadata{
			Name: "Mach-O universal binary",
			Type: "64-bit",
		}
	}

	return nil
}
