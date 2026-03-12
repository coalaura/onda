package wav

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var riff [4]byte

	_, err := io.ReadFull(r, riff[:])
	if err != nil {
		return nil, err
	}

	if riff != [4]byte{'R', 'I', 'F', 'F'} {
		return nil, types.ErrUnknownFormat
	}

	_, err = r.Seek(4, io.SeekCurrent)
	if err != nil {
		return nil, err
	}

	var wave [4]byte

	_, err = io.ReadFull(r, wave[:])
	if err != nil {
		return nil, err
	}

	if wave != [4]byte{'W', 'A', 'V', 'E'} {
		return nil, types.ErrUnknownFormat
	}

	var chunk [4]byte

	_, err = io.ReadFull(r, chunk[:])
	if err != nil {
		return nil, err
	}

	if chunk != [4]byte{'f', 'm', 't', ' '} {
		return &types.Metadata{
			Name: "WAV audio",
		}, nil
	}

	_, err = r.Seek(4, io.SeekCurrent)
	if err != nil {
		return nil, err
	}

	var format uint16

	err = binary.Read(r, binary.LittleEndian, &format)
	if err != nil {
		return nil, err
	}

	var channels uint16

	err = binary.Read(r, binary.LittleEndian, &channels)
	if err != nil {
		return nil, err
	}

	var sampleRate uint32

	err = binary.Read(r, binary.LittleEndian, &sampleRate)
	if err != nil {
		return nil, err
	}

	_, err = r.Seek(6, io.SeekCurrent)
	if err != nil {
		return nil, err
	}

	var bits uint16

	err = binary.Read(r, binary.LittleEndian, &bits)
	if err != nil {
		return nil, err
	}

	var codec string

	if format == 1 {
		codec = "PCM"
	} else {
		codec = fmt.Sprintf("codec %d", format)
	}

	return &types.Metadata{
		Name: "WAV audio",
		Type: codec,
		Info: fmt.Sprintf("%d Hz, %d-bit, %d ch", sampleRate, bits, channels),
	}, nil
}
