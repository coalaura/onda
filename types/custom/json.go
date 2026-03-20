package custom

import (
	"bytes"

	"github.com/coalaura/onda/types"
)

func DetectJSON(b types.Buffer) *types.Metadata {
	limit := min(b.Len(), 4096)
	data := b[:limit]

	trimmed := bytes.TrimSpace(data)
	if len(trimmed) < 2 {
		return nil
	}

	if trimmed[0] == '{' || trimmed[0] == '[' {
		if bytes.Contains(trimmed, []byte(`":`)) || bytes.Contains(trimmed, []byte(`": `)) {
			return &types.Metadata{Kind: types.KindJSONDocument}
		}
	}

	return nil
}
