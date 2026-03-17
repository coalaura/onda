package custom

import "github.com/coalaura/onda/types"

func DetectZlib(b types.Buffer) *types.Metadata {
	if b.Len() < 2 {
		return nil
	}

	cmf := b[0]
	flg := b[1]

	if cmf&0x0f != 8 {
		return nil
	}

	if cmf>>4 > 7 {
		return nil
	}

	if (uint16(cmf)<<8|uint16(flg))%31 != 0 {
		return nil
	}

	return &types.Metadata{
		Kind: types.KindZlibData,
	}
}
