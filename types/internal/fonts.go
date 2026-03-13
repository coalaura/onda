package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature(types.KindTrueTypeFont, types.TypeNone, 0, []byte{
		0x00, 0x01, 0x00, 0x00,
	})

	types.RegisterSignature(types.KindOpenTypeFont, types.TypeNone, 0, []byte("OTTO"))
	types.RegisterSignature(types.KindWOFFFont, types.TypeNone, 0, []byte("wOFF"))
	types.RegisterSignature(types.KindWOFF2Font, types.TypeNone, 0, []byte("wOF2"))
	types.RegisterSignature(types.KindTrueTypeCollection, types.TypeNone, 0, []byte("ttcf"))
	types.RegisterSignature(types.KindEOTFont, types.TypeNone, 34, []byte("LP"))
}
