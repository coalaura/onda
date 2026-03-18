package types

import "encoding/binary"

type Buffer []byte

func (b Buffer) Len() int {
	return len(b)
}

func (b Buffer) Has(offset int, magic []byte) bool {
	if offset < 0 || offset+len(magic) > len(b) {
		return false
	}

	return string(b[offset:offset+len(magic)]) == string(magic)
}

func (b Buffer) HasMask(offset int, magic string, mask string) bool {
	if offset < 0 || offset+len(magic) > len(b) || len(magic) != len(mask) {
		return false
	}

	target := b[offset : offset+len(magic)]

	// BCE pog
	if len(magic) > 0 {
		_ = target[len(magic)-1]
		_ = mask[len(magic)-1]
		_ = magic[len(magic)-1]
	}

	for i := 0; i < len(magic); i++ {
		if target[i]&mask[i] != magic[i] {
			return false
		}
	}

	return true
}

func (b Buffer) U16LE(offset int) (uint16, bool) {
	if offset < 0 || offset+2 > len(b) {
		return 0, false
	}

	return binary.LittleEndian.Uint16(b[offset:]), true
}

func (b Buffer) U16BE(offset int) (uint16, bool) {
	if offset < 0 || offset+2 > len(b) {
		return 0, false
	}

	return binary.BigEndian.Uint16(b[offset:]), true
}

func (b Buffer) U32LE(offset int) (uint32, bool) {
	if offset < 0 || offset+4 > len(b) {
		return 0, false
	}

	return binary.LittleEndian.Uint32(b[offset:]), true
}

func (b Buffer) U32BE(offset int) (uint32, bool) {
	if offset < 0 || offset+4 > len(b) {
		return 0, false
	}

	return binary.BigEndian.Uint32(b[offset:]), true
}
