package custom

import (
	"bytes"

	"github.com/coalaura/onda/types"
)

func DetectOgg(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte("OggS")) {
		return nil
	}

	limit := min(b.Len(), 4096)
	data := b[:limit]

	switch {
	case bytes.Contains(data, []byte("OpusHead")):
		return &types.Metadata{Name: "Opus audio"}
	case bytes.Contains(data, []byte{0x01, 'v', 'o', 'r', 'b', 'i', 's'}):
		return &types.Metadata{Name: "Vorbis audio"}
	case bytes.Contains(data, []byte("Speex   ")):
		return &types.Metadata{Name: "Speex audio"}
	case bytes.Contains(data, []byte{0x80, 't', 'h', 'e', 'o', 'r', 'a'}):
		return &types.Metadata{Name: "Theora video"}
	case bytes.Contains(data, []byte{0x7f, 'F', 'L', 'A', 'C'}):
		return &types.Metadata{Name: "FLAC audio", Type: "Ogg container"}
	default:
		return &types.Metadata{Name: "Ogg container"}
	}
}
