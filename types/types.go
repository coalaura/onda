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
	Info string // Optional extra information
}

var (
	ErrUnknownFormat = errors.New("unknown file format")

	registry []Sniffer
)

func Register(sn Sniffer) {
	registry = append(registry, sn)
}
