package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature(types.KindFLACAudio, types.TypeNone, 0, []byte("fLaC"))
	types.RegisterSignature(types.KindMIDISequence, types.TypeNone, 0, []byte("MThd"))
	types.RegisterSignature(types.KindAUAudio, types.TypeNone, 0, []byte(".snd"))
	types.RegisterSignature(types.KindAMRAudio, types.TypeNone, 0, []byte("#!AMR\n"))
	types.RegisterSignature(types.KindAMRWBAudio, types.TypeNone, 0, []byte("#!AMR-WB\n"))
	types.RegisterSignature(types.KindAACAudio, types.TypeADTS, 0, []byte{0xff, 0xf1})
	types.RegisterSignature(types.KindAACAudio, types.TypeADTS, 0, []byte{0xff, 0xf9})
	types.RegisterSignature(types.KindCAFAudio, types.TypeNone, 0, []byte("caff"))
	types.RegisterSignature(types.KindWavPackAudio, types.TypeNone, 0, []byte("wvpk"))
	types.RegisterSignature(types.KindMonkeysAudio, types.TypeNone, 0, []byte("MAC "))
	types.RegisterSignature(types.KindDSFAudio, types.TypeNone, 0, []byte("DSD "))
	types.RegisterSignature(types.KindTAKAudio, types.TypeNone, 0, []byte("tBaK"))
	types.RegisterSignature(types.KindVOCAudio, types.TypeNone, 0, []byte("Creative Voice File\x1a"))
	types.RegisterSignature(types.KindMusepackAudio, types.TypeStreamVersion7, 0, []byte("MP+"))
	types.RegisterSignature(types.KindMusepackAudio, types.TypeStreamVersion8, 0, []byte("MPCK"))
	types.RegisterSignature(types.KindImpulseTrackerModule, types.TypeNone, 0, []byte("IMPM"))
	types.RegisterSignature(types.KindFastTrackerModule, types.TypeNone, 0, []byte("Extended Module: "))
	types.RegisterSignature(types.KindScreamTrackerModule, types.TypeNone, 44, []byte("SCRM"))
	types.RegisterSignature(types.KindTTAAudio, types.TypeNone, 0, []byte("TTA1"))
}
