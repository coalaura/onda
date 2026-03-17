package custom

import "github.com/coalaura/onda/types"

func DetectLZMA(b types.Buffer) *types.Metadata {
	if b.Len() < 13 {
		return nil
	}

	props := b[0]
	if props > 224 {
		return nil
	}

	dictSize, ok := b.U32LE(1)
	if !ok {
		return nil
	}

	if dictSize == 0 {
		return nil
	}

	if dictSize > (1 << 30) {
		return nil
	}

	return &types.Metadata{
		Kind: types.KindLZMAData,
	}
}
