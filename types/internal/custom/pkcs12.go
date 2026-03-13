package custom

import "github.com/coalaura/onda/types"

func DetectPKCS12(b types.Buffer) *types.Metadata {
	if b.Len() < 32 {
		return nil
	}

	if b[0] != 0x30 {
		return nil
	}

	if !containsASN1OID(b, []byte{0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x01, 0x07, 0x01}) {
		return nil
	}

	return &types.Metadata{Kind: types.KindPKCS12}
}

func containsASN1OID(b types.Buffer, oid []byte) bool {
	if len(oid) == 0 || b.Len() < len(oid) {
		return false
	}

	limit := min(b.Len(), 4096) - len(oid)
	for i := 0; i <= limit; i++ {
		if b.Has(i, oid) {
			return true
		}
	}

	return false
}
