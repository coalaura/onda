package custom

import "github.com/coalaura/onda/types"

func DetectPCX(b types.Buffer) *types.Metadata {
	if b.Len() < 128 || b[0] != 0x0a {
		return nil
	}

	if !isPCXVersion(b[1]) {
		return nil
	}

	if b[2] != 1 {
		return nil
	}

	if !isPCXBitsPerPixel(b[3]) {
		return nil
	}

	xMin, ok := b.U16LE(4)
	if !ok {
		return nil
	}

	yMin, ok := b.U16LE(6)
	if !ok {
		return nil
	}

	xMax, ok := b.U16LE(8)
	if !ok {
		return nil
	}

	yMax, ok := b.U16LE(10)
	if !ok {
		return nil
	}

	if xMax < xMin || yMax < yMin {
		return nil
	}

	if b[64] != 0 {
		return nil
	}

	if !isPCXColorPlanes(b[65]) {
		return nil
	}

	bytesPerLine, ok := b.U16LE(66)
	if !ok || bytesPerLine == 0 {
		return nil
	}

	return &types.Metadata{
		Kind: types.KindPCXImage,
	}
}

func isPCXVersion(v byte) bool {
	switch v {
	case 0, 2, 3, 4, 5:
		return true
	default:
		return false
	}
}

func isPCXBitsPerPixel(v byte) bool {
	switch v {
	case 1, 2, 4, 8:
		return true
	default:
		return false
	}
}

func isPCXColorPlanes(v byte) bool {
	switch v {
	case 1, 3, 4:
		return true
	default:
		return false
	}
}
