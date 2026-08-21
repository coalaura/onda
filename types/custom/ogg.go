package custom

import "github.com/coalaura/wtf/types"

func DetectOgg(b types.Buffer) *types.Metadata {
	if b.Len() < 27 || !b.Has(0, []byte("OggS")) {
		return nil
	}

	limit := min(b.Len(), 4096)
	offset := 0

	for offset+27 <= limit {
		if !b.Has(offset, []byte("OggS")) || b[offset+4] != 0 || b[offset+5]&^byte(0x07) != 0 {
			return nil
		}

		segmentCount := int(b[offset+26])
		headerEnd := offset + 27 + segmentCount
		if headerEnd > limit {
			break
		}

		payloadSize := 0

		for _, size := range b[offset+27 : headerEnd] {
			payloadSize += int(size)
		}

		pageEnd := headerEnd + payloadSize
		if pageEnd > limit {
			break
		}

		packetStart := headerEnd
		packetOffset := headerEnd

		for _, size := range b[offset+27 : headerEnd] {
			packetOffset += int(size)
			if size == 255 {
				continue
			}

			if meta := detectOggPacket(b[packetStart:packetOffset]); meta != nil {
				return meta
			}

			packetStart = packetOffset
		}

		offset = pageEnd
	}

	return &types.Metadata{Kind: types.KindOggContainer}
}

func detectOggPacket(packet types.Buffer) *types.Metadata {
	if packet.Has(0, []byte("fishead\x00")) && packet.Len() >= 10 {
		switch uint16(packet[8]) | uint16(packet[9])<<8 {
		case 3:
			return &types.Metadata{Kind: types.KindOggSkeleton, Type: types.TypeOggSkeleton3}
		case 4:
			return &types.Metadata{Kind: types.KindOggSkeleton, Type: types.TypeOggSkeleton4}
		}
	}

	switch {
	case packet.Has(0, []byte("OpusHead")):
		return &types.Metadata{Kind: types.KindOggContainer, Type: types.TypeOpusAudio}
	case packet.Has(0, []byte{0x01, 'v', 'o', 'r', 'b', 'i', 's'}):
		return &types.Metadata{Kind: types.KindOggContainer, Type: types.TypeVorbisAudio}
	case packet.Has(0, []byte("Speex   ")):
		return &types.Metadata{Kind: types.KindOggContainer, Type: types.TypeSpeexAudio}
	case packet.Has(0, []byte{0x80, 't', 'h', 'e', 'o', 'r', 'a'}):
		return &types.Metadata{Kind: types.KindOggContainer, Type: types.TypeTheoraVideo}
	case packet.Has(0, []byte{0x7f, 'F', 'L', 'A', 'C'}):
		return &types.Metadata{Kind: types.KindOggContainer, Type: types.TypeFLACAudio}
	default:
		return nil
	}
}
