package custom

import (
	"slices"

	"github.com/coalaura/wtf/types"
)

func DetectISOBaseMedia(b types.Buffer) *types.Metadata {
	if b.Len() < 16 {
		return nil
	}

	boxSize32, ok := b.U32BE(0)
	if !ok || !b.Has(4, []byte("ftyp")) {
		return nil
	}

	brandOffset := 8
	compatibleOffset := 16
	boxEnd := uint64(boxSize32)

	switch boxSize32 {
	case 1:
		high, ok := b.U32BE(8)
		if !ok {
			return nil
		}

		low, ok := b.U32BE(12)
		if !ok {
			return nil
		}

		boxEnd = uint64(high)<<32 | uint64(low)
		brandOffset = 16
		compatibleOffset = 24
	case 0:
		boxEnd = uint64(b.Len())
	}

	if boxEnd < uint64(compatibleOffset) || b.Len() < compatibleOffset {
		return nil
	}

	if boxEnd > uint64(b.Len()) {
		return nil
	}

	boxEnd = min(boxEnd, uint64(maxScanSize))

	if !hasISOBrand(b, brandOffset, compatibleOffset, int(boxEnd), "isom", "iso2", "iso3", "iso4", "iso5", "iso6", "mp41", "mp42", "dash", "avif", "avis", "heic", "heix", "hevc", "hevx", "mif1", "msf1", "mjp2", "M4A ", "M4B ", "M4P ", "M4V ", "f4v ", "qt  ", "3gp4", "3gp5", "3gp6", "3gs7", "3ge6", "3gg6", "3gp1", "3gp2", "3g2a", "3g2b", "crx ", "braw", "MSNV") {
		return nil
	}

	boxLimit := int(boxEnd)

	// Get major brand
	var majorBrand string

	if brandOffset+4 <= b.Len() {
		majorBrand = string(b[brandOffset : brandOffset+4])
	}

	// Check major brand first
	switch majorBrand {
	case "qt  ":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeQuickTimeMovie}
	case "M4V ":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeM4VVideo}
	case "f4v ":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeF4VVideo}
	case "M4A ", "M4B ", "M4P ":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeMPEG4Audio}
	case "avif":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeAVIFImage}
	case "avis":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeAVIFImageSequence}
	case "heic", "heix", "hevc", "hevx", "mif1", "msf1":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeHEIFImage}
	case "mjp2":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeMotionJPEG2000Video}
	case "crx ":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeCanonRAW3Image}
	case "braw":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeBlackmagicRAWVideo}
	case "3g2a", "3g2b":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.Type3G2Video}
	}

	// Check for 3GPP prefix in major brand (3gp1-3gp6, 3gs7, 3ge6, 3gg6)
	if len(majorBrand) == 4 && (majorBrand[:3] == "3gp" || majorBrand[:3] == "3g2") {
		if majorBrand[3] >= '1' && majorBrand[3] <= '9' {
			return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.Type3GPPVideo}
		}
	}

	// For generic major brands (isom, iso2, mp41, etc.), check compatible brands
	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "avif") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeAVIFImage}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "avis") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeAVIFImageSequence}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "heic", "heix", "hevc", "hevx", "mif1", "msf1") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeHEIFImage}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "M4V ") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeM4VVideo}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "f4v ") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeF4VVideo}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "3g2a", "3g2b") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.Type3G2Video}
	}

	if hasISOBrandPrefix(b, brandOffset, compatibleOffset, boxLimit, "3gp", "3g2") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.Type3GPPVideo}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "M4A ", "M4B ", "M4P ") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeMPEG4Audio}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "qt  ") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeQuickTimeMovie}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "mjp2") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeMotionJPEG2000Video}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "crx ") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeCanonRAW3Image}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "braw") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeBlackmagicRAWVideo}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "isom", "iso2", "iso3", "iso4", "iso5", "iso6", "mp41", "mp42", "dash") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeMP4Video}
	}

	return &types.Metadata{
		Kind: types.KindISOBaseMedia,
	}
}

func hasISOBrand(b types.Buffer, majorOffset int, compatOffset int, boxEnd int, brands ...string) bool {
	if majorOffset+4 <= b.Len() {
		if slices.Contains(brands, string(b[majorOffset:majorOffset+4])) {
			return true
		}
	}

	for off := compatOffset; off+4 <= boxEnd && off+4 <= b.Len(); off += 4 {
		if slices.Contains(brands, string(b[off:off+4])) {
			return true
		}
	}

	return false
}

func hasISOBrandPrefix(b types.Buffer, majorOffset int, compatOffset int, boxEnd int, prefixes ...string) bool {
	if majorOffset+3 <= b.Len() {
		if slices.Contains(prefixes, string(b[majorOffset:majorOffset+3])) {
			return true
		}
	}

	for off := compatOffset; off+4 <= boxEnd && off+4 <= b.Len(); off += 4 {
		if slices.Contains(prefixes, string(b[off:off+3])) {
			return true
		}
	}

	return false
}
