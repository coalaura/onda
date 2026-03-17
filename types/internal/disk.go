//go:build ignore

package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature(types.KindAndroidSparseImage, types.TypeNone, 0, []byte{0x3a, 0xff, 0x26, 0xed})
	types.RegisterSignature(types.KindAPFSFilesystem, types.TypeNone, 32, []byte("NXSB"))
	types.RegisterSignature(types.KindBcachefsFilesystem, types.TypeNone, 4096, []byte{0xc6, 0x85, 0x73, 0xf6, 0x4e, 0x1a, 0x45, 0xca, 0x82, 0x65, 0xf5, 0x7f, 0x48, 0xba, 0x6d, 0x81})
	types.RegisterSignature(types.KindBitLockerDiskEncryption, types.TypeNone, 3, []byte("-FVE-FS-"))
	types.RegisterSignature(types.KindBtrfsFilesystem, types.TypeNone, 0x10040, []byte("_BHRfS_M"))
	types.RegisterSignature(types.KindCramFS, types.TypeNone, 0, []byte{0x45, 0x3d, 0xcd, 0x28})
	types.RegisterSignature(types.KindExFATFilesystem, types.TypeNone, 3, []byte("EXFAT   "))
	types.RegisterSignature(types.KindF2FSFilesystem, types.TypeNone, 1024, []byte{0x10, 0x20, 0xf5, 0xf2})
	types.RegisterSignature(types.KindHFSPlusFilesystem, types.TypeNone, 1024, []byte("H+"))
	types.RegisterSignature(types.KindHFSPlusFilesystem, types.TypeNone, 1024, []byte("HX"))
	types.RegisterSignature(types.KindISO9660Image, types.TypeNone, 0x8001, []byte("CD001"))
	types.RegisterSignature(types.KindJFFS2Filesystem, types.TypeNone, 0, []byte{0x85, 0x19})
	types.RegisterSignature(types.KindLUKSDiskEncryption, types.TypeNone, 0, []byte{'L', 'U', 'K', 'S', 0xba, 0xbe})
	types.RegisterSignature(types.KindNILFS2Filesystem, types.TypeNone, 1032, []byte{0x02, 0x09, 0x01, 0x12})
	types.RegisterSignature(types.KindNTFSFilesystem, types.TypeNone, 3, []byte("NTFS    "))
	types.RegisterSignature(types.KindQCOWDiskImage, types.TypeQCOW2, 0, []byte("QFI\xfb"))
	types.RegisterSignature(types.KindROMFS, types.TypeNone, 0, []byte("-rom1fs-"))
	types.RegisterSignature(types.KindUBIFSFilesystem, types.TypeNone, 0, []byte{0x31, 0x18, 0x10, 0x06})
	types.RegisterSignature(types.KindUniversalDiskFormat, types.TypeNone, 32769, []byte("NSR02"))
	types.RegisterSignature(types.KindUniversalDiskFormat, types.TypeNone, 32769, []byte("NSR03"))
	types.RegisterSignature(types.KindVHDDiskImage, types.TypeNone, 0, []byte("cxsparse"))
	types.RegisterSignature(types.KindVHDXDiskImage, types.TypeNone, 0, []byte("vhdxfile"))
	types.RegisterSignature(types.KindVirtualBoxDiskImage, types.TypeNone, 0x40, []byte("<<< Oracle VM VirtualBox Disk Image >>>\n"))
	types.RegisterSignature(types.KindVMwareDiskImage, types.TypeVMDK, 0, []byte("KDMV"))
	types.RegisterSignature(types.KindZFSFilesystem, types.TypeNone, 0, []byte{0x00, 0xba, 0xb1, 0x0c})
	types.RegisterSignature(types.KindZFSFilesystem, types.TypeNone, 0, []byte{0x0c, 0xb1, 0xba, 0x00})
}
