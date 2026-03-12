package mp3

import (
	"fmt"
	"io"

	"github.com/coalaura/onda/types"
)

type sniffer struct{}

func init() {
	types.Register(sniffer{})
}

func (sniffer) Sniff(r io.ReadSeeker) (*types.Metadata, error) {
	var hdr [3]byte

	_, err := io.ReadFull(r, hdr[:])
	if err != nil {
		return nil, err
	}

	if hdr == [3]byte{'I', 'D', '3'} {
		return &types.Metadata{
			Name: "MP3 audio",
			Type: "ID3 tagged",
		}, nil
	}

	_, err = r.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	var frame [2]byte

	_, err = io.ReadFull(r, frame[:])
	if err != nil {
		return nil, err
	}

	if frame[0] == 0xff && (frame[1]&0xe0) == 0xe0 {
		return &types.Metadata{
			Name: "MP3 audio",
			Info: fmt.Sprintf("frame sync 0x%02x%02x", frame[0], frame[1]),
		}, nil
	}

	return nil, types.ErrUnknownFormat
}
