package custom

import "github.com/coalaura/onda/types"

func DetectMP3(b types.Buffer) *types.Metadata {
	if b.Has(0, []byte("ID3")) {
		return &types.Metadata{
			Name: "MPEG Audio",
			Type: "MP3 (ID3 Tagged)",
		}
	}

	if b.Len() < 4 {
		return nil
	}

	if b[0] != 0xff || (b[1]&0xe0) != 0xe0 {
		return nil
	}

	version := (b[1] >> 3) & 0x03
	layer := (b[1] >> 1) & 0x03
	bitrate := (b[2] >> 4) & 0x0f
	sampleRate := (b[2] >> 2) & 0x03

	if version == 1 || layer == 0 || bitrate == 0 || bitrate == 0x0f || sampleRate == 0x03 {
		return nil
	}

	return &types.Metadata{
		Name: "MPEG Audio",
		Type: "MP3",
	}
}
