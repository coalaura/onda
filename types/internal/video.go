package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature(types.KindASFContainer, types.TypeNone, 0, []byte{0x30, 0x26, 0xb2, 0x75, 0x8e, 0x66, 0xcf, 0x11, 0xa6, 0xd9, 0x00, 0xaa, 0x00, 0x62, 0xce, 0x6c})
	types.RegisterSignature(types.KindBinkVideo, types.TypeNone, 0, []byte("BIKb"))
	types.RegisterSignature(types.KindBinkVideo, types.TypeNone, 0, []byte("BIKf"))
	types.RegisterSignature(types.KindBinkVideo, types.TypeNone, 0, []byte("BIKg"))
	types.RegisterSignature(types.KindBinkVideo, types.TypeNone, 0, []byte("BIKh"))
	types.RegisterSignature(types.KindBinkVideo, types.TypeNone, 0, []byte("BIKi"))
	types.RegisterSignature(types.KindDiracVideo, types.TypeNone, 0, []byte("BBCD"))
	types.RegisterSignature(types.KindDVRMSVideo, types.TypeNone, 0, []byte{0x05, 0x16, 0x11, 0xc3, 0x53, 0x12, 0x04, 0x11, 0xa3, 0x14, 0x00, 0x90, 0x27, 0x51, 0x46, 0xf3})
	types.RegisterSignature(types.KindFLVVideo, types.TypeNone, 0, []byte("FLV"))
	types.RegisterSignature(types.KindIVFVideo, types.TypeNone, 0, []byte("DKIF"))
	types.RegisterSignature(types.KindM3U8Playlist, types.TypeNone, 0, []byte("#EXTM3U"))
	types.RegisterSignature(types.KindMaterialExchangeFormat, types.TypeNone, 0, []byte{0x06, 0x0e, 0x2b, 0x34, 0x02, 0x05, 0x01, 0x01, 0x0d, 0x01, 0x02, 0x01, 0x01, 0x02})
	types.RegisterSignature(types.KindMPEGProgramStream, types.TypeNone, 0, []byte{0x00, 0x00, 0x01, 0xba})
	types.RegisterSignature(types.KindMPEGVideo, types.TypeMPEG12Video, 0, []byte{0x00, 0x00, 0x01, 0xb3})
	types.RegisterSignature(types.KindNullsoftVideo, types.TypeNone, 0, []byte("NSVf"))
	types.RegisterSignature(types.KindRealMedia, types.TypeNone, 0, []byte{'.', 'R', 'M', 'F'})
	types.RegisterSignature(types.KindRIFFContainer, types.TypeNone, 0, []byte("RIFF"))
	types.RegisterSignature(types.KindSiliconGraphicsMovie, types.TypeNone, 0, []byte("MOVI"))
	types.RegisterSignature(types.KindSmackerVideo, types.TypeNone, 0, []byte("SMK2"))
	types.RegisterSignature(types.KindSmackerVideo, types.TypeNone, 0, []byte("SMK4"))
	types.RegisterSignature(types.KindWTVVideo, types.TypeNone, 0, []byte{0xb7, 0xd8, 0x00, 0x20, 0x37, 0x49, 0xda, 0x11, 0xa6, 0x4e, 0x00, 0x07, 0xe9, 0x5e, 0xad, 0x8d})
	types.RegisterSignature(types.KindYUV4MPEG2Video, types.TypeNone, 0, []byte("YUV4MPEG2 "))

	types.RegisterMaskedSignature(types.KindRIFFContainer, types.TypeAfterEffectsProjectAEP, 0, []byte("RIFF\x00\x00\x00\x00FXTC"), []byte{0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff})
	types.RegisterMaskedSignature(types.KindRIFFContainer, types.TypeAVIVideo, 0, []byte("RIFF\x00\x00\x00\x00AVI "), []byte{0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff})
	types.RegisterMaskedSignature(types.KindRIFFContainer, types.TypeCubaseProjectCPR, 0, []byte("RIFF\x00\x00\x00\x00SMED"), []byte{0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff})
	types.RegisterMaskedSignature(types.KindRIFFContainer, types.TypeWindowsAnimatedCursor, 0, []byte("RIFF\x00\x00\x00\x00ACON"), []byte{0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff})
}
