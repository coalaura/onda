package internal

import "github.com/coalaura/wtf/types"

func init() {
	types.RegisterSignature(types.KindAutodesk123DDesign, types.TypeNone, 0, []byte("123D\x00"))
	types.RegisterSignature(types.KindAutodeskAliasStudio, types.TypeNone, 0, []byte("AliasStudio"))
	types.RegisterSignature(types.KindBlenderCache, types.TypeNone, 0, []byte("BPHYSICS"))
	types.RegisterSignature(types.KindHalfLifeModel, types.TypeHalfLife1, 0, []byte("IDST\x0a\x00\x00\x00"))
	types.RegisterSignature(types.KindHalfLifeModel, types.TypeHalfLife2, 0, []byte("IDST\x2c\x00\x00\x00"))
	types.RegisterSignature(types.KindHalfLifeModel, types.TypeHalfLife2, 0, []byte("IDST\x2d\x00\x00\x00"))
	types.RegisterSignature(types.KindHalfLifeModel, types.TypeHalfLife2, 0, []byte("IDST\x2e\x00\x00\x00"))
	types.RegisterSignature(types.KindHalfLifeModel, types.TypeHalfLife2, 0, []byte("IDST\x2f\x00\x00\x00"))
	types.RegisterSignature(types.KindHalfLifeModel, types.TypeHalfLife2, 0, []byte("IDST\x30\x00\x00\x00"))
	types.RegisterSignature(types.KindHalfLifeModel, types.TypeHalfLife2, 0, []byte("IDST\x31\x00\x00\x00"))
	types.RegisterSignature(types.KindMayaProject, types.TypeNone, 0, []byte("MayaProject\x00"))
	types.RegisterSignature(types.KindSpineAnimation, types.TypeSpineBinary, 0, []byte{0x13, 0x53, 0x70, 0x69, 0x6e, 0x65, 0x20, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e})
}
