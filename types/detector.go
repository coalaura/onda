package types

import (
	"cmp"
	"slices"
	"sync"
)

type priorityClass int

const (
	classSignature priorityClass = iota
	classCustom
	classWeak
	classFallback
)

type Detector interface {
	Detect(Buffer) *Metadata
}

type DetectFunc func(Buffer) *Metadata

type Signature struct {
	Kind   KindID
	Type   TypeID
	Offset int
	Magic  []byte
	Mask   []byte
}

type registered struct {
	d     Detector
	class priorityClass
	len   int
}

var (
	detectors []registered
	sortOnce  sync.Once
)

// Register adds a standard custom detector.
func Register(d Detector) {
	detectors = append(detectors, registered{d: d, class: classCustom})
}

// RegisterWeak adds a structural detector that is prone to false positives (e.g., LZMA).
func RegisterWeak(d Detector) {
	detectors = append(detectors, registered{d: d, class: classWeak})
}

// RegisterFallback adds a detector of last resort (e.g., Plain Text).
func RegisterFallback(d Detector) {
	detectors = append(detectors, registered{d: d, class: classFallback})
}

func RegisterSignature(kind KindID, typ TypeID, offset int, magic []byte) {
	sig := Signature{Kind: kind, Type: typ, Offset: offset, Magic: magic}

	detectors = append(detectors, registered{d: sig, class: classSignature, len: len(magic)})
}

func RegisterMaskedSignature(kind KindID, typ TypeID, offset int, magic []byte, mask []byte) {
	maskedMagic := make([]byte, len(magic))

	for i := range magic {
		maskedMagic[i] = magic[i] & mask[i]
	}

	sig := Signature{Kind: kind, Type: typ, Offset: offset, Magic: maskedMagic, Mask: mask}

	detectors = append(detectors, registered{d: sig, class: classSignature, len: len(magic)})
}

func Detect(name string, data []byte) (*Metadata, error) {
	sortOnce.Do(func() {
		slices.SortStableFunc(detectors, func(a, b registered) int {
			if a.class != b.class {
				return cmp.Compare(a.class, b.class)
			}

			if a.class == classSignature {
				return cmp.Compare(b.len, a.len)
			}

			return 0
		})
	})

	buf := Buffer(data)

	for _, reg := range detectors {
		if meta := reg.d.Detect(buf); meta != nil {
			meta.File = name

			return meta, nil
		}
	}

	return nil, ErrUnknownFormat
}

func (df DetectFunc) Detect(b Buffer) *Metadata {
	return df(b)
}

func (s Signature) Detect(b Buffer) *Metadata {
	var ok bool

	if len(s.Mask) == 0 {
		ok = b.Has(s.Offset, s.Magic)
	} else {
		ok = b.HasMask(s.Offset, s.Magic, s.Mask)
	}

	if !ok {
		return nil
	}

	return &Metadata{
		Kind: s.Kind,
		Type: s.Type,
	}
}
