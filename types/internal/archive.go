package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature(types.Kind7ZipArchive, types.TypeNone, 0, []byte{'7', 'z', 0xbc, 0xaf, 0x27, 0x1c})
	types.RegisterSignature(types.KindARArchive, types.TypeNone, 0, []byte("!<arch>\n"))
	types.RegisterSignature(types.KindBzip2Archive, types.TypeNone, 0, []byte("BZh"))
	types.RegisterSignature(types.KindCabinetArchive, types.TypeNone, 0, []byte("MSCF"))
	types.RegisterSignature(types.KindCPIOArchive, types.TypeBinaryBigEndian, 0, []byte{0x71, 0xc7})
	types.RegisterSignature(types.KindCPIOArchive, types.TypeBinaryLittleEndian, 0, []byte{0xc7, 0x71})
	types.RegisterSignature(types.KindCPIOArchive, types.TypeNewASCII, 0, []byte("070701"))
	types.RegisterSignature(types.KindCPIOArchive, types.TypeNewASCIIWithCRC, 0, []byte("070702"))
	types.RegisterSignature(types.KindCPIOArchive, types.TypeOldASCII, 0, []byte("070707"))
	types.RegisterSignature(types.KindDebianPackage, types.TypeNone, 8, []byte("debian-binary"))
	types.RegisterSignature(types.KindGzipArchive, types.TypeNone, 0, []byte{0x1f, 0x8b})
	types.RegisterSignature(types.KindLZ4Frame, types.TypeNone, 0, []byte{0x04, 0x22, 0x4d, 0x18})
	types.RegisterSignature(types.KindLZIPArchive, types.TypeNone, 0, []byte("LZIP"))
	types.RegisterSignature(types.KindLZOPArchive, types.TypeNone, 0, []byte{0x89, 'L', 'Z', 'O', 0x00, 0x0d, 0x0a, 0x1a, 0x0a})
	types.RegisterSignature(types.KindRARArchive, types.TypeRAR4, 0, []byte{'R', 'a', 'r', '!', 0x1a, 0x07, 0x00})
	types.RegisterSignature(types.KindRARArchive, types.TypeRAR5, 0, []byte{'R', 'a', 'r', '!', 0x1a, 0x07, 0x01, 0x00})
	types.RegisterSignature(types.KindRPMPackage, types.TypeNone, 0, []byte{0xed, 0xab, 0xee, 0xdb})
	types.RegisterSignature(types.KindRubyGemPackage, types.TypeNone, 0, []byte("RubyGems"))
	types.RegisterSignature(types.KindSnappyFramedData, types.TypeNone, 0, []byte{0xff, 0x06, 0x00, 0x00, 's', 'N', 'a', 'P', 'p', 'Y'})
	types.RegisterSignature(types.KindSquashFSFilesystem, types.TypeNone, 0, []byte("hsqs"))
	types.RegisterSignature(types.KindStuffItArchive, types.TypeNone, 0, []byte("SIT!"))
	types.RegisterSignature(types.KindStuffItArchive, types.TypeNone, 0, []byte("SITD"))
	types.RegisterSignature(types.KindUnixCompressArchive, types.TypeNone, 0, []byte{0x1f, 0x9d})
	types.RegisterSignature(types.KindWindowsImagingFormat, types.TypeNone, 0, []byte("MSWIM\x00\x00\x00"))
	types.RegisterSignature(types.KindXARArchive, types.TypeNone, 0, []byte("xar!"))
	types.RegisterSignature(types.KindXZArchive, types.TypeNone, 0, []byte{0xfd, '7', 'z', 'X', 'Z', 0x00})
	types.RegisterSignature(types.KindZIPArchive, types.TypeEmpty, 0, []byte{'P', 'K', 5, 6})
	types.RegisterSignature(types.KindZIPArchive, types.TypeSpanned, 0, []byte{'P', 'K', 7, 8})
	types.RegisterSignature(types.KindZstandardArchive, types.TypeNone, 0, []byte{0x28, 0xb5, 0x2f, 0xfd})

	types.RegisterMaskedSignature(types.KindGzipData, types.TypeBGZF, 0, []byte{0x1f, 0x8b, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 'B', 'C', 0x02, 0x00}, []byte{0xff, 0xff, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 0xff, 0xff})
}
