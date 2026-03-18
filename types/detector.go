//go:generate go run ../gen/optimizer

package types

type Detector interface {
	Detect(Buffer) *Metadata
}

type DetectFunc func(Buffer) *Metadata

func (df DetectFunc) Detect(b Buffer) *Metadata {
	return df(b)
}

func Detect(name string, data []byte) (*Metadata, error) {
	buf := Buffer(data)

	if meta := detectOptimized(buf); meta != nil {
		meta.File = name

		return meta, nil
	}

	for _, d := range detectors {
		if meta := d.Detect(buf); meta != nil {
			meta.File = name

			return meta, nil
		}
	}

	return nil, ErrUnknownFormat
}
