package custom

import (
	"encoding/binary"

	"github.com/coalaura/onda/types"
)

func DetectLZMA(b types.Buffer) *types.Metadata {
	if b.Len() < 13 {
		return nil
	}

	props := b[0]
	if props > 224 {
		return nil
	}

	dictSize, ok := b.U32LE(1)
	if !ok {
		return nil
	}

	if dictSize == 0 {
		return nil
	}

	if dictSize > (1 << 30) {
		return nil
	}

	// first byte of compressed stream (byte 13) MUST be 0x00
	if b.Len() > 13 && b[13] != 0x00 {
		return nil
	}

	uncompressedSize := binary.LittleEndian.Uint64(b[5:13])

	// check if size is reasonable
	if uncompressedSize != ^uint64(0) && uncompressedSize > (1<<50) {
		return nil
	}

	return &types.Metadata{
		Kind: types.KindLZMAData,
	}
}
