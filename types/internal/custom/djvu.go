package custom

import "github.com/coalaura/onda/types"

func DetectDjVu(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte("AT&TFORM")) {
		return nil
	}

	if b.Len() < 16 {
		return nil
	}

	switch {
	case b.Has(12, []byte("DJVU")), b.Has(12, []byte("DJVM")), b.Has(12, []byte("DJVI")), b.Has(12, []byte("THUM")):
		return &types.Metadata{
			Kind: types.KindDjVuDocument,
		}
	default:
		return nil
	}
}
