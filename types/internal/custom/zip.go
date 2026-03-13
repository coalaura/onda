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
		return &types.Metadata{Kind: types.KindZIPArchive}
	}

	lowerName := strings.ToLower(fileName)

	if lowerName == "mimetype" {
		if data := firstZipStoredData(b); len(data) > 0 {
			if bytes.Equal(data, []byte("application/epub+zip")) {
				return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeEPUBDocument}
			}

			if bytes.Equal(data, []byte("application/vnd.oasis.opendocument.text")) {
				return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOpenDocumentTextODT}
			}

			if bytes.Equal(data, []byte("application/vnd.oasis.opendocument.spreadsheet")) {
				return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOpenDocumentSpreadsheetODS}
			}

			if bytes.Equal(data, []byte("application/vnd.oasis.opendocument.presentation")) {
				return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOpenDocumentPresentationODP}
			}
		}
	}

	if lowerName == "[content_types].xml" {
		switch {
		case bytes.Contains(b, []byte("word/")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordDocumentDOCX}
		case bytes.Contains(b, []byte("xl/")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelSpreadsheetXLSX}
		case bytes.Contains(b, []byte("ppt/")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointPresentationPPTX}
		}
	}

	if lowerName == "androidmanifest.xml" || bytes.Contains(b, []byte("classes.dex")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidPackageAPK}
	}

	if lowerName == "meta-inf/manifest.mf" || bytes.Contains(b, []byte("META-INF/MANIFEST.MF")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeJavaArchiveJAR}
	}

	if strings.HasPrefix(lowerName, "payload/") || bytes.Contains(b, []byte("Payload/")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeIOSApplicationArchiveIPA}
	}

	return &types.Metadata{Kind: types.KindZIPArchive}
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
