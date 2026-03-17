package custom

import (
	"bytes"

	"github.com/coalaura/onda/types"
)

func DetectTorrent(b types.Buffer) *types.Metadata {
	if b.Len() < 32 || b[0] != 'd' {
		return nil
	}

	limit := min(b.Len(), 4096)
	data := b[:limit]

	if !bytes.Contains(data, []byte("8:announce")) || !bytes.Contains(data, []byte("4:info")) {
		return nil
	}

	return &types.Metadata{Kind: types.KindTorrentFile}
}
