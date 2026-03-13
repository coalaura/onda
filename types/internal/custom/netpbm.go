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
			Name: "Netpbm Image",
			Type: "PBM ASCII",
		}
	case '2':
		return &types.Metadata{
			Name: "Netpbm Image",
			Type: "PGM ASCII",
		}
	case '3':
		return &types.Metadata{
			Name: "Netpbm Image",
			Type: "PPM ASCII",
		}
	case '4':
		return &types.Metadata{
			Name: "Netpbm Image",
			Type: "PBM binary",
		}
	case '5':
		return &types.Metadata{
			Name: "Netpbm Image",
			Type: "PGM binary",
		}
	case '6':
		return &types.Metadata{
			Name: "Netpbm Image",
			Type: "PPM binary",
		}
	case '7':
		return &types.Metadata{
			Name: "Netpbm Image",
			Type: "PAM",
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
