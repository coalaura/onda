package types

type Detector interface {
	Detect(Buffer) *Metadata
}

type DetectFunc func(Buffer) *Metadata

type Signature struct {
	Name   string
	Type   string
	Offset int
	Magic  []byte
	Mask   []byte
}

var detectors []Detector

func Register(d Detector) {
	detectors = append(detectors, d)
}

func RegisterSignature(name string, typ string, offset int, magic []byte) {
	Register(Signature{
		Name:   name,
		Type:   typ,
		Offset: offset,
		Magic:  magic,
	})
}

func RegisterMaskedSignature(name string, typ string, offset int, magic []byte, mask []byte) {
	Register(Signature{
		Name:   name,
		Type:   typ,
		Offset: offset,
		Magic:  magic,
		Mask:   mask,
	})
}

func Detect(name string, data []byte) (*Metadata, error) {
	buf := Buffer(data)

	for _, d := range detectors {
		meta := d.Detect(buf)
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
		Name: s.Name,
		Type: s.Type,
	}
}
