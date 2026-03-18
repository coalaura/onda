package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature(types.KindBCFData, types.TypeNone, 0, []byte("BCF\x02"))
	types.RegisterSignature(types.KindHDF4Data, types.TypeNone, 0, []byte{0x0e, 0x03, 0x13, 0x01})
	types.RegisterSignature(types.KindLevelDB, types.TypeNone, 0, []byte{0x57, 0xfb, 0x80, 0x8b, 0x24, 0x75, 0x47, 0xdb})
	types.RegisterSignature(types.KindMATLABData, types.TypeNone, 0, []byte("MATLAB 5.0 MAT-file"))
	types.RegisterSignature(types.KindNumPyArray, types.TypeNone, 0, []byte("\x93NUMPY"))
	types.RegisterSignature(types.KindRData, types.TypeNone, 0, []byte("RDX2\n"))
	types.RegisterSignature(types.KindRData, types.TypeNone, 0, []byte("RDX3\n"))
	types.RegisterSignature(types.KindSASData, types.TypeNone, 12, []byte{0xc2, 0xea, 0x81, 0x60, 0xb3, 0x14, 0x11, 0xcf, 0xbd, 0x92, 0x08, 0x00, 0x09, 0xc7, 0x31, 0x8c, 0x18, 0x1f, 0x10, 0x11})
	types.RegisterSignature(types.KindSPSSData, types.TypeNone, 0, []byte("$FL2"))
	types.RegisterSignature(types.KindStataData, types.TypeNone, 0, []byte("<stata_dta>"))
}
