package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature("PNG Image", "", 0, []byte{
		137, 'P', 'N', 'G', 13, 10, 26, 10,
	})

	types.RegisterSignature("JPEG Image", "", 0, []byte{
		0xff, 0xd8, 0xff,
	})

	types.RegisterSignature("GIF Image", "GIF87a", 0, []byte("GIF87a"))
	types.RegisterSignature("GIF Image", "GIF89a", 0, []byte("GIF89a"))

	types.RegisterSignature("BMP Image", "", 0, []byte("BM"))

	types.RegisterSignature("TIFF Image", "Little-Endian", 0, []byte{
		'I', 'I', 42, 0,
	})

	types.RegisterSignature("TIFF Image", "Big-Endian", 0, []byte{
		'M', 'M', 0, 42,
	})

	types.RegisterSignature("QOI Image", "", 0, []byte("qoif"))

	types.RegisterSignature("Photoshop Document", "PSD", 0, []byte("8BPS"))
	types.RegisterSignature("Photoshop Document", "PSB", 0, []byte("8BPB"))

	types.RegisterSignature("Radiance HDR Image", "", 0, []byte("#?RADIANCE"))
	types.RegisterSignature("Radiance HDR Image", "", 0, []byte("#?RGBE"))

	types.RegisterSignature("OpenEXR Image", "", 0, []byte{
		0x76, 0x2f, 0x31, 0x01,
	})

	types.RegisterSignature("DDS Image", "", 0, []byte("DDS "))

	types.RegisterSignature("JPEG XL Image", "Codestream", 0, []byte{
		0xff, 0x0a,
	})

	types.RegisterSignature("JPEG XL Image", "Container", 0, []byte{
		0x00, 0x00, 0x00, 0x0c, 'J', 'X', 'L', ' ', 0x0d, 0x0a, 0x87, 0x0a,
	})

	types.RegisterSignature("KTX Texture", "KTX", 0, []byte{
		0xab, 'K', 'T', 'X', ' ', '1', '1', 0xbb, 0x0d, 0x0a, 0x1a, 0x0a,
	})

	types.RegisterSignature("KTX Texture", "KTX2", 0, []byte{
		0xab, 'K', 'T', 'X', ' ', '2', '0', 0xbb, 0x0d, 0x0a, 0x1a, 0x0a,
	})

	types.RegisterSignature("ICNS Icon", "", 0, []byte("icns"))
	types.RegisterSignature("Farbfeld Image", "", 0, []byte("farbfeld"))

	types.RegisterSignature("JPEG 2000 Image", "", 0, []byte{
		0x00, 0x00, 0x00, 0x0c, 'j', 'P', ' ', ' ', 0x0d, 0x0a, 0x87, 0x0a,
	})

	types.RegisterSignature("BPG Image", "", 0, []byte{'B', 'P', 'G', 0xfb})
	types.RegisterSignature("JPEG XR Image", "Little-Endian", 0, []byte{'I', 'I', 0xbc, 0x01})
	types.RegisterSignature("JPEG XR Image", "Big-Endian", 0, []byte{'M', 'M', 0x01, 0xbc})
	types.RegisterSignature("Sun Raster Image", "", 0, []byte{0x59, 0xa6, 0x6a, 0x95})
}
