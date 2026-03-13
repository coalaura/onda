package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature("CRX Browser Extension", "", 0, []byte("Cr24"))
	types.RegisterSignature("Dalvik Executable", "DEX 035", 0, []byte("dex\n035\x00"))
	types.RegisterSignature("Dalvik Executable", "DEX 036", 0, []byte("dex\n036\x00"))
	types.RegisterSignature("Dalvik Executable", "DEX 037", 0, []byte("dex\n037\x00"))
	types.RegisterSignature("Dalvik Executable", "DEX 038", 0, []byte("dex\n038\x00"))
	types.RegisterSignature("Dalvik Executable", "DEX 039", 0, []byte("dex\n039\x00"))
	types.RegisterSignature("Dalvik Executable", "DEX 040", 0, []byte("dex\n040\x00"))
	types.RegisterSignature("Dalvik Executable", "DEX 041", 0, []byte("dex\n041\x00"))
	types.RegisterSignature("Android OAT", "", 0, []byte("oat\n"))
	types.RegisterSignature("Android ODEX", "", 0, []byte("dey\n"))

	types.RegisterSignature("Java Class", "", 0, []byte{
		0xca, 0xfe, 0xba, 0xbe,
	})

	types.RegisterSignature("Java KeyStore", "", 0, []byte{0xfe, 0xed, 0xfe, 0xed})

	types.RegisterSignature("WebAssembly Module", "", 0, []byte{
		0x00, 0x61, 0x73, 0x6d,
	})

	types.RegisterSignature("NES ROM", "", 0, []byte{
		'N', 'E', 'S', 0x1a,
	})

	types.RegisterSignature("PCAP Capture", "Little-Endian", 0, []byte{
		0xd4, 0xc3, 0xb2, 0xa1,
	})

	types.RegisterSignature("PCAP Capture", "Big-Endian", 0, []byte{
		0xa1, 0xb2, 0xc3, 0xd4,
	})

	types.RegisterSignature("PCAP Capture", "Nanosecond Little-Endian", 0, []byte{
		0x4d, 0x3c, 0xb2, 0xa1,
	})

	types.RegisterSignature("PCAP Capture", "Nanosecond Big-Endian", 0, []byte{
		0xa1, 0xb2, 0x3c, 0x4d,
	})

	types.RegisterSignature("PCAPNG Capture", "", 0, []byte{
		0x0a, 0x0d, 0x0d, 0x0a,
	})

	types.RegisterSignature("Shockwave Flash", "Uncompressed", 0, []byte("FWS"))
	types.RegisterSignature("Shockwave Flash", "Zlib Compressed", 0, []byte("CWS"))
	types.RegisterSignature("Shockwave Flash", "LZMA Compressed", 0, []byte("ZWS"))

	types.RegisterSignature("LLVM Bitcode", "", 0, []byte{
		'B', 'C', 0xc0, 0xde,
	})

	types.RegisterSignature("LLVM Bitcode", "Wrapper", 0, []byte{0xde, 0xc0, 0x17, 0x0b})

	types.RegisterSignature("Windows Shortcut", "", 0, []byte{
		0x4c, 0x00, 0x00, 0x00,
		0x01, 0x14, 0x02, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0xc0, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x46,
	})
}
