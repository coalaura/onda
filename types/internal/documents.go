package internal

import "github.com/coalaura/onda/types"

func init() {
	types.RegisterSignature("PDF document", "", 0, []byte("%PDF-"))
	types.RegisterSignature("PostScript document", "", 0, []byte("%!PS-Adobe-"))
	types.RegisterSignature("Rich Text Format document", "", 0, []byte("{\\rtf"))
	types.RegisterSignature("SQLite database", "", 0, []byte("SQLite format 3\x00"))
	types.RegisterSignature("PGP message", "", 0, []byte("-----BEGIN PGP"))
	types.RegisterSignature("Apache Parquet", "", 0, []byte("PAR1"))
	types.RegisterSignature("FITS astronomical image", "", 0, []byte("SIMPLE  ="))
	types.RegisterSignature("DICOM medical image", "", 128, []byte("DICM"))
	types.RegisterSignature("Windows registry hive", "", 0, []byte("regf"))
	types.RegisterSignature("Apple binary property list", "", 0, []byte("bplist00"))

	types.RegisterSignature("Avro object container", "", 0, []byte{
		'O', 'b', 'j', 0x01,
	})

	types.RegisterSignature("ORC columnar data", "", 0, []byte("ORC"))
}
