package custom

import "github.com/coalaura/onda/types"

func DetectMetafile(b types.Buffer) *types.Metadata {
	if b.Has(0, []byte{0x01, 0x00, 0x00, 0x00}) {
		return &types.Metadata{
			Name: "Enhanced Metafile image",
		}
	}

	if b.Has(0, []byte{0xd7, 0xcd, 0xc6, 0x9a}) {
		return &types.Metadata{
			Name: "Windows Metafile image",
		}
	}

	return nil
}
