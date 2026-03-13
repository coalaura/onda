package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature("QCOW disk image", "QCOW2", 0, []byte("QFI\xfb"))
	types.RegisterSignature("VHDX disk image", "", 0, []byte("vhdxfile"))
	types.RegisterSignature("VMware disk image", "VMDK", 0, []byte("KDMV"))
	types.RegisterSignature("Android sparse image", "", 0, []byte{0x3a, 0xff, 0x26, 0xed})

	types.RegisterSignature("VirtualBox disk image", "", 0x40, []byte("<<< Oracle VM VirtualBox Disk Image >>>\n"))

	types.RegisterSignature("ISO 9660 image", "", 0x8001, []byte("CD001"))
}
