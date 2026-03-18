package custom

import "github.com/coalaura/onda/types"

func DetectFAT(b types.Buffer) *types.Metadata {
	if b.Len() < 90 {
		return nil
	}

	if !((b[0] == 0xeb && b[2] == 0x90) || b[0] == 0xe9) {
		return nil
	}

	bps, ok := b.U16LE(11) // Bytes per sector
	if !ok || (bps != 512 && bps != 1024 && bps != 2048 && bps != 4096) {
		return nil
	}

	spc := b[13]
	if spc == 0 || (spc&(spc-1)) != 0 {
		return nil
	}

	fats := b[16]
	if fats == 0 || fats > 4 {
		return nil
	}

	if b.Has(82, []byte("FAT32   ")) {
		return &types.Metadata{Kind: types.KindFATFilesystem, Type: types.TypeFAT32}
	}

	if b.Has(54, []byte("FAT12   ")) {
		return &types.Metadata{Kind: types.KindFATFilesystem, Type: types.TypeFAT12}
	}

	if b.Has(54, []byte("FAT16   ")) {
		return &types.Metadata{Kind: types.KindFATFilesystem, Type: types.TypeFAT16}
	}

	return nil
}
