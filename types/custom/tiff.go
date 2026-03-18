package custom

import (
	"bytes"

	"github.com/coalaura/onda/types"
)

func DetectTIFFSubtypes(b types.Buffer) *types.Metadata {
	if !isTIFFHeader(b) {
		if b.Has(0, []byte{'I', 'I', 'R', 'O', 0x08, 0x00}) || b.Has(0, []byte{'M', 'M', 'O', 'R', 0x00, 0x00}) {
			return &types.Metadata{Kind: types.KindOlympusRAWImage}
		}

		return nil
	}

	if b.Has(0, []byte{'I', 'I', 'U', 0x00}) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypePanasonicRAWRW2}
	}

	if b.Has(0, []byte{'I', 'I', 0x2a, 0x00, 0x10, 0x00, 0x00, 0x00, 'C', 'R'}) {
		return &types.Metadata{Kind: types.KindCanonRAWImage}
	}

	limit := min(b.Len(), 1024)
	data := b[:limit]

	if bytes.Contains(data, []byte("Nikon")) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeNikonRAWNEF}
	}

	if bytes.Contains(data, []byte("PENTAX")) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypePentaxRAWPEF}
	}

	if bytes.Contains(data, []byte("SONY DSC")) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeSonyRAWARW}
	}

	if bytes.Contains(data, []byte("SONY SR2")) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeSonyRAWSR2}
	}

	if bytes.Contains(data, []byte("DNGVersion")) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeAdobeDNGDNG}
	}

	if b.Has(0, []byte{'I', 'I'}) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeLittleEndian}
	}

	return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeBigEndian}
}

func isTIFFHeader(b types.Buffer) bool {
	return b.Has(0, []byte{'I', 'I', 42, 0}) || b.Has(0, []byte{'M', 'M', 0, 42})
}
