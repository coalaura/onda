package custom

import "github.com/coalaura/onda/types"

func DetectGzipData(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte{0x1f, 0x8b}) {
		return nil
	}

	if b.Len() >= 16 && b.Has(12, []byte{'B', 'C', 0x02, 0x00}) {
		return &types.Metadata{Kind: types.KindGzipData, Type: types.TypeBGZF}
	}

	return nil
}
