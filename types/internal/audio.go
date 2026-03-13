package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature("FLAC Audio", "", 0, []byte("fLaC"))
	types.RegisterSignature("MIDI Sequence", "", 0, []byte("MThd"))
	types.RegisterSignature("AU Audio", "", 0, []byte(".snd"))
	types.RegisterSignature("AMR Audio", "", 0, []byte("#!AMR\n"))
	types.RegisterSignature("AMR-WB Audio", "", 0, []byte("#!AMR-WB\n"))
	types.RegisterSignature("AAC Audio", "ADTS", 0, []byte{0xff, 0xf1})
	types.RegisterSignature("AAC Audio", "ADTS", 0, []byte{0xff, 0xf9})
	types.RegisterSignature("CAF Audio", "", 0, []byte("caff"))
	types.RegisterSignature("WavPack Audio", "", 0, []byte("wvpk"))
	types.RegisterSignature("Monkey's Audio", "", 0, []byte("MAC "))
	types.RegisterSignature("DSF Audio", "", 0, []byte("DSD "))
	types.RegisterSignature("TAK Audio", "", 0, []byte("tBaK"))
	types.RegisterSignature("VOC Audio", "", 0, []byte("Creative Voice File\x1a"))
	types.RegisterSignature("Musepack Audio", "Stream Version 7", 0, []byte("MP+"))
	types.RegisterSignature("Musepack Audio", "Stream Version 8", 0, []byte("MPCK"))
	types.RegisterSignature("Impulse Tracker Module", "", 0, []byte("IMPM"))
	types.RegisterSignature("FastTracker Module", "", 0, []byte("Extended Module: "))
	types.RegisterSignature("Scream Tracker Module", "", 44, []byte("SCRM"))
}
