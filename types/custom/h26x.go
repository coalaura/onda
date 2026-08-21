package custom

import "github.com/coalaura/wtf/types"

func DetectH26x(b types.Buffer) *types.Metadata {
	startCodeOffset := findAnnexBStartCode(b)
	if startCodeOffset == -1 {
		return nil
	}

	nalOffset := startCodeOffset
	if nalOffset >= b.Len() {
		return nil
	}

	nalHeader := b[nalOffset]

	if nalHeader&0x80 != 0 {
		return nil
	}

	h264Type := nalHeader & 0x1F

	if nalOffset+1 < b.Len() {
		h265Type := (nalHeader >> 1) & 0x3F
		temporalIdPlus1 := b[nalOffset+1] & 0x07

		if temporalIdPlus1 >= 1 {
			switch h265Type {
			case 32, 33, 34:
				return &types.Metadata{Kind: types.KindH265Video}
			case 19, 20:
				return &types.Metadata{Kind: types.KindH265Video, Confidence: types.ConfidenceMedium}
			}
		}
	}

	switch h264Type {
	case 7:
		return &types.Metadata{Kind: types.KindH264Video}
	case 8:
		return &types.Metadata{Kind: types.KindH264Video, Confidence: types.ConfidenceMedium}
	case 5:
		return &types.Metadata{Kind: types.KindH264Video, Confidence: types.ConfidenceMedium}
	}

	return nil
}

func findAnnexBStartCode(b types.Buffer) int {
	if b.Len() < 4 {
		return -1
	}

	var zeroCount int

	for i := 0; i < b.Len() && i < 8; i++ {
		if b[i] == 0x00 {
			zeroCount++
		} else if b[i] == 0x01 && zeroCount >= 2 {
			return i + 1
		} else {
			break
		}
	}

	return -1
}
