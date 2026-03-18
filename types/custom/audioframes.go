package custom

import "github.com/coalaura/onda/types"

func DetectMPEGAudioFrames(b types.Buffer) *types.Metadata {
	if b.Len() >= 6 && b[0] == 0x0b && b[1] == 0x77 {
		bsid := (b[5] >> 3) & 0x1f
		if bsid <= 10 {
			return &types.Metadata{Kind: types.KindMPEGAudioFrame, Type: types.TypeAC3}
		}

		return &types.Metadata{Kind: types.KindMPEGAudioFrame, Type: types.TypeEAC3}
	}

	return nil
}
