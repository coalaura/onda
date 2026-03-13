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
}
