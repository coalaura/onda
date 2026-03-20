package custom

import (
	"bytes"

	"github.com/coalaura/onda/types"
)

func DetectXMLSubtypes(b types.Buffer) *types.Metadata {
	limit := min(b.Len(), 1024)
	data := b[:limit]

	trimmed := bytes.TrimSpace(data)
	if len(trimmed) < 5 || trimmed[0] != '<' {
		return nil
	}

	if bytes.HasPrefix(trimmed, []byte("<?xml")) || bytes.HasPrefix(trimmed, []byte("<!DOCTYPE")) {
		if bytes.Contains(data, []byte("<plist")) || bytes.Contains(data, []byte("<!DOCTYPE plist")) {
			return &types.Metadata{Kind: types.KindAppleXMLPropertyList}
		}

		if bytes.Contains(data, []byte("<kml")) {
			return &types.Metadata{Kind: types.KindKeyholeMarkupLanguage}
		}

		if bytes.Contains(data, []byte("<gpx")) {
			return &types.Metadata{Kind: types.KindGPSExchangeFormat}
		}

		if bytes.Contains(data, []byte("<rss")) {
			return &types.Metadata{Kind: types.KindRSSFeed}
		}

		if bytes.Contains(data, []byte("<feed")) {
			return &types.Metadata{Kind: types.KindAtomFeed}
		}

		if bytes.Contains(data, []byte("<soap:Envelope")) {
			return &types.Metadata{Kind: types.KindSOAPMessage}
		}
	}

	return nil
}
