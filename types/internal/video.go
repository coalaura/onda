package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature("FLV Video", "", 0, []byte("FLV"))
	types.RegisterSignature("IVF Video", "", 0, []byte("DKIF"))
	types.RegisterSignature("MPEG Program Stream", "", 0, []byte{0x00, 0x00, 0x01, 0xba})
	types.RegisterSignature("ASF Container", "", 0, []byte{0x30, 0x26, 0xb2, 0x75, 0x8e, 0x66, 0xcf, 0x11, 0xa6, 0xd9, 0x00, 0xaa, 0x00, 0x62, 0xce, 0x6c})
}
