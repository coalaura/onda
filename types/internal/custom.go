package internal

import (
	"github.com/coalaura/onda/types"
	"github.com/coalaura/onda/types/custom"
)

func init() {
	types.Register(types.DetectFunc(custom.DetectAppleDiskImage))
	types.Register(types.DetectFunc(custom.DetectEBML))
	types.Register(types.DetectFunc(custom.DetectELF))
	types.Register(types.DetectFunc(custom.DetectExt))
	types.Register(types.DetectFunc(custom.DetectFAT))
	types.Register(types.DetectFunc(custom.DetectISOBaseMedia))
	types.Register(types.DetectFunc(custom.DetectMachO))
	types.Register(types.DetectFunc(custom.DetectMPEGAudioFrames))
	types.Register(types.DetectFunc(custom.DetectNetpbm))
	types.Register(types.DetectFunc(custom.DetectOgg))
	types.Register(types.DetectFunc(custom.DetectOLE))
	types.Register(types.DetectFunc(custom.DetectPE))
	types.Register(types.DetectFunc(custom.DetectPKCS12))
	types.Register(types.DetectFunc(custom.DetectPYC))
	types.Register(types.DetectFunc(custom.DetectSVG))
	types.Register(types.DetectFunc(custom.DetectTar))
	types.Register(types.DetectFunc(custom.DetectTIFFSubtypes))
	types.Register(types.DetectFunc(custom.DetectTorrent))
	types.Register(types.DetectFunc(custom.DetectZIPContainer))

	types.RegisterWeak(types.DetectFunc(custom.DetectLZMA))
	types.RegisterWeak(types.DetectFunc(custom.DetectMP3))
	types.RegisterWeak(types.DetectFunc(custom.DetectMPEGTransport))
	types.RegisterWeak(types.DetectFunc(custom.DetectPCX))
	types.RegisterWeak(types.DetectFunc(custom.DetectZlib))
}
