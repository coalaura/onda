package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature("FLAC audio", "", 0, []byte("fLaC"))
	types.RegisterSignature("MIDI sequence", "", 0, []byte("MThd"))
	types.RegisterSignature("Ogg container", "", 0, []byte("OggS"))
	types.RegisterSignature("AU audio", "", 0, []byte(".snd"))
	types.RegisterSignature("AMR audio", "", 0, []byte("#!AMR\n"))
	types.RegisterSignature("AMR-WB audio", "", 0, []byte("#!AMR-WB\n"))
	types.RegisterSignature("AAC audio", "ADTS", 0, []byte{0xff, 0xf1})
	types.RegisterSignature("AAC audio", "ADTS", 0, []byte{0xff, 0xf9})
	types.RegisterSignature("CAF audio", "", 0, []byte("caff"))
	types.RegisterSignature("WavPack audio", "", 0, []byte("wvpk"))
	types.RegisterSignature("Monkey's Audio", "", 0, []byte("MAC "))
	types.RegisterSignature("DSF audio", "", 0, []byte("DSD "))
	types.RegisterSignature("TAK audio", "", 0, []byte("tBaK"))
	types.RegisterSignature("VOC audio", "", 0, []byte("Creative Voice File\x1a"))
	types.RegisterSignature("Musepack audio", "stream version 7", 0, []byte("MP+"))
	types.RegisterSignature("Musepack audio", "stream version 8", 0, []byte("MPCK"))
	types.RegisterSignature("Impulse Tracker module", "", 0, []byte("IMPM"))
	types.RegisterSignature("FastTracker module", "", 0, []byte("Extended Module: "))
	types.RegisterSignature("Scream Tracker module", "", 44, []byte("SCRM"))
}
