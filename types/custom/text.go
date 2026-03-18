package custom

import (
	"bytes"
	"unicode/utf8"

	"github.com/coalaura/onda/types"
)

func DetectText(b types.Buffer) *types.Metadata {
	if b.Len() == 0 {
		return nil
	}

	if !passesTextLengthGate(b) {
		return nil
	}

	if b.Has(0, []byte{0xef, 0xbb, 0xbf}) {
		if isLikelyTextUTF8(b[3:]) {
			return &types.Metadata{Kind: types.KindTextFile, Type: types.TypeUTF8Text}
		}
	}

	if isLikelyASCIIText(b) {
		return &types.Metadata{Kind: types.KindTextFile, Type: types.TypeASCIIText}
	}

	if isLikelyTextUTF8(b) {
		return &types.Metadata{Kind: types.KindTextFile, Type: types.TypeUTF8Text}
	}

	return nil
}

func passesTextLengthGate(b types.Buffer) bool {
	if b.Len() >= 16 {
		return true
	}

	return bytes.ContainsAny(b, " \t\r\n")
}

func isLikelyASCIIText(b types.Buffer) bool {
	limit := min(b.Len(), 4096)
	if limit == 0 {
		return false
	}

	var printable int

	target := b[:limit]

	for _, c := range target {
		if c == 0 {
			return false
		}

		if c == '\n' || c == '\r' || c == '\t' || (c >= 32 && c <= 126) {
			printable++
		}
	}

	return printable*100/limit >= 95
}

func isLikelyTextUTF8(b types.Buffer) bool {
	limit := min(b.Len(), 4096)
	if limit == 0 {
		return false
	}

	target := b[:limit]
	if !utf8.Valid(target) {
		return false
	}

	var printable int

	for _, c := range target {
		if c == 0 {
			return false
		}

		if c == '\n' || c == '\r' || c == '\t' || c >= 32 {
			printable++
		}
	}

	return printable*100/limit >= 90
}
