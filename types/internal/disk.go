package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature("QCOW Disk Image", "QCOW2", 0, []byte("QFI\xfb"))
	types.RegisterSignature("VHDX Disk Image", "", 0, []byte("vhdxfile"))
	types.RegisterSignature("VMware Disk Image", "VMDK", 0, []byte("KDMV"))
	types.RegisterSignature("Android Sparse Image", "", 0, []byte{0x3a, 0xff, 0x26, 0xed})

	types.RegisterSignature("VirtualBox Disk Image", "", 0x40, []byte("<<< Oracle VM VirtualBox Disk Image >>>\n"))

	types.RegisterSignature("ISO 9660 Image", "", 0x8001, []byte("CD001"))
}
