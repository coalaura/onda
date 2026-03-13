package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature("CRX browser extension", "", 0, []byte("Cr24"))

	types.RegisterSignature("Java class", "", 0, []byte{
		0xca, 0xfe, 0xba, 0xbe,
	})

	types.RegisterSignature("WebAssembly module", "", 0, []byte{
		0x00, 0x61, 0x73, 0x6d,
	})

	types.RegisterSignature("NES ROM", "", 0, []byte{
		'N', 'E', 'S', 0x1a,
	})

	types.RegisterSignature("PCAP capture", "little-endian", 0, []byte{
		0xd4, 0xc3, 0xb2, 0xa1,
	})

	types.RegisterSignature("PCAP capture", "big-endian", 0, []byte{
		0xa1, 0xb2, 0xc3, 0xd4,
	})

	types.RegisterSignature("PCAP capture", "nanosecond little-endian", 0, []byte{
		0x4d, 0x3c, 0xb2, 0xa1,
	})

	types.RegisterSignature("PCAP capture", "nanosecond big-endian", 0, []byte{
		0xa1, 0xb2, 0x3c, 0x4d,
	})

	types.RegisterSignature("PCAPNG capture", "", 0, []byte{
		0x0a, 0x0d, 0x0d, 0x0a,
	})

	types.RegisterSignature("Shockwave Flash", "uncompressed", 0, []byte("FWS"))
	types.RegisterSignature("Shockwave Flash", "zlib compressed", 0, []byte("CWS"))
	types.RegisterSignature("Shockwave Flash", "lzma compressed", 0, []byte("ZWS"))
}
