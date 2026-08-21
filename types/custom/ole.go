package custom

import (
	"bytes"

	"github.com/coalaura/wtf/types"
)

var (
	oleWordDocument       = []byte{'W', 0, 'o', 0, 'r', 0, 'd', 0, 'D', 0, 'o', 0, 'c', 0, 'u', 0, 'm', 0, 'e', 0, 'n', 0, 't', 0, 0, 0}
	oleWorkbook           = []byte{'W', 0, 'o', 0, 'r', 0, 'k', 0, 'b', 0, 'o', 0, 'o', 0, 'k', 0, 0, 0}
	oleBook               = []byte{'B', 0, 'o', 0, 'o', 0, 'k', 0, 0, 0}
	olePowerPointDocument = []byte{'P', 0, 'o', 0, 'w', 0, 'e', 0, 'r', 0, 'P', 0, 'o', 0, 'i', 0, 'n', 0, 't', 0, ' ', 0, 'D', 0, 'o', 0, 'c', 0, 'u', 0, 'm', 0, 'e', 0, 'n', 0, 't', 0, 0, 0}
	oleMSI                = []byte{0x84, 0x10, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}
	oleMSP                = []byte{'M', 0, 'S', 0, 'P', 0}
	oleOutlookMessage     = []byte{'_', 0, '_', 0, 's', 0, 'u', 0, 'b', 0, 's', 0, 't', 0, 'g', 0, '1', 0, '.', 0, '0', 0, '_', 0}
	oleVisioDocument      = []byte{'V', 0, 'i', 0, 's', 0, 'i', 0, 'o', 0, 'D', 0, 'o', 0, 'c', 0, 'u', 0, 'm', 0, 'e', 0, 'n', 0, 't', 0}
	oleProject            = []byte{'M', 0, 'S', 0, 'P', 0, 'r', 0, 'o', 0, 'j', 0, 'e', 0, 'c', 0, 't', 0, '.', 0, 'P', 0, 'r', 0, 'o', 0, 'j', 0, 'e', 0, 'c', 0, 't', 0}
	olePublisher          = []byte{'P', 0, 'u', 0, 'b', 0, 'l', 0, 'i', 0, 's', 0, 'h', 0, 'e', 0, 'r', 0, 'D', 0, 'o', 0, 'c', 0, 'u', 0, 'm', 0, 'e', 0, 'n', 0, 't', 0}
)

func DetectOLE(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte{0xd0, 0xcf, 0x11, 0xe0, 0xa1, 0xb1, 0x1a, 0xe1}) {
		return nil
	}

	if containsOLE(b, oleWordDocument) || containsOLE(b, []byte("MSWordDoc")) || containsOLE(b, []byte("Word.Document.")) {
		return &types.Metadata{Kind: types.KindOLECompoundDocument, Type: types.TypeMicrosoftWordDocument}
	}

	if containsOLE(b, oleWorkbook) || containsOLE(b, oleBook) || containsOLE(b, []byte("Excel.Sheet.")) {
		return &types.Metadata{Kind: types.KindOLECompoundDocument, Type: types.TypeMicrosoftExcelWorkbook}
	}

	if containsOLE(b, olePowerPointDocument) || containsOLE(b, []byte("PowerPoint.Show.")) {
		return &types.Metadata{Kind: types.KindOLECompoundDocument, Type: types.TypeMicrosoftPowerPointPresentation}
	}

	if containsOLE(b, oleMSI) {
		return &types.Metadata{Kind: types.KindOLECompoundDocument, Type: types.TypeMicrosoftInstaller}
	}

	if containsOLE(b, oleMSP) {
		return &types.Metadata{Kind: types.KindOLECompoundDocument, Type: types.TypeMicrosoftInstallerPatch}
	}

	if containsOLE(b, oleOutlookMessage) {
		return &types.Metadata{Kind: types.KindOLECompoundDocument, Type: types.TypeMicrosoftOutlookMessage}
	}

	if containsOLE(b, oleVisioDocument) {
		return &types.Metadata{Kind: types.KindOLECompoundDocument, Type: types.TypeMicrosoftVisioDrawing}
	}

	if containsOLE(b, oleProject) {
		return &types.Metadata{Kind: types.KindOLECompoundDocument, Type: types.TypeMicrosoftProjectDocument}
	}

	if containsOLE(b, olePublisher) {
		return &types.Metadata{Kind: types.KindOLECompoundDocument, Type: types.TypeMicrosoftPublisherDocument}
	}

	return &types.Metadata{
		Kind: types.KindOLECompoundDocument,
	}
}

func containsOLE(b types.Buffer, magic []byte) bool {
	prefixEnd := min(b.Len(), maxScanSize)
	if bytes.Contains(b[:prefixEnd], magic) {
		return true
	}

	suffixStart := max(prefixEnd, b.Len()-maxScanSize)

	return bytes.Contains(b[suffixStart:], magic)
}
