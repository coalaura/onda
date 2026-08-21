package custom

import "github.com/coalaura/wtf/types"

func Detect3DStudio(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte{0x4d, 0x4d}) {
		return nil
	}

	chunkSize, ok := b.U32LE(2)
	if !ok || chunkSize < 12 {
		return nil
	}

	child, ok := b.U16LE(6)
	if !ok || child != 0x0002 && child != 0x3d3d && child != 0xb000 {
		return nil
	}

	return &types.Metadata{Kind: types.Kind3DStudioMaxModel}
}
