package custom

import "github.com/coalaura/onda/types"

func DetectEBML(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte{0x1a, 0x45, 0xdf, 0xa3}) {
		return nil
	}

	limit := min(b.Len(), 256)

	for i := 4; i+4 <= limit; i++ {
		if b.Has(i, []byte("webm")) {
			return &types.Metadata{
				Name: "EBML Container",
				Type: "WebM",
			}
		}

		if b.Has(i, []byte("matroska")) {
			return &types.Metadata{
				Name: "EBML Container",
				Type: "Matroska",
			}
		}
	}

	return &types.Metadata{
		Name: "EBML Container",
	}
}
