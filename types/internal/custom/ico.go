package custom

import "github.com/coalaura/onda/types"

func DetectICO(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte{0, 0}) {
		return nil
	}

	reserved, ok := b.U16LE(0)
	if !ok {
		return nil
	}

	typ, ok := b.U16LE(2)
	if !ok {
		return nil
	}

	if reserved != 0 {
		return nil
	}

	switch typ {
	case 1:
		return &types.Metadata{
			Kind: types.KindICOCURImage,
			Type: types.TypeWindowsIcon,
		}
	case 2:
		return &types.Metadata{
			Kind: types.KindICOCURImage,
			Type: types.TypeWindowsCursor,
		}
	default:
		return nil
	}
}
