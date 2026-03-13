package custom

import "github.com/coalaura/onda/types"

func DetectMPEGTransport(b types.Buffer) *types.Metadata {
	if b.Len() >= 3*188 && b[0] == 0x47 && b[188] == 0x47 && b[376] == 0x47 {
		return &types.Metadata{Name: "MPEG transport stream", Type: "TS"}
	}

	if b.Len() >= 4+3*192 && b[4] == 0x47 && b[196] == 0x47 && b[388] == 0x47 {
		return &types.Metadata{Name: "MPEG transport stream", Type: "M2TS"}
	}

	return nil
}
