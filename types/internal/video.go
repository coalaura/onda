package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature(types.KindASFContainer, types.TypeNone, 0, []byte{0x30, 0x26, 0xb2, 0x75, 0x8e, 0x66, 0xcf, 0x11, 0xa6, 0xd9, 0x00, 0xaa, 0x00, 0x62, 0xce, 0x6c})
	types.RegisterSignature(types.KindFLVVideo, types.TypeNone, 0, []byte("FLV"))
	types.RegisterSignature(types.KindIVFVideo, types.TypeNone, 0, []byte("DKIF"))
	types.RegisterSignature(types.KindM3U8Playlist, types.TypeNone, 0, []byte("#EXTM3U"))
	types.RegisterSignature(types.KindMPEGProgramStream, types.TypeNone, 0, []byte{0x00, 0x00, 0x01, 0xba})
	types.RegisterSignature(types.KindRealMedia, types.TypeNone, 0, []byte{'.', 'R', 'M', 'F'})
	types.RegisterSignature(types.KindRIFFContainer, types.TypeNone, 0, []byte("RIFF"))
	types.RegisterSignature(types.KindWTVVideo, types.TypeNone, 0, []byte{0xb7, 0xd8, 0x00, 0x20, 0x37, 0x49, 0xda, 0x11, 0xa6, 0x4e, 0x00, 0x07, 0xe9, 0x5e, 0xad, 0x8d})

	types.RegisterMaskedSignature(types.KindRIFFContainer, types.TypeAVIVideo, 0, []byte("RIFF\x00\x00\x00\x00AVI "), []byte{0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff})
}
