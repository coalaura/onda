package custom

import "github.com/coalaura/onda/types"

func DetectMachO(b types.Buffer) *types.Metadata {
	if b.Has(0, []byte{0xfe, 0xed, 0xfa, 0xce}) {
		if !isValidMachOHeaderBE(b, 4) {
			return nil
		}

		return &types.Metadata{
			Name: "Mach-O binary",
			Type: "32-bit big-endian",
		}
	}

	if b.Has(0, []byte{0xce, 0xfa, 0xed, 0xfe}) {
		if !isValidMachOHeaderLE(b, 4) {
			return nil
		}

		return &types.Metadata{
			Name: "Mach-O binary",
			Type: "32-bit little-endian",
		}
	}

	if b.Has(0, []byte{0xfe, 0xed, 0xfa, 0xcf}) {
		if !isValidMachOHeaderBE(b, 4) {
			return nil
		}

		return &types.Metadata{
			Name: "Mach-O binary",
			Type: "64-bit big-endian",
		}
	}

	if b.Has(0, []byte{0xcf, 0xfa, 0xed, 0xfe}) {
		if !isValidMachOHeaderLE(b, 4) {
			return nil
		}

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

		if nfatArch == 0 || nfatArch > 32 {
			return nil
		}

		if !isValidFatArchCPUType(b, 8) {
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

		if nfatArch == 0 || nfatArch > 32 {
			return nil
		}

		if !isValidFatArchCPUType(b, 8) {
			return nil
		}

		return &types.Metadata{
			Name: "Mach-O universal binary",
			Type: "64-bit",
		}
	}

	return nil
}

func isValidMachOHeaderBE(b types.Buffer, cpuOffset int) bool {
	cpuType, ok := b.U32BE(cpuOffset)
	if !ok {
		return false
	}

	return isKnownMachOCPUType(cpuType)
}

func isValidMachOHeaderLE(b types.Buffer, cpuOffset int) bool {
	cpuType, ok := b.U32LE(cpuOffset)
	if !ok {
		return false
	}

	return isKnownMachOCPUType(cpuType)
}

func isValidFatArchCPUType(b types.Buffer, cpuOffset int) bool {
	cpuType, ok := b.U32BE(cpuOffset)
	if !ok {
		return false
	}

	return isKnownMachOCPUType(cpuType)
}

func isKnownMachOCPUType(cpuType uint32) bool {
	masked := cpuType & 0x00ffffff

	switch masked {
	case 6, 7, 12, 14, 18:
		return true
	default:
		return false
	}
}
