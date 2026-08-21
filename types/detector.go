//go:generate go run ../gen/optimizer

package types

import "io"

type Detector interface {
	Detect(Buffer) *Metadata
}

type DetectFunc func(Buffer) *Metadata

func (df DetectFunc) Detect(b Buffer) *Metadata {
	return df(b)
}

func Detect(name string, data []byte) (*Metadata, error) {
	buf := Buffer(data)

	meta := DetectAppleDiskImage(buf)
	if meta != nil {
		meta.File = name

		return meta, nil
	}

	meta = detectOptimized(buf)
	if meta != nil {
		meta.File = name

		return meta, nil
	}

	return nil, ErrUnknownFormat
}

// DetectReaderAt detects a file from a bounded prefix and refines containers
// using metadata read at their real file offsets.
func DetectReaderAt(name string, data []byte, r io.ReaderAt, size int64) (*Metadata, error) {
	meta := detectOptimized(Buffer(data))
	if meta != nil {
		switch meta.Kind {
		case KindZIPArchive:
			refined := detectZIPReaderAt(r, size)
			if refined != nil {
				meta = preferZIPMetadata(meta, refined)
			}
		case KindOLECompoundDocument:
			refined := detectOLEReaderAt(r, size)
			if refined != nil {
				meta = refined
			}
		}

		meta.File = name

		return meta, nil
	}

	if size >= 512 {
		tail := make([]byte, 512)

		if !readAtExactly(r, tail, size-int64(len(tail))) {
			return nil, io.ErrUnexpectedEOF
		}

		meta = DetectAppleDiskImage(Buffer(tail))
		if meta != nil {
			meta.File = name

			return meta, nil
		}
	}

	return nil, ErrUnknownFormat
}
