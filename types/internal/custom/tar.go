package custom

import "github.com/coalaura/onda/types"

func DetectTar(b types.Buffer) *types.Metadata {
	if b.Has(257, []byte("ustar\x00")) || b.Has(257, []byte("ustar  ")) {
		return &types.Metadata{
			Kind: types.KindTARArchive,
		}
	}

	return nil
}
