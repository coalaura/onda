package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature("TrueType font", "", 0, []byte{
		0x00, 0x01, 0x00, 0x00,
	})

	types.RegisterSignature("OpenType font", "", 0, []byte("OTTO"))
	types.RegisterSignature("WOFF font", "", 0, []byte("wOFF"))
	types.RegisterSignature("WOFF2 font", "", 0, []byte("wOF2"))
	types.RegisterSignature("TrueType collection", "", 0, []byte("ttcf"))
}
