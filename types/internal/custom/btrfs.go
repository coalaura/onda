package custom

import "github.com/coalaura/onda/types"

func DetectBtrfs(b types.Buffer) *types.Metadata {
	if !b.Has(0x10040, []byte("_BHRfS_M")) {
		return nil
	}

	return &types.Metadata{
		Kind: types.KindBtrfsFilesystem,
	}
}
