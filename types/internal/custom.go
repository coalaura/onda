package internal

import (
	"github.com/coalaura/wtf/types"
	"github.com/coalaura/wtf/types/custom"
)

func init() {
	types.Register(types.DetectFunc(custom.Detect3DStudio))
	types.Register(types.DetectFunc(custom.DetectAC3))
	types.Register(types.DetectFunc(custom.DetectEBML))
	types.Register(types.DetectFunc(custom.DetectELF))
	types.Register(types.DetectFunc(custom.DetectExt))
	types.Register(types.DetectFunc(custom.DetectFAT))
	types.Register(types.DetectFunc(custom.DetectISOBaseMedia))
	types.Register(types.DetectFunc(custom.DetectMachO))
	types.Register(types.DetectFunc(custom.DetectNetpbm))
	types.Register(types.DetectFunc(custom.DetectOgg))
	types.Register(types.DetectFunc(custom.DetectOLE))
	types.Register(types.DetectFunc(custom.DetectPE))
	types.Register(types.DetectFunc(custom.DetectPKCS12))
	types.Register(types.DetectFunc(custom.DetectPYC))
	types.Register(types.DetectFunc(custom.DetectQNX4))
	types.Register(types.DetectFunc(custom.DetectSQLiteSHM))
	types.Register(types.DetectFunc(custom.DetectTar))
	types.Register(types.DetectFunc(custom.DetectTIFFSubtypes))
	types.Register(types.DetectFunc(custom.DetectTorrent))
	types.Register(types.DetectFunc(custom.DetectZIPContainer))

	types.RegisterWeak(types.DetectFunc(custom.DetectAV1))
	types.RegisterWeak(types.DetectFunc(custom.DetectH26x))
	types.RegisterWeak(types.DetectFunc(custom.DetectJSON))
	types.RegisterWeak(types.DetectFunc(custom.DetectLZMA))
	types.RegisterWeak(types.DetectFunc(custom.DetectMP3))
	types.RegisterWeak(types.DetectFunc(custom.DetectMPEGTransport))
	types.RegisterWeak(types.DetectFunc(custom.DetectPCX))
	types.RegisterWeak(types.DetectFunc(custom.DetectXMLSubtypes))
	types.RegisterWeak(types.DetectFunc(custom.DetectZlib))
}
