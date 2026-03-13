package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature(types.KindPNGImage, types.TypeNone, 0, []byte{
		137, 'P', 'N', 'G', 13, 10, 26, 10,
	})

	types.RegisterSignature(types.KindJPEGImage, types.TypeNone, 0, []byte{
		0xff, 0xd8, 0xff,
	})

	types.RegisterSignature(types.KindGIFImage, types.TypeGIF87a, 0, []byte("GIF87a"))
	types.RegisterSignature(types.KindGIFImage, types.TypeGIF89a, 0, []byte("GIF89a"))

	types.RegisterSignature(types.KindBMPImage, types.TypeNone, 0, []byte("BM"))

	types.RegisterSignature(types.KindTIFFImage, types.TypeLittleEndian, 0, []byte{
		'I', 'I', 42, 0,
	})

	types.RegisterSignature(types.KindTIFFImage, types.TypeBigEndian, 0, []byte{
		'M', 'M', 0, 42,
	})

	types.RegisterSignature(types.KindQOIImage, types.TypeNone, 0, []byte("qoif"))

	types.RegisterSignature(types.KindPhotoshopDocument, types.TypePSD, 0, []byte("8BPS"))
	types.RegisterSignature(types.KindPhotoshopDocument, types.TypePSB, 0, []byte("8BPB"))

	types.RegisterSignature(types.KindRadianceHDRImage, types.TypeNone, 0, []byte("#?RADIANCE"))
	types.RegisterSignature(types.KindRadianceHDRImage, types.TypeNone, 0, []byte("#?RGBE"))

	types.RegisterSignature(types.KindOpenEXRImage, types.TypeNone, 0, []byte{
		0x76, 0x2f, 0x31, 0x01,
	})

	types.RegisterSignature(types.KindDDSImage, types.TypeNone, 0, []byte("DDS "))

	types.RegisterSignature(types.KindJPEGXLImage, types.TypeCodestream, 0, []byte{
		0xff, 0x0a,
	})

	types.RegisterSignature(types.KindJPEGXLImage, types.TypeContainer, 0, []byte{
		0x00, 0x00, 0x00, 0x0c, 'J', 'X', 'L', ' ', 0x0d, 0x0a, 0x87, 0x0a,
	})

	types.RegisterSignature(types.KindKTXTexture, types.TypeKTX, 0, []byte{
		0xab, 'K', 'T', 'X', ' ', '1', '1', 0xbb, 0x0d, 0x0a, 0x1a, 0x0a,
	})

	types.RegisterSignature(types.KindKTXTexture, types.TypeKTX2, 0, []byte{
		0xab, 'K', 'T', 'X', ' ', '2', '0', 0xbb, 0x0d, 0x0a, 0x1a, 0x0a,
	})

	types.RegisterSignature(types.KindICNSIcon, types.TypeNone, 0, []byte("icns"))
	types.RegisterSignature(types.KindFarbfeldImage, types.TypeNone, 0, []byte("farbfeld"))

	types.RegisterSignature(types.KindJPEG2000Image, types.TypeNone, 0, []byte{
		0x00, 0x00, 0x00, 0x0c, 'j', 'P', ' ', ' ', 0x0d, 0x0a, 0x87, 0x0a,
	})

	types.RegisterSignature(types.KindBPGImage, types.TypeNone, 0, []byte{'B', 'P', 'G', 0xfb})
	types.RegisterSignature(types.KindJPEGXRImage, types.TypeLittleEndian, 0, []byte{'I', 'I', 0xbc, 0x01})
	types.RegisterSignature(types.KindJPEGXRImage, types.TypeBigEndian, 0, []byte{'M', 'M', 0x01, 0xbc})
	types.RegisterSignature(types.KindSunRasterImage, types.TypeNone, 0, []byte{0x59, 0xa6, 0x6a, 0x95})
	types.RegisterSignature(types.KindGIMPXCFImage, types.TypeNone, 0, []byte("gimp xcf "))
	types.RegisterSignature(types.KindCanonRAWImage, types.TypeNone, 0, []byte{'I', 'I', 0x2a, 0x00, 0x10, 0x00, 0x00, 0x00, 'C', 'R'})
	types.RegisterSignature(types.KindOlympusRAWImage, types.TypeNone, 0, []byte{'I', 'I', 'R', 'O', 0x08, 0x00})
	types.RegisterSignature(types.KindOlympusRAWImage, types.TypeNone, 0, []byte{'M', 'M', 'O', 'R', 0x00, 0x00})
	types.RegisterSignature(types.KindFujifilmRAWImage, types.TypeNone, 0, []byte("FUJIFILMCCD-RAW "))
}
