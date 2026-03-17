package custom

import (
	"bytes"

	"github.com/coalaura/onda/types"
)

func DetectPKCS12(b types.Buffer) *types.Metadata {
	if b.Len() < 32 || b[0] != 0x30 {
		return nil
	}

	limit := min(b.Len(), 4096)
	if !bytes.Contains(b[:limit], []byte{0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x01, 0x07, 0x01}) {
		return nil
	}

	return &types.Metadata{Kind: types.KindPKCS12}
}
