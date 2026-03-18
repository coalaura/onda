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
		return &types.Metadata{Kind: types.KindOggContainer, Type: types.TypeOpusAudio}
	case bytes.Contains(data, []byte{0x01, 'v', 'o', 'r', 'b', 'i', 's'}):
		return &types.Metadata{Kind: types.KindOggContainer, Type: types.TypeVorbisAudio}
	case bytes.Contains(data, []byte("Speex   ")):
		return &types.Metadata{Kind: types.KindOggContainer, Type: types.TypeSpeexAudio}
	case bytes.Contains(data, []byte{0x80, 't', 'h', 'e', 'o', 'r', 'a'}):
		return &types.Metadata{Kind: types.KindOggContainer, Type: types.TypeTheoraVideo}
	case bytes.Contains(data, []byte{0x7f, 'F', 'L', 'A', 'C'}):
		return &types.Metadata{Kind: types.KindOggContainer, Type: types.TypeFLACAudio}
	default:
		return &types.Metadata{Kind: types.KindOggContainer}
	}
}
