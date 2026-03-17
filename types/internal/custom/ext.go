package custom

import "github.com/coalaura/onda/types"

const (
	superblockOffset = 1024
	magicOffset      = superblockOffset + 56
	compatOffset     = superblockOffset + 92
	incompatOffset   = superblockOffset + 96
	roCompatOffset   = superblockOffset + 100
)

func DetectExt(b types.Buffer) *types.Metadata {
	if b.Len() < superblockOffset+104 {
		return nil
	}

	magic, ok := b.U16LE(magicOffset)
	if !ok || magic != 0xef53 {
		return nil
	}

	compat, ok := b.U32LE(compatOffset)
	if !ok {
		return nil
	}

	incompat, ok := b.U32LE(incompatOffset)
	if !ok {
		return nil
	}

	roCompat, ok := b.U32LE(roCompatOffset)
	if !ok {
		return nil
	}

	typ := types.TypeExt2

	if incompat&0x40 != 0 {
		typ = types.TypeExt4
	} else if compat&0x4 != 0 || roCompat&0x4 != 0 {
		typ = types.TypeExt3
	}

	return &types.Metadata{
		Kind: types.KindExtFilesystem,
		Type: typ,
	}
}
