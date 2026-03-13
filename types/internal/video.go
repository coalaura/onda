package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature("FLV video", "", 0, []byte("FLV"))
}
