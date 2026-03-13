package custom

import (
	"bytes"
	"strings"

	"github.com/coalaura/onda/types"
)

func DetectZIPContainer(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte{'P', 'K', 3, 4}) {
		return nil
	}

	fileName, ok := firstZipLocalName(b)
	if !ok {
		return &types.Metadata{Name: "ZIP archive"}
	}

	lowerName := strings.ToLower(fileName)

	if lowerName == "mimetype" {
		if data := firstZipStoredData(b); len(data) > 0 {
			if bytes.Equal(data, []byte("application/epub+zip")) {
				return &types.Metadata{Name: "EPUB document"}
			}

			if bytes.Equal(data, []byte("application/vnd.oasis.opendocument.text")) {
				return &types.Metadata{Name: "OpenDocument text", Type: "ODT"}
			}

			if bytes.Equal(data, []byte("application/vnd.oasis.opendocument.spreadsheet")) {
				return &types.Metadata{Name: "OpenDocument spreadsheet", Type: "ODS"}
			}

			if bytes.Equal(data, []byte("application/vnd.oasis.opendocument.presentation")) {
				return &types.Metadata{Name: "OpenDocument presentation", Type: "ODP"}
			}
		}
	}

	if lowerName == "[content_types].xml" {
		switch {
		case bytes.Contains(b, []byte("word/")):
			return &types.Metadata{Name: "Microsoft Word document", Type: "DOCX"}
		case bytes.Contains(b, []byte("xl/")):
			return &types.Metadata{Name: "Microsoft Excel spreadsheet", Type: "XLSX"}
		case bytes.Contains(b, []byte("ppt/")):
			return &types.Metadata{Name: "Microsoft PowerPoint presentation", Type: "PPTX"}
		}
	}

	if lowerName == "androidmanifest.xml" || bytes.Contains(b, []byte("classes.dex")) {
		return &types.Metadata{Name: "Android package", Type: "APK"}
	}

	if lowerName == "meta-inf/manifest.mf" || bytes.Contains(b, []byte("META-INF/MANIFEST.MF")) {
		return &types.Metadata{Name: "Java archive", Type: "JAR"}
	}

	if strings.HasPrefix(lowerName, "payload/") || bytes.Contains(b, []byte("Payload/")) {
		return &types.Metadata{Name: "iOS application archive", Type: "IPA"}
	}

	return &types.Metadata{Name: "ZIP archive"}
}

func firstZipLocalName(b types.Buffer) (string, bool) {
	if !b.Has(0, []byte{'P', 'K', 3, 4}) {
		return "", false
	}

	nameLen, ok := b.U16LE(26)
	if !ok || nameLen == 0 {
		return "", false
	}

	_, ok = b.U16LE(28)
	if !ok {
		return "", false
	}

	nameStart := 30
	nameEnd := nameStart + int(nameLen)
	if nameEnd > b.Len() {
		return "", false
	}

	return string(b[nameStart:nameEnd]), true
}

func firstZipStoredData(b types.Buffer) []byte {
	if !b.Has(0, []byte{'P', 'K', 3, 4}) {
		return nil
	}

	compression, ok := b.U16LE(8)
	if !ok || compression != 0 {
		return nil
	}

	nameLen, ok := b.U16LE(26)
	if !ok {
		return nil
	}

	extraLen, ok := b.U16LE(28)
	if !ok {
		return nil
	}

	dataLen, ok := b.U32LE(18)
	if !ok || dataLen == 0 {
		return nil
	}

	dataStart := 30 + int(nameLen) + int(extraLen)
	dataEnd := dataStart + int(dataLen)
	if dataStart < 0 || dataEnd > b.Len() {
		return nil
	}

	return b[dataStart:dataEnd]
}
