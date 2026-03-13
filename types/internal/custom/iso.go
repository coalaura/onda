package custom

import "github.com/coalaura/onda/types"

func DetectISOBaseMedia(b types.Buffer) *types.Metadata {
	size, ok := b.U32BE(0)
	if !ok {
		return nil
	}

	if size < 8 {
		return nil
	}

	if !b.Has(4, []byte("ftyp")) {
		return nil
	}

	if b.Has(8, []byte("avif")) {
		return &types.Metadata{
			Name: "AVIF image",
		}
	}

	if b.Has(8, []byte("avis")) {
		return &types.Metadata{
			Name: "AVIF image",
			Type: "sequence",
		}
	}

	if b.Has(8, []byte("heic")) || b.Has(8, []byte("heix")) || b.Has(8, []byte("hevc")) || b.Has(8, []byte("hevx")) || b.Has(8, []byte("mif1")) || b.Has(8, []byte("msf1")) {
		return &types.Metadata{
			Name: "HEIF image",
		}
	}

	if b.Has(8, []byte("isom")) || b.Has(8, []byte("iso2")) || b.Has(8, []byte("mp41")) || b.Has(8, []byte("mp42")) {
		return &types.Metadata{
			Name: "MP4 video",
		}
	}

	if b.Has(8, []byte("M4A ")) {
		return &types.Metadata{
			Name: "MPEG-4 audio",
			Type: "M4A",
		}
	}

	if b.Has(8, []byte("qt  ")) {
		return &types.Metadata{
			Name: "QuickTime movie",
		}
	}

	if b.Has(8, []byte("3gp")) || b.Has(8, []byte("3gs")) {
		return &types.Metadata{
			Name: "3GPP media",
		}
	}

	return &types.Metadata{
		Name: "ISO Base Media file",
	}
}
