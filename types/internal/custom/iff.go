package custom

import "github.com/coalaura/onda/types"

func DetectIFF(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte("FORM")) {
		return nil
	}

	if b.Has(8, []byte("AIFF")) {
		return &types.Metadata{
			Name: "AIFF audio",
		}
	}

	if b.Has(8, []byte("AIFC")) {
		return &types.Metadata{
			Name: "AIFF audio",
			Type: "AIFC",
		}
	}

	return &types.Metadata{
		Name: "IFF container",
	}
}
