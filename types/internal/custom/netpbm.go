package custom

import "github.com/coalaura/onda/types"

func DetectNetpbm(b types.Buffer) *types.Metadata {
	if b.Len() < 2 {
		return nil
	}

	if b[0] != 'P' {
		return nil
	}

	switch b[1] {
	case '1':
		return &types.Metadata{
			Name: "Netpbm image",
			Type: "PBM ASCII",
		}
	case '2':
		return &types.Metadata{
			Name: "Netpbm image",
			Type: "PGM ASCII",
		}
	case '3':
		return &types.Metadata{
			Name: "Netpbm image",
			Type: "PPM ASCII",
		}
	case '4':
		return &types.Metadata{
			Name: "Netpbm image",
			Type: "PBM binary",
		}
	case '5':
		return &types.Metadata{
			Name: "Netpbm image",
			Type: "PGM binary",
		}
	case '6':
		return &types.Metadata{
			Name: "Netpbm image",
			Type: "PPM binary",
		}
	case '7':
		return &types.Metadata{
			Name: "Netpbm image",
			Type: "PAM",
		}
	default:
		return nil
	}
}
