package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature("TrueType Font", "", 0, []byte{
		0x00, 0x01, 0x00, 0x00,
	})

	types.RegisterSignature("OpenType Font", "", 0, []byte("OTTO"))
	types.RegisterSignature("WOFF Font", "", 0, []byte("wOFF"))
	types.RegisterSignature("WOFF2 Font", "", 0, []byte("wOF2"))
	types.RegisterSignature("TrueType Collection", "", 0, []byte("ttcf"))
}
