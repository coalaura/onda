package custom

import "github.com/coalaura/onda/types"

func DetectCRX(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte("Cr24")) {
		return nil
	}

	ver, ok := b.U32LE(4)
	if !ok {
		return &types.Metadata{Kind: types.KindCRXBrowserExtension}
	}

	switch ver {
	case 2:
		return &types.Metadata{Kind: types.KindCRXBrowserExtension, Type: types.TypeCRXVersion2}
	case 3:
		return &types.Metadata{Kind: types.KindCRXBrowserExtension, Type: types.TypeCRXVersion3}
	default:
		return &types.Metadata{Kind: types.KindCRXBrowserExtension}
	}
}
