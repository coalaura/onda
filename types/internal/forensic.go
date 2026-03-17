//go:build ignore

package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature(types.KindAdvancedForensicFormat, types.TypeNone, 0, []byte("AFF\x00"))
	types.RegisterSignature(types.KindEnCaseImage, types.TypeNone, 0, []byte("EVF\x09\x0d\x0a\xff\x00"))
	types.RegisterSignature(types.KindMediaDescriptor, types.TypeNone, 0, []byte("MEDIA DESCRIPTOR"))
	types.RegisterSignature(types.KindPowerISODAA, types.TypeNone, 0, []byte("DAA\x1a"))
}
