package custom

import (
	"bytes"

	"github.com/coalaura/onda/types"
)

func DetectSVG(b types.Buffer) *types.Metadata {
	if b.Len() < 4 {
		return nil
	}

	limit := min(b.Len(), 4096)
	data := b[:limit]

	if !bytes.Contains(data, []byte("<svg")) {
		return nil
	}

	trimmed := bytes.TrimSpace(data)
	if len(trimmed) > 0 && trimmed[0] == '<' {
		return &types.Metadata{Kind: types.KindSVGImage}
	}

	return nil
}
