package custom

import "github.com/coalaura/onda/types"

func DetectTIFFSubtypes(b types.Buffer) *types.Metadata {
	if !isTIFFHeader(b) {
		return nil
	}

	if b.Has(0, []byte{'I', 'I', 'U', 0x00}) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypePanasonicRAWRW2}
	}

	if hasBytesPrefix(b, []byte("Nikon")) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeNikonRAWNEF}
	}

	if hasBytesPrefix(b, []byte("PENTAX")) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypePentaxRAWPEF}
	}

	if hasBytesPrefix(b, []byte("SONY DSC")) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeSonyRAWARW}
	}

	if hasBytesPrefix(b, []byte("SONY SR2")) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeSonyRAWSR2}
	}

	if hasBytesPrefix(b, []byte("DNGVersion")) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeAdobeDNGDNG}
	}

	return nil
}

func isTIFFHeader(b types.Buffer) bool {
	return b.Has(0, []byte{'I', 'I', 42, 0}) || b.Has(0, []byte{'M', 'M', 0, 42})
}

func hasBytesPrefix(b types.Buffer, needle []byte) bool {
	if len(needle) == 0 || b.Len() < len(needle) {
		return false
	}

	limit := min(b.Len(), 1024) - len(needle)
	for i := 0; i <= limit; i++ {
		if b.Has(i, needle) {
			return true
		}
	}

	return false
}
