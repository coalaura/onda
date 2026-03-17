//go:build ignore

package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature(types.KindCloneCDControl, types.TypeNone, 0, []byte("[CloneCD]"))
	types.RegisterSignature(types.KindDolphinRVZ, types.TypeNone, 0, []byte("RVZ\x01"))
	types.RegisterSignature(types.KindMAMECHD, types.TypeNone, 0, []byte("MComprHD"))
	types.RegisterSignature(types.KindNintendoSwitchNRO, types.TypeNone, 16, []byte("NRO0"))
	types.RegisterSignature(types.KindNintendoSwitchNSO, types.TypeNone, 0, []byte("NSO0"))
	types.RegisterSignature(types.KindPlayStationExecutable, types.TypeNone, 0, []byte("SCE\x00"))
	types.RegisterSignature(types.KindPlayStationPortableExecutable, types.TypeNone, 0, []byte("\x00PBP"))
	types.RegisterSignature(types.KindSourceEngineBSP, types.TypeNone, 0, []byte("VBSP"))
	types.RegisterSignature(types.KindValvePak, types.TypeNone, 0, []byte{0x34, 0x12, 0xaa, 0x55})
	types.RegisterSignature(types.KindWiiUArchive, types.TypeNone, 0, []byte("WUA\x01"))
	types.RegisterSignature(types.KindXbox360Executable, types.TypeNone, 0, []byte("XEX2"))
	types.RegisterSignature(types.KindXboxExecutable, types.TypeNone, 0, []byte("XBEH"))
}
