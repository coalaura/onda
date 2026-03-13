package custom

import "github.com/coalaura/onda/types"

func DetectRIFF(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte("RIFF")) {
		return nil
	}

	if b.Has(8, []byte("WAVE")) {
		return &types.Metadata{
			Name: "RIFF Container",
			Type: "WAV Audio",
		}
	}

	if b.Has(8, []byte("AVI ")) {
		return &types.Metadata{
			Name: "RIFF Container",
			Type: "AVI Video",
		}
	}

	if b.Has(8, []byte("WEBP")) {
		return &types.Metadata{
			Name: "RIFF Container",
			Type: "WebP Image",
		}
	}

	return &types.Metadata{
		Name: "RIFF Container",
	}
}
