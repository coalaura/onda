package custom

import "github.com/coalaura/onda/types"

func DetectOLE(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte{0xd0, 0xcf, 0x11, 0xe0, 0xa1, 0xb1, 0x1a, 0xe1}) {
		return nil
	}

	return &types.Metadata{
		Kind: types.KindOLECompoundDocument,
	}
}
