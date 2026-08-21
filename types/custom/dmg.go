package custom

import "github.com/coalaura/wtf/types"

func DetectAppleDiskImage(b types.Buffer) *types.Metadata {
	if b.Len() < 512 {
		return nil
	}

	offset := b.Len() - 512

	version, versionOK := b.U32BE(offset + 4)
	headerSize, sizeOK := b.U32BE(offset + 8)
	segmentNumber, segmentOK := b.U32BE(offset + 56)
	segmentCount, countOK := b.U32BE(offset + 60)

	if b.Has(offset, []byte("koly")) && versionOK && version == 4 && sizeOK && headerSize == 512 && segmentOK && segmentNumber > 0 && countOK && segmentCount > 0 && segmentNumber <= segmentCount {
		return &types.Metadata{
			Kind: types.KindAppleDiskImage,
		}
	}

	return nil
}
