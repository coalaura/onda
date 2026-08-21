package custom

import "github.com/coalaura/wtf/types"

func DetectMachO(b types.Buffer) *types.Metadata {
	if b.Has(0, []byte{0xfe, 0xed, 0xfa, 0xce}) {
		if !isValidMachOHeaderBE(b, 28) {
			return nil
		}

		return &types.Metadata{
			Kind: types.KindMachOBinary,
			Type: types.Type32BitBigEndian,
		}
	}

	if b.Has(0, []byte{0xce, 0xfa, 0xed, 0xfe}) {
		if !isValidMachOHeaderLE(b, 28) {
			return nil
		}

		return &types.Metadata{
			Kind: types.KindMachOBinary,
			Type: types.Type32BitLittleEndian,
		}
	}

	if b.Has(0, []byte{0xfe, 0xed, 0xfa, 0xcf}) {
		if !isValidMachOHeaderBE(b, 32) {
			return nil
		}

		return &types.Metadata{
			Kind: types.KindMachOBinary,
			Type: types.Type64BitBigEndian,
		}
	}

	if b.Has(0, []byte{0xcf, 0xfa, 0xed, 0xfe}) {
		if !isValidMachOHeaderLE(b, 32) {
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

func isValidMachOHeaderBE(b types.Buffer, headerSize int) bool {
	if b.Len() < headerSize {
		return false
	}

	cpuType, _ := b.U32BE(4)
	ncmds, _ := b.U32BE(16)
	sizeofcmds, _ := b.U32BE(20)

	return isKnownMachOCPUType(cpuType) && ncmds <= 4096 && uint64(headerSize)+uint64(sizeofcmds) <= uint64(b.Len())
}

func isValidMachOHeaderLE(b types.Buffer, headerSize int) bool {
	if b.Len() < headerSize {
		return false
	}

	cpuType, _ := b.U32LE(4)
	ncmds, _ := b.U32LE(16)
	sizeofcmds, _ := b.U32LE(20)

	return isKnownMachOCPUType(cpuType) && ncmds <= 4096 && uint64(headerSize)+uint64(sizeofcmds) <= uint64(b.Len())
}

func isValidFatMachO(b types.Buffer, archSize int) bool {
	nfatArch, ok := b.U32BE(4)
	if !ok || nfatArch == 0 || nfatArch > 32 || 8+int(nfatArch)*archSize > b.Len() {
		return false
	}

	tableEnd := uint64(8) + uint64(nfatArch)*uint64(archSize)

	for i := range int(nfatArch) {
		offset := 8 + i*archSize

		cpuType, _ := b.U32BE(offset)
		if !isKnownMachOCPUType(cpuType) {
			return false
		}

		if archSize == 20 {
			archOffset, _ := b.U32BE(offset + 8)
			archLength, _ := b.U32BE(offset + 12)

			if uint64(archOffset) < tableEnd || archLength == 0 {
				return false
			}

			continue
		}

		offsetHigh, _ := b.U32BE(offset + 8)
		offsetLow, _ := b.U32BE(offset + 12)

		lengthHigh, _ := b.U32BE(offset + 16)
		lengthLow, _ := b.U32BE(offset + 20)

		archOffset := uint64(offsetHigh)<<32 | uint64(offsetLow)
		archLength := uint64(lengthHigh)<<32 | uint64(lengthLow)

		if archOffset < tableEnd || archLength == 0 {
			return false
		}
	}

	return true
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
