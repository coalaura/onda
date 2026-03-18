package custom

import (
	"bytes"

	"github.com/coalaura/onda/types"
)

func DetectEBML(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte{0x1a, 0x45, 0xdf, 0xa3}) {
		return nil
	}

	limit := min(b.Len(), 256)
	data := b[:limit]

	if bytes.Contains(data, []byte("webm")) {
		return &types.Metadata{
			Kind: types.KindEBMLContainer,
			Type: types.TypeWebM,
		}
	}

	if bytes.Contains(data, []byte("matroska")) {
		return &types.Metadata{
			Kind: types.KindEBMLContainer,
			Type: types.TypeMatroska,
		}
	}

	return &types.Metadata{
		Kind: types.KindEBMLContainer,
	}
}
