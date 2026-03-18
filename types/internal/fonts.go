package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature(types.KindEOTFont, types.TypeNone, 34, []byte("LP"))
	types.RegisterSignature(types.KindGlyphBitmapDistributionFormat, types.TypeNone, 0, []byte("STARTFONT "))
	types.RegisterSignature(types.KindOpenTypeFont, types.TypeNone, 0, []byte("OTTO"))
	types.RegisterSignature(types.KindPCFFont, types.TypeNone, 0, []byte("\x01fcp"))
	types.RegisterSignature(types.KindPSFFont, types.TypeNone, 0, []byte{0x36, 0x04})
	types.RegisterSignature(types.KindPSFFont, types.TypeNone, 0, []byte{0x72, 0xb5, 0x4a, 0x86})
	types.RegisterSignature(types.KindTrueTypeCollection, types.TypeNone, 0, []byte("ttcf"))
	types.RegisterSignature(types.KindTrueTypeFont, types.TypeNone, 0, []byte{0x00, 0x01, 0x00, 0x00})
	types.RegisterSignature(types.KindWOFF2Font, types.TypeNone, 0, []byte("wOF2"))
	types.RegisterSignature(types.KindWOFFFont, types.TypeNone, 0, []byte("wOFF"))
}
