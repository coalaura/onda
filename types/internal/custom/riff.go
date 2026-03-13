package custom

import "github.com/coalaura/onda/types"

func DetectRIFF(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte("RIFF")) {
		return nil
	}

	if b.Has(8, []byte("WAVE")) {
		return &types.Metadata{
			Name: "WAV audio",
		}
	}

	if b.Has(8, []byte("AVI ")) {
		return &types.Metadata{
			Name: "AVI video",
		}
	}

	if b.Has(8, []byte("WEBP")) {
		return &types.Metadata{
			Name: "WebP image",
		}
	}

	return &types.Metadata{
		Name: "RIFF container",
	}
}
