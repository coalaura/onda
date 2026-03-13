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

	if b.Has(0, []byte{0x7f, 0xfe, 0x80, 0x01}) || b.Has(0, []byte{0xfe, 0x7f, 0x01, 0x80}) || b.Has(0, []byte{0x1f, 0xff, 0xe8, 0x00}) || b.Has(0, []byte{0xff, 0x1f, 0x00, 0xe8}) {
		return &types.Metadata{Kind: types.KindDTSAudio}
	}

	if b.Len() < 4 {
		return nil
	}

	if b[0] != 0xff || (b[1]&0xe0) != 0xe0 {
		return nil
	}

	layerBits := (b[1] >> 1) & 0x03
	if layerBits == 2 {
		return &types.Metadata{Kind: types.KindMPEGAudioFrame, Type: types.TypeMPEGLayer2}
	}

	if layerBits == 1 {
		return &types.Metadata{Kind: types.KindMPEGAudioFrame, Type: types.TypeMPEGLayer3}
	}

	return nil
}
