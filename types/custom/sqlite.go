package custom

import (
	"bytes"

	"github.com/coalaura/onda/types"
)

func DetectSQLiteSHM(b types.Buffer) *types.Metadata {
	if b.Len() < 96 {
		return nil
	}

	if !bytes.Equal(b[:48], b[48:96]) {
		return nil
	}

	var typ types.TypeID

	if b[0] == 0x18 {
		typ = types.TypeLittleEndian
	} else if b[3] == 0x18 {
		typ = types.TypeBigEndian
	} else {
		return nil
	}

	return &types.Metadata{
		Kind: types.KindSQLite3SharedMemory,
		Type: typ,
	}
}
