package custom

import "github.com/coalaura/onda/types"

func DetectTorrent(b types.Buffer) *types.Metadata {
	if b.Len() < 32 {
		return nil
	}

	if b[0] != 'd' {
		return nil
	}

	if !containsASCII(b, "8:announce") || !containsASCII(b, "4:info") {
		return nil
	}

	return &types.Metadata{Kind: types.KindTorrentFile}
}

func containsASCII(b types.Buffer, s string) bool {
	if len(s) == 0 || b.Len() < len(s) {
		return false
	}

	limit := min(b.Len(), 4096) - len(s)
	for i := 0; i <= limit; i++ {
		if b.Has(i, []byte(s)) {
			return true
		}
	}

	return false
}
