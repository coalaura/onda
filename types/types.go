//go:generate go run ../gen/sorter

package types

import (
	"errors"
	"io"
)

type Sniffer interface {
	Sniff(io.ReadSeeker) (*Metadata, error)
}

type Metadata struct {
	File string

	// Kind is the primary format family (for example: ZIP Archive).
	// Keep this as the broad, stable classification.
	Kind KindID

	// Type is the optional subtype within the Kind (for example: Microsoft Word Document (DOCX)).
	// Use TypeNone when there is no meaningful subtype.
	Type TypeID
}

var ErrUnknownFormat = errors.New("unknown file format")
