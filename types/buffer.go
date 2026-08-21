package types

import (
	"bytes"
	"encoding/binary"
)

type Buffer []byte

func (b Buffer) Len() int {
	return len(b)
}

func (b Buffer) Has(offset int, magic []byte) bool {
	if offset < 0 || offset > len(b)-len(magic) {
		return false
	}

	return bytes.Equal(b[offset:offset+len(magic)], magic)
}

func (b Buffer) HasMask(offset int, magic string, mask string) bool {
	if offset < 0 || len(magic) != len(mask) || offset > len(b)-len(magic) {
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
		if target[i]&mask[i] != magic[i]&mask[i] {
			return false
		}
	}

	return true
}

func (b Buffer) U16LE(offset int) (uint16, bool) {
	if offset < 0 || offset > len(b)-2 {
		return 0, false
	}

	return binary.LittleEndian.Uint16(b[offset:]), true
}

func (b Buffer) U16BE(offset int) (uint16, bool) {
	if offset < 0 || offset > len(b)-2 {
		return 0, false
	}

	return binary.BigEndian.Uint16(b[offset:]), true
}

func (b Buffer) U32LE(offset int) (uint32, bool) {
	if offset < 0 || offset > len(b)-4 {
		return 0, false
	}

	return binary.LittleEndian.Uint32(b[offset:]), true
}

func (b Buffer) U32BE(offset int) (uint32, bool) {
	if offset < 0 || offset > len(b)-4 {
		return 0, false
	}

	return binary.BigEndian.Uint32(b[offset:]), true
}
