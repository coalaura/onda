//go:build ignore

package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature(types.KindAndroidSparseImage, types.TypeNone, 0, []byte{0x3a, 0xff, 0x26, 0xed})
	types.RegisterSignature(types.KindAPFSFilesystem, types.TypeNone, 32, []byte("NXSB"))
	types.RegisterSignature(types.KindBtrfsFilesystem, types.TypeNone, 0x10040, []byte("_BHRfS_M"))
	types.RegisterSignature(types.KindCramFS, types.TypeNone, 0, []byte{0x45, 0x3d, 0xcd, 0x28})
	types.RegisterSignature(types.KindExFATFilesystem, types.TypeNone, 3, []byte("EXFAT   "))
	types.RegisterSignature(types.KindHFSPlusFilesystem, types.TypeNone, 1024, []byte("H+"))
	types.RegisterSignature(types.KindHFSPlusFilesystem, types.TypeNone, 1024, []byte("HX"))
	types.RegisterSignature(types.KindISO9660Image, types.TypeNone, 0x8001, []byte("CD001"))
	types.RegisterSignature(types.KindLUKSDiskEncryption, types.TypeNone, 0, []byte{'L', 'U', 'K', 'S', 0xba, 0xbe})
	types.RegisterSignature(types.KindNTFSFilesystem, types.TypeNone, 3, []byte("NTFS    "))
	types.RegisterSignature(types.KindQCOWDiskImage, types.TypeQCOW2, 0, []byte("QFI\xfb"))
	types.RegisterSignature(types.KindROMFS, types.TypeNone, 0, []byte("-rom1fs-"))
	types.RegisterSignature(types.KindVHDDiskImage, types.TypeNone, 0, []byte("cxsparse"))
	types.RegisterSignature(types.KindVHDXDiskImage, types.TypeNone, 0, []byte("vhdxfile"))
	types.RegisterSignature(types.KindVirtualBoxDiskImage, types.TypeNone, 0x40, []byte("<<< Oracle VM VirtualBox Disk Image >>>\n"))
	types.RegisterSignature(types.KindVMwareDiskImage, types.TypeVMDK, 0, []byte("KDMV"))
}
