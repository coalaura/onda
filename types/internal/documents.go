package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature(types.KindPDFDocument, types.TypeNone, 0, []byte("%PDF-"))
	types.RegisterSignature(types.KindPostScriptDocument, types.TypeNone, 0, []byte("%!PS-Adobe-"))
	types.RegisterSignature(types.KindRichTextFormatDocument, types.TypeNone, 0, []byte("{\\rtf"))
	types.RegisterSignature(types.KindSQLiteDatabase, types.TypeNone, 0, []byte("SQLite format 3\x00"))
	types.RegisterSignature(types.KindHDF5Data, types.TypeNone, 0, []byte{0x89, 'H', 'D', 'F', 0x0d, 0x0a, 0x1a, 0x0a})
	types.RegisterSignature(types.KindPGPMessage, types.TypeNone, 0, []byte("-----BEGIN PGP"))
	types.RegisterSignature(types.KindApacheParquet, types.TypeNone, 0, []byte("PAR1"))
	types.RegisterSignature(types.KindApacheArrowFile, types.TypeNone, 0, []byte("ARROW1"))
	types.RegisterSignature(types.KindFITSAstronomicalImage, types.TypeNone, 0, []byte("SIMPLE  ="))
	types.RegisterSignature(types.KindDICOMMedicalImage, types.TypeNone, 128, []byte("DICM"))
	types.RegisterSignature(types.KindWindowsRegistryHive, types.TypeNone, 0, []byte("regf"))
	types.RegisterSignature(types.KindAppleBinaryPropertyList, types.TypeNone, 0, []byte("bplist00"))
	types.RegisterSignature(types.KindKeePassDatabase, types.TypeKDBX, 0, []byte{0x03, 0xd9, 0xa2, 0x9a, 0x67, 0xfb, 0x4b, 0xb5})

	types.RegisterSignature(types.KindAvroObjectContainer, types.TypeNone, 0, []byte{
		'O', 'b', 'j', 0x01,
	})

	types.RegisterSignature(types.KindORCColumnarData, types.TypeNone, 0, []byte("ORC"))
	types.RegisterSignature(types.KindWARCFile, types.TypeNone, 0, []byte("WARC/"))
}
