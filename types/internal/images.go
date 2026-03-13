package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature("PNG image", "", 0, []byte{
		137, 'P', 'N', 'G', 13, 10, 26, 10,
	})

	types.RegisterSignature("JPEG image", "", 0, []byte{
		0xff, 0xd8, 0xff,
	})

	types.RegisterSignature("GIF image", "GIF87a", 0, []byte("GIF87a"))
	types.RegisterSignature("GIF image", "GIF89a", 0, []byte("GIF89a"))

	types.RegisterSignature("BMP image", "", 0, []byte("BM"))

	types.RegisterSignature("TIFF image", "little-endian", 0, []byte{
		'I', 'I', 42, 0,
	})

	types.RegisterSignature("TIFF image", "big-endian", 0, []byte{
		'M', 'M', 0, 42,
	})

	types.RegisterSignature("QOI image", "", 0, []byte("qoif"))

	types.RegisterSignature("Photoshop document", "PSD", 0, []byte("8BPS"))
	types.RegisterSignature("Photoshop document", "PSB", 0, []byte("8BPB"))

	types.RegisterSignature("Radiance HDR image", "", 0, []byte("#?RADIANCE"))
	types.RegisterSignature("Radiance HDR image", "", 0, []byte("#?RGBE"))

	types.RegisterSignature("OpenEXR image", "", 0, []byte{
		0x76, 0x2f, 0x31, 0x01,
	})

	types.RegisterSignature("DDS image", "", 0, []byte("DDS "))

	types.RegisterSignature("JPEG XL image", "codestream", 0, []byte{
		0xff, 0x0a,
	})

	types.RegisterSignature("JPEG XL image", "container", 0, []byte{
		0x00, 0x00, 0x00, 0x0c, 'J', 'X', 'L', ' ', 0x0d, 0x0a, 0x87, 0x0a,
	})

	types.RegisterSignature("KTX texture", "KTX", 0, []byte{
		0xab, 'K', 'T', 'X', ' ', '1', '1', 0xbb, 0x0d, 0x0a, 0x1a, 0x0a,
	})

	types.RegisterSignature("KTX texture", "KTX2", 0, []byte{
		0xab, 'K', 'T', 'X', ' ', '2', '0', 0xbb, 0x0d, 0x0a, 0x1a, 0x0a,
	})

	types.RegisterSignature("ICNS icon", "", 0, []byte("icns"))
	types.RegisterSignature("Farbfeld image", "", 0, []byte("farbfeld"))

	types.RegisterSignature("JPEG 2000 image", "", 0, []byte{
		0x00, 0x00, 0x00, 0x0c, 'j', 'P', ' ', ' ', 0x0d, 0x0a, 0x87, 0x0a,
	})

	types.RegisterSignature("BPG image", "", 0, []byte{'B', 'P', 'G', 0xfb})
	types.RegisterSignature("JPEG XR image", "little-endian", 0, []byte{'I', 'I', 0xbc, 0x01})
	types.RegisterSignature("JPEG XR image", "big-endian", 0, []byte{'M', 'M', 0x01, 0xbc})
	types.RegisterSignature("Sun raster image", "", 0, []byte{0x59, 0xa6, 0x6a, 0x95})
}
