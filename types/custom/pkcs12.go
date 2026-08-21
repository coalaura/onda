package custom

import "github.com/coalaura/wtf/types"

func DetectPKCS12(b types.Buffer) *types.Metadata {
	if b.Len() < 32 {
		return nil
	}

	offset, length, ok := asn1Value(b, 0, 0x30)
	if !ok {
		return nil
	}

	end := offset + length

	versionOffset, versionLength, ok := asn1Value(b, offset, 0x02)
	if !ok || versionLength != 1 || versionOffset+versionLength > end || b[versionOffset] != 3 {
		return nil
	}

	offset = versionOffset + versionLength

	contentOffset, contentLength, ok := asn1Value(b, offset, 0x30)
	if !ok || contentOffset+contentLength > end {
		return nil
	}

	contentEnd := contentOffset + contentLength

	oidOffset, oidLength, ok := asn1Value(b, contentOffset, 0x06)
	if !ok || oidLength != 9 || oidOffset+oidLength > contentEnd || !b.Has(oidOffset, []byte{0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x01, 0x07, 0x01}) {
		return nil
	}

	return &types.Metadata{Kind: types.KindPKCS12}
}

func asn1Value(b types.Buffer, offset int, tag byte) (int, int, bool) {
	if offset < 0 || offset+2 > b.Len() || b[offset] != tag {
		return 0, 0, false
	}

	length := int(b[offset+1])
	headerSize := 2

	if length&0x80 != 0 {
		lengthBytes := length & 0x7f
		if lengthBytes == 0 || lengthBytes > 4 || offset+2+lengthBytes > b.Len() {
			return 0, 0, false
		}

		length = 0

		for i := range lengthBytes {
			length = length<<8 | int(b[offset+2+i])
		}

		headerSize += lengthBytes
	}

	valueOffset := offset + headerSize
	if length < 0 || valueOffset > b.Len() || length > b.Len()-valueOffset {
		return 0, 0, false
	}

	return valueOffset, length, true
}
