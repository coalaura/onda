package types

import (
	"errors"
	"io"
)

type Sniffer interface {
	Sniff(io.ReadSeeker) (*Metadata, error)
}

type Metadata struct {
	File string // File name
	Name string // Main human-readable type
	Type string // Optional subtype
}

var ErrUnknownFormat = errors.New("unknown file format")
