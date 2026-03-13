package custom

import "github.com/coalaura/onda/types"

func DetectNetpbm(b types.Buffer) *types.Metadata {
	if b.Len() < 3 {
		return nil
	}

	if b[0] != 'P' {
		return nil
	}

	if !isNetpbmDelimiter(b[2]) {
		return nil
	}

	switch b[1] {
	case '1':
		return &types.Metadata{
			Kind: types.KindNetpbmImage,
			Type: types.TypePBMASCII,
		}
	case '2':
		return &types.Metadata{
			Kind: types.KindNetpbmImage,
			Type: types.TypePGMASCII,
		}
	case '3':
		return &types.Metadata{
			Kind: types.KindNetpbmImage,
			Type: types.TypePPMASCII,
		}
	case '4':
		return &types.Metadata{
			Kind: types.KindNetpbmImage,
			Type: types.TypePBMBinary,
		}
	case '5':
		return &types.Metadata{
			Kind: types.KindNetpbmImage,
			Type: types.TypePGMBinary,
		}
	case '6':
		return &types.Metadata{
			Kind: types.KindNetpbmImage,
			Type: types.TypePPMBinary,
		}
	case '7':
		return &types.Metadata{
			Kind: types.KindNetpbmImage,
			Type: types.TypePAM,
		}
	default:
		return nil
	}
}

func isNetpbmDelimiter(c byte) bool {
	switch c {
	case ' ', '\t', '\r', '\n', '\f':
		return true
	default:
		return false
	}
}
