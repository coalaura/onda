package custom

import "github.com/coalaura/onda/types"

func DetectDjVu(b types.Buffer) *types.Metadata {
	if b.Has(0, []byte("FORM")) {
		switch {
		case b.Has(8, []byte("DJVU")):
			return &types.Metadata{Kind: types.KindDjVuDocument}
		case b.Has(8, []byte("DJVM")):
			return &types.Metadata{Kind: types.KindDjVuDocument}
		case b.Has(8, []byte("DJVI")):
			return &types.Metadata{Kind: types.KindDjVuDocument}
		case b.Has(8, []byte("THUM")):
			return &types.Metadata{Kind: types.KindDjVuDocument}
		}
	}

	if b.Has(0, []byte("AT&TFORM")) {
		if b.Len() < 16 {
			return nil
		}

		switch {
		case b.Has(12, []byte("DJVU")), b.Has(12, []byte("DJVM")), b.Has(12, []byte("DJVI")), b.Has(12, []byte("THUM")):
			return &types.Metadata{Kind: types.KindDjVuDocument}
		}
	}

	return nil
}
