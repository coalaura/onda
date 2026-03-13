package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature("PDF document", "", 0, []byte("%PDF-"))
	types.RegisterSignature("PostScript document", "", 0, []byte("%!PS-Adobe-"))
	types.RegisterSignature("Rich Text Format document", "", 0, []byte("{\\rtf"))
	types.RegisterSignature("SQLite database", "", 0, []byte("SQLite format 3\x00"))
	types.RegisterSignature("HDF5 data", "", 0, []byte{0x89, 'H', 'D', 'F', 0x0d, 0x0a, 0x1a, 0x0a})
	types.RegisterSignature("PGP message", "", 0, []byte("-----BEGIN PGP"))
	types.RegisterSignature("Apache Parquet", "", 0, []byte("PAR1"))
	types.RegisterSignature("Apache Arrow file", "", 0, []byte("ARROW1"))
	types.RegisterSignature("FITS astronomical image", "", 0, []byte("SIMPLE  ="))
	types.RegisterSignature("DICOM medical image", "", 128, []byte("DICM"))
	types.RegisterSignature("Windows registry hive", "", 0, []byte("regf"))
	types.RegisterSignature("Apple binary property list", "", 0, []byte("bplist00"))
	types.RegisterSignature("KeePass database", "KDBX", 0, []byte{0x03, 0xd9, 0xa2, 0x9a, 0x67, 0xfb, 0x4b, 0xb5})

	types.RegisterSignature("Avro object container", "", 0, []byte{
		'O', 'b', 'j', 0x01,
	})

	types.RegisterSignature("ORC columnar data", "", 0, []byte("ORC"))
}
