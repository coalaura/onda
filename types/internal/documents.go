package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature("PDF Document", "", 0, []byte("%PDF-"))
	types.RegisterSignature("PostScript Document", "", 0, []byte("%!PS-Adobe-"))
	types.RegisterSignature("Rich Text Format Document", "", 0, []byte("{\\rtf"))
	types.RegisterSignature("SQLite Database", "", 0, []byte("SQLite format 3\x00"))
	types.RegisterSignature("HDF5 Data", "", 0, []byte{0x89, 'H', 'D', 'F', 0x0d, 0x0a, 0x1a, 0x0a})
	types.RegisterSignature("PGP Message", "", 0, []byte("-----BEGIN PGP"))
	types.RegisterSignature("Apache Parquet", "", 0, []byte("PAR1"))
	types.RegisterSignature("Apache Arrow File", "", 0, []byte("ARROW1"))
	types.RegisterSignature("FITS Astronomical Image", "", 0, []byte("SIMPLE  ="))
	types.RegisterSignature("DICOM Medical Image", "", 128, []byte("DICM"))
	types.RegisterSignature("Windows Registry Hive", "", 0, []byte("regf"))
	types.RegisterSignature("Apple Binary Property List", "", 0, []byte("bplist00"))
	types.RegisterSignature("KeePass Database", "KDBX", 0, []byte{0x03, 0xd9, 0xa2, 0x9a, 0x67, 0xfb, 0x4b, 0xb5})

	types.RegisterSignature("Avro Object Container", "", 0, []byte{
		'O', 'b', 'j', 0x01,
	})

	types.RegisterSignature("ORC Columnar Data", "", 0, []byte("ORC"))
}
