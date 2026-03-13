package types

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

var detectors []Detector
var fallbackDetectors []Detector

func Register(d Detector) {
	detectors = append(detectors, d)
}

func RegisterFallback(d Detector) {
	fallbackDetectors = append(fallbackDetectors, d)
}

func RegisterSignature(kind KindID, typ TypeID, offset int, magic []byte) {
	Register(Signature{
		Kind:   kind,
		Type:   typ,
		Offset: offset,
		Magic:  magic,
	})
}

func RegisterMaskedSignature(kind KindID, typ TypeID, offset int, magic []byte, mask []byte) {
	Register(Signature{
		Kind:   kind,
		Type:   typ,
		Offset: offset,
		Magic:  magic,
		Mask:   mask,
	})
}

func Detect(name string, data []byte) (*Metadata, error) {
	buf := Buffer(data)

	for i := len(detectors) - 1; i >= 0; i-- {
		meta := detectors[i].Detect(buf)
		if meta == nil {
			continue
		}

		meta.File = name

		return meta, nil
	}

	for i := len(fallbackDetectors) - 1; i >= 0; i-- {
		meta := fallbackDetectors[i].Detect(buf)
		if meta == nil {
			continue
		}

		meta.File = name

		return meta, nil
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
