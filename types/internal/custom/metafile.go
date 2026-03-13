package custom

import "github.com/coalaura/onda/types"

func DetectMetafile(b types.Buffer) *types.Metadata {
	if isEMF(b) {
		return &types.Metadata{
			Kind: types.KindMetafileImage,
			Type: types.TypeEnhancedMetafileEMF,
		}
	}

	if b.Has(0, []byte{0xd7, 0xcd, 0xc6, 0x9a}) {
		return &types.Metadata{
			Kind: types.KindMetafileImage,
			Type: types.TypeWindowsMetafileWMF,
		}
	}

	return nil
}

func isEMF(b types.Buffer) bool {
	if !b.Has(0, []byte{0x01, 0x00, 0x00, 0x00}) {
		return false
	}

	size, ok := b.U32LE(4)
	if !ok || size < 88 {
		return false
	}

	return b.Has(40, []byte{0x20, 0x45, 0x4d, 0x46})
}
