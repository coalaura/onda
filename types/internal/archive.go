package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature("ZIP archive", "", 0, []byte{'P', 'K', 3, 4})
	types.RegisterSignature("ZIP archive", "empty", 0, []byte{'P', 'K', 5, 6})
	types.RegisterSignature("ZIP archive", "spanned", 0, []byte{'P', 'K', 7, 8})

	types.RegisterSignature("7-Zip archive", "", 0, []byte{
		'7', 'z', 0xbc, 0xaf, 0x27, 0x1c,
	})

	types.RegisterSignature("RAR archive", "RAR4", 0, []byte{
		'R', 'a', 'r', '!', 0x1a, 0x07, 0x00,
	})

	types.RegisterSignature("RAR archive", "RAR5", 0, []byte{
		'R', 'a', 'r', '!', 0x1a, 0x07, 0x01, 0x00,
	})

	types.RegisterSignature("Gzip archive", "", 0, []byte{
		0x1f, 0x8b,
	})

	types.RegisterSignature("Bzip2 archive", "", 0, []byte("BZh"))

	types.RegisterSignature("XZ archive", "", 0, []byte{
		0xfd, '7', 'z', 'X', 'Z', 0x00,
	})

	types.RegisterSignature("Zstandard archive", "", 0, []byte{
		0x28, 0xb5, 0x2f, 0xfd,
	})

	types.RegisterSignature("LZ4 frame", "", 0, []byte{
		0x04, 0x22, 0x4d, 0x18,
	})

	types.RegisterSignature("LZIP archive", "", 0, []byte("LZIP"))
	types.RegisterSignature("LZOP archive", "", 0, []byte{0x89, 'L', 'Z', 'O', 0x00, 0x0d, 0x0a, 0x1a, 0x0a})

	types.RegisterSignature("Snappy framed data", "", 0, []byte{
		0xff, 0x06, 0x00, 0x00, 's', 'N', 'a', 'P', 'p', 'Y',
	})

	types.RegisterSignature("Cabinet archive", "", 0, []byte("MSCF"))
	types.RegisterSignature("XAR archive", "", 0, []byte("xar!"))
	types.RegisterSignature("Windows Imaging Format", "", 0, []byte("MSWIM\x00\x00\x00"))
	types.RegisterSignature("AR archive", "", 0, []byte("!<arch>\n"))
	types.RegisterSignature("Debian package", "", 8, []byte("debian-binary"))
	types.RegisterSignature("SquashFS filesystem", "", 0, []byte("hsqs"))

	types.RegisterSignature("CPIO archive", "new ASCII", 0, []byte("070701"))
	types.RegisterSignature("CPIO archive", "new ASCII with CRC", 0, []byte("070702"))
	types.RegisterSignature("CPIO archive", "old ASCII", 0, []byte("070707"))
	types.RegisterSignature("CPIO archive", "binary big-endian", 0, []byte{0x71, 0xc7})
	types.RegisterSignature("CPIO archive", "binary little-endian", 0, []byte{0xc7, 0x71})

	types.RegisterSignature("RPM package", "", 0, []byte{
		0xed, 0xab, 0xee, 0xdb,
	})
}
