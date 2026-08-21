package custom

import (
	"bytes"

	"github.com/coalaura/wtf/types"
)

type tiffByteOrder uint8

const (
	tiffLittleEndian tiffByteOrder = iota
	tiffBigEndian
)

func DetectTIFFSubtypes(b types.Buffer) *types.Metadata {
	if b.Has(0, []byte{'I', 'I', 'U', 0x00, 0x18, 0x00, 0x00, 0x00}) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypePanasonicRAW}
	}

	if !isTIFFHeader(b) {
		if b.Has(0, []byte{'I', 'I', 'R', 'O', 0x08, 0x00}) || b.Has(0, []byte{'M', 'M', 'O', 'R', 0x00, 0x00}) {
			return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeOlympusRAW}
		}

		return nil
	}

	if b.Has(0, []byte{'I', 'I', 0x2a, 0x00, 0x10, 0x00, 0x00, 0x00, 'C', 'R'}) {
		return &types.Metadata{Kind: types.KindCanonRAWImage, Type: types.TypeHE}
	}

	order, ifd0, ok := tiffHeaderInfo(b)
	if !ok {
		if b.Has(0, []byte{'I', 'I'}) {
			return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeLittleEndian}
		}

		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeBigEndian}
	}

	if tiffHasTag(b, order, ifd0, 0xc612) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeAdobeDNG}
	}

	makeValue, ok := tiffASCIIValueForTag(b, order, ifd0, 0x010f)
	if ok {
		if bytes.HasPrefix(makeValue, []byte("NIKON")) || bytes.HasPrefix(makeValue, []byte("Nikon")) {
			return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeNikonRAW}
		}

		if bytes.HasPrefix(makeValue, []byte("OLYMPUS")) || bytes.HasPrefix(makeValue, []byte("Olympus")) {
			return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeOlympusRAW}
		}

		if bytes.HasPrefix(makeValue, []byte("PENTAX")) {
			return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypePentaxRAW}
		}

		if bytes.HasPrefix(makeValue, []byte("SONY")) {
			if bytes.Contains(makeValue, []byte("SR2")) {
				return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeSonyRAWSR2}
			}

			return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeSonyRAW}
		}
	}

	limit := min(b.Len(), 4096)
	data := b[:limit]

	if bytes.Contains(data, []byte("Nikon")) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeNikonRAW}
	}

	if bytes.Contains(data, []byte("OLYMPUS")) || bytes.Contains(data, []byte("Olympus")) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeOlympusRAW}
	}

	if bytes.Contains(data, []byte("PENTAX")) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypePentaxRAW}
	}

	if bytes.Contains(data, []byte("SONY DSC")) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeSonyRAW}
	}

	if bytes.Contains(data, []byte("SONY SR2")) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeSonyRAWSR2}
	}

	if bytes.Contains(data, []byte("DNGVersion")) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeAdobeDNG}
	}

	if b.Has(0, []byte{'I', 'I'}) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeLittleEndian}
	}

	return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeBigEndian}
}

func tiffHeaderInfo(b types.Buffer) (tiffByteOrder, uint32, bool) {
	if b.Has(0, []byte{'I', 'I', 42, 0}) {
		ifd0, ok := b.U32LE(4)
		if !ok || ifd0 >= uint32(b.Len()) {
			return 0, 0, false
		}

		return tiffLittleEndian, ifd0, true
	}

	if b.Has(0, []byte{'M', 'M', 0, 42}) {
		ifd0, ok := b.U32BE(4)
		if !ok || ifd0 >= uint32(b.Len()) {
			return 0, 0, false
		}

		return tiffBigEndian, ifd0, true
	}

	return 0, 0, false
}

func tiffHasTag(b types.Buffer, order tiffByteOrder, ifdOffset uint32, tag uint16) bool {
	count, ok := tiffU16(b, order, int(ifdOffset))
	if !ok {
		return false
	}

	entryBase := int(ifdOffset) + 2
	entryEnd := entryBase + int(count)*12
	if entryBase < 0 || entryEnd > b.Len() {
		return false
	}

	for i := range int(count) {
		entry := entryBase + i*12

		currentTag, ok := tiffU16(b, order, entry)
		if !ok {
			return false
		}

		if currentTag == tag {
			return true
		}
	}

	return false
}

func tiffASCIIValueForTag(b types.Buffer, order tiffByteOrder, ifdOffset uint32, tag uint16) ([]byte, bool) {
	count, ok := tiffU16(b, order, int(ifdOffset))
	if !ok {
		return nil, false
	}

	entryBase := int(ifdOffset) + 2
	entryEnd := entryBase + int(count)*12

	if entryBase < 0 || entryEnd > b.Len() {
		return nil, false
	}

	for i := range int(count) {
		entry := entryBase + i*12

		currentTag, ok := tiffU16(b, order, entry)
		if !ok || currentTag != tag {
			continue
		}

		fieldType, ok := tiffU16(b, order, entry+2)
		if !ok || fieldType != 2 {
			return nil, false
		}

		valueCount, ok := tiffU32(b, order, entry+4)
		if !ok || valueCount == 0 {
			return nil, false
		}

		if valueCount <= 4 {
			end := entry + 8 + int(valueCount)
			if end > b.Len() {
				return nil, false
			}

			return bytes.TrimRight(b[entry+8:end], "\x00"), true
		}

		valueOffset, ok := tiffU32(b, order, entry+8)
		if !ok {
			return nil, false
		}

		start := int(valueOffset)
		end := start + int(valueCount)

		if start < 0 || end > b.Len() {
			return nil, false
		}

		return bytes.TrimRight(b[start:end], "\x00"), true
	}

	return nil, false
}

func tiffU16(b types.Buffer, order tiffByteOrder, offset int) (uint16, bool) {
	if order == tiffLittleEndian {
		return b.U16LE(offset)
	}

	return b.U16BE(offset)
}

func tiffU32(b types.Buffer, order tiffByteOrder, offset int) (uint32, bool) {
	if order == tiffLittleEndian {
		return b.U32LE(offset)
	}

	return b.U32BE(offset)
}

func isTIFFHeader(b types.Buffer) bool {
	return b.Has(0, []byte{'I', 'I', 42, 0}) || b.Has(0, []byte{'M', 'M', 0, 42})
}
