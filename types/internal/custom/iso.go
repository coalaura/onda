package custom

import "github.com/coalaura/onda/types"

func DetectISOBaseMedia(b types.Buffer) *types.Metadata {
	boxSize32, ok := b.U32BE(0)
	if !ok || !b.Has(4, []byte("ftyp")) {
		return nil
	}

	brandOffset := 8
	compatibleOffset := 16
	boxEnd := int(boxSize32)

	if boxSize32 == 1 {
		high, ok := b.U32BE(8)
		if !ok {
			return nil
		}

		low, ok := b.U32BE(12)
		if !ok {
			return nil
		}

		if high != 0 {
			return nil
		}

		boxEnd = int(low)
		brandOffset = 16
		compatibleOffset = 24
	}

	if boxEnd <= compatibleOffset {
		return nil
	}

	if boxEnd > b.Len() {
		boxEnd = b.Len()
	}

	if !hasISOBrand(b, brandOffset, compatibleOffset, boxEnd, "isom", "iso2", "iso3", "iso4", "iso5", "iso6", "mp41", "mp42", "dash", "avif", "avis", "heic", "heix", "hevc", "hevx", "mif1", "msf1", "M4A ", "M4B ", "M4P ", "M4V ", "f4v ", "qt  ", "3gp4", "3gp5", "3gp6", "3gs7", "3ge6", "3gg6", "3gp1", "3gp2", "3g2a", "3g2b") {
		return nil
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxEnd, "avif") {
		return &types.Metadata{
			Kind: types.KindISOBaseMedia,
			Type: types.TypeAVIFImage,
		}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxEnd, "avis") {
		return &types.Metadata{
			Kind: types.KindISOBaseMedia,
			Type: types.TypeAVIFImageSequence,
		}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxEnd, "heic", "heix", "hevc", "hevx", "mif1", "msf1") {
		return &types.Metadata{
			Kind: types.KindISOBaseMedia,
			Type: types.TypeHEIFImage,
		}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxEnd, "M4V ") {
		return &types.Metadata{
			Kind: types.KindISOBaseMedia,
			Type: types.TypeM4VVideo,
		}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxEnd, "f4v ") {
		return &types.Metadata{
			Kind: types.KindISOBaseMedia,
			Type: types.TypeF4VVideo,
		}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxEnd, "3g2a", "3g2b") {
		return &types.Metadata{
			Kind: types.KindISOBaseMedia,
			Type: types.Type3G2Media,
		}
	}

	if hasISOBrandPrefix(b, brandOffset, compatibleOffset, boxEnd, "3gp", "3g2") {
		return &types.Metadata{
			Kind: types.KindISOBaseMedia,
			Type: types.Type3GPPMedia,
		}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxEnd, "M4A ", "M4B ", "M4P ") {
		return &types.Metadata{
			Kind: types.KindISOBaseMedia,
			Type: types.TypeMPEG4AudioM4AFamily,
		}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxEnd, "qt  ") {
		return &types.Metadata{
			Kind: types.KindISOBaseMedia,
			Type: types.TypeQuickTimeMovie,
		}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxEnd, "isom", "iso2", "iso3", "iso4", "iso5", "iso6", "mp41", "mp42", "dash") {
		return &types.Metadata{
			Kind: types.KindISOBaseMedia,
			Type: types.TypeMP4Video,
		}
	}

	return &types.Metadata{
		Kind: types.KindISOBaseMedia,
	}
}

func hasISOBrand(b types.Buffer, majorOffset int, compatOffset int, boxEnd int, brands ...string) bool {
	for _, brand := range brands {
		if len(brand) != 4 {
			continue
		}

		if b.Has(majorOffset, []byte(brand)) {
			return true
		}

		for off := compatOffset; off+4 <= boxEnd; off += 4 {
			if b.Has(off, []byte(brand)) {
				return true
			}
		}
	}

	return false
}

func hasISOBrandPrefix(b types.Buffer, majorOffset int, compatOffset int, boxEnd int, prefixes ...string) bool {
	if hasISOBrandPrefixAt(b, majorOffset, prefixes...) {
		return true
	}

	for off := compatOffset; off+4 <= boxEnd; off += 4 {
		if hasISOBrandPrefixAt(b, off, prefixes...) {
			return true
		}
	}

	return false
}

func hasISOBrandPrefixAt(b types.Buffer, offset int, prefixes ...string) bool {
	if offset < 0 || offset+4 > b.Len() {
		return false
	}

	for _, prefix := range prefixes {
		if len(prefix) != 3 {
			continue
		}

		if b[offset] == prefix[0] && b[offset+1] == prefix[1] && b[offset+2] == prefix[2] {
			return true
		}
	}

	return false
}
