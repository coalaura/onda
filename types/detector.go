package types

import (
	"io"
)

func Detect(name string, r io.ReadSeeker) (*Metadata, error) {
	for _, sn := range registry {
		_, err := r.Seek(0, io.SeekStart)
		if err != nil {
			continue
		}

		meta, err := sn.Sniff(r)
		if err == nil && meta != nil {
			meta.File = name

			return meta, nil
		}
	}

	return nil, ErrUnknownFormat
}
