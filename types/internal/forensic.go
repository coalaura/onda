package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature(types.KindAdvancedForensicFormat, types.TypeNone, 0, []byte("AFF\x00"))
	types.RegisterSignature(types.KindEnCaseImage, types.TypeNone, 0, []byte("EVF\x09\x0d\x0a\xff\x00"))
	types.RegisterSignature(types.KindMCAPCapture, types.TypeNone, 0, []byte{0x89, 'M', 'C', 'A', 'P', '0', '\r', '\n'})
	types.RegisterSignature(types.KindMediaDescriptor, types.TypeNone, 0, []byte("MEDIA DESCRIPTOR"))
	types.RegisterSignature(types.KindMicrosoftNetworkMonitor, types.TypeNone, 0, []byte("MACROSOFT\x00"))
	types.RegisterSignature(types.KindMicrosoftNetworkMonitor, types.TypeNone, 0, []byte("NetMon\x00\x00"))
	types.RegisterSignature(types.KindPowerISODAA, types.TypeNone, 0, []byte("DAA\x1a"))
	types.RegisterSignature(types.KindROSBag, types.TypeNone, 0, []byte("#ROSBAG V2.0\n"))
	types.RegisterSignature(types.KindSnoopCapture, types.TypeNone, 0, []byte("snoop\x00\x00\x00"))
}
