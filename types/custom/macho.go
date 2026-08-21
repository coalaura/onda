package custom

import "github.com/coalaura/wtf/types"

func DetectMachO(b types.Buffer) *types.Metadata {
	if b.Has(0, []byte{0xfe, 0xed, 0xfa, 0xce}) {
		if !isValidMachOHeaderBE(b, 4) {
			return nil
		}

		return &types.Metadata{
			Kind: types.KindMachOBinary,
			Type: types.Type32BitBigEndian,
		}
	}

	if b.Has(0, []byte{0xce, 0xfa, 0xed, 0xfe}) {
		if !isValidMachOHeaderLE(b, 4) {
			return nil
		}

		return &types.Metadata{
			Kind: types.KindMachOBinary,
			Type: types.Type32BitLittleEndian,
		}
	}

	if b.Has(0, []byte{0xfe, 0xed, 0xfa, 0xcf}) {
		if !isValidMachOHeaderBE(b, 4) {
			return nil
		}

		return &types.Metadata{
			Kind: types.KindMachOBinary,
			Type: types.Type64BitBigEndian,
		}
	}

	if b.Has(0, []byte{0xcf, 0xfa, 0xed, 0xfe}) {
		if !isValidMachOHeaderLE(b, 4) {
			return nil
		}

		return &types.Metadata{
			Kind: types.KindMachOBinary,
			Type: types.Type64BitLittleEndian,
		}
	}

	if b.Has(0, []byte{0xca, 0xfe, 0xba, 0xbe}) {
		// Java class files: minor version at offset 4-5, major version at offset 6-7
		// Valid major version range: 45 (Java 1.1) to ~75 (Java 21+)
		if b.Len() >= 8 {
			majorVersion, ok := b.U16BE(6)
			if ok && majorVersion >= 45 && majorVersion <= 75 {
				return &types.Metadata{Kind: types.KindJavaClassBytecode}
			}
		}

		if !isValidFatMachO(b, 20) {
			return nil
		}

		return &types.Metadata{
			Kind: types.KindMachOUniversalBinary,
		}
	}

	if b.Has(0, []byte{0xca, 0xfe, 0xba, 0xbf}) {
		if !isValidFatMachO(b, 32) {
			return nil
		}

		return &types.Metadata{
			Kind: types.KindMachOUniversalBinary,
			Type: types.Type64Bit,
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

func isValidFatMachO(b types.Buffer, archSize int) bool {
	nfatArch, ok := b.U32BE(4)
	if !ok || nfatArch == 0 || nfatArch > 32 || 8+int(nfatArch)*archSize > b.Len() {
		return false
	}

	cpuType, _ := b.U32BE(8)

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
