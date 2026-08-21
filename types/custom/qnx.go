package custom

import "github.com/coalaura/wtf/types"

const qnx4SuperOffset = 512

func DetectQNX4(b types.Buffer) *types.Metadata {
	if !b.Has(qnx4SuperOffset, []byte{'/', 0}) {
		return nil
	}

	rootBlock, ok := b.U32LE(qnx4SuperOffset + 20)
	if !ok || rootBlock == 0 {
		return nil
	}

	rootBlocks, ok := b.U32LE(qnx4SuperOffset + 24)
	if !ok || rootBlocks == 0 {
		return nil
	}

	rootOffset := uint64(rootBlock-1) * 512
	rootSize := min(uint64(rootBlocks)*512, uint64(maxScanSize))

	if rootOffset > uint64(b.Len()) || rootSize > uint64(b.Len())-rootOffset {
		return nil
	}

	for offset := int(rootOffset); offset+64 <= int(rootOffset+rootSize); offset += 64 {
		if b.Has(offset, []byte(".bitmap\x00")) {
			return &types.Metadata{Kind: types.KindQNX4Filesystem}
		}
	}

	return nil
}
