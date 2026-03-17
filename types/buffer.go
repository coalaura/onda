package types

import "bytes"

type Buffer []byte

func (b Buffer) Len() int {
	return len(b)
}

func (b Buffer) Has(offset int, magic []byte) bool {
	if offset < 0 || offset+len(magic) > len(b) {
		return false
	}

	return bytes.Equal(b[offset:offset+len(magic)], magic)
}

func (b Buffer) HasMask(offset int, magic []byte, mask []byte) bool {
	if len(magic) != len(mask) {
		return false
	}

	if offset < 0 || offset+len(magic) > len(b) {
		return false
	}

	target := b[offset : offset+len(magic)]

	// BCE pog
	if len(magic) > 0 {
		_ = target[len(magic)-1]
		_ = mask[len(magic)-1]
	}

	for i, m := range magic {
		if target[i]&mask[i] != m&mask[i] {
			return false
		}
	}

	return true
}

func (b Buffer) U16LE(offset int) (uint16, bool) {
	if offset < 0 || offset+2 > len(b) {
		return 0, false
	}

	return uint16(b[offset]) | uint16(b[offset+1])<<8, true
}

func (b Buffer) U16BE(offset int) (uint16, bool) {
	if offset < 0 || offset+2 > len(b) {
		return 0, false
	}

	return uint16(b[offset])<<8 | uint16(b[offset+1]), true
}

func (b Buffer) U32LE(offset int) (uint32, bool) {
	if offset < 0 || offset+4 > len(b) {
		return 0, false
	}

	return uint32(b[offset]) | uint32(b[offset+1])<<8 | uint32(b[offset+2])<<16 | uint32(b[offset+3])<<24, true
}

func (b Buffer) U32BE(offset int) (uint32, bool) {
	if offset < 0 || offset+4 > len(b) {
		return 0, false
	}

	return uint32(b[offset])<<24 | uint32(b[offset+1])<<16 | uint32(b[offset+2])<<8 | uint32(b[offset+3]), true
}
