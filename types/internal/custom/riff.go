package custom

import "github.com/coalaura/onda/types"

func DetectRIFF(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte("RIFF")) {
		return nil
	}

	if b.Has(8, []byte("WAVE")) {
		return &types.Metadata{
			Kind: types.KindRIFFContainer,
			Type: types.TypeWAVAudio,
		}
	}

	if b.Has(8, []byte("AVI ")) {
		return &types.Metadata{
			Kind: types.KindRIFFContainer,
			Type: types.TypeAVIVideo,
		}
	}

	if b.Has(8, []byte("WEBP")) {
		return &types.Metadata{
			Kind: types.KindRIFFContainer,
			Type: types.TypeWebPImage,
		}
	}

	return &types.Metadata{
		Kind: types.KindRIFFContainer,
	}
}
