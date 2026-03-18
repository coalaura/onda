package custom

import "github.com/coalaura/onda/types"

func DetectAppleDiskImage(b types.Buffer) *types.Metadata {
	if b.Len() < 512 {
		return nil
	}

	if b.Has(b.Len()-512, []byte("koly")) {
		return &types.Metadata{
			Kind: types.KindAppleDiskImage,
		}
	}

	return nil
}
