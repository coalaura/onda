package custom

import (
	"bytes"

	"github.com/coalaura/onda/types"
)

var (
	oleWordDocument       = []byte{'W', 0, 'o', 0, 'r', 0, 'd', 0, 'D', 0, 'o', 0, 'c', 0, 'u', 0, 'm', 0, 'e', 0, 'n', 0, 't', 0}
	oleWorkbook           = []byte{'W', 0, 'o', 0, 'r', 0, 'k', 0, 'b', 0, 'o', 0, 'o', 0, 'k', 0}
	oleBook               = []byte{'B', 0, 'o', 0, 'o', 0, 'k', 0}
	olePowerPointDocument = []byte{'P', 0, 'o', 0, 'w', 0, 'e', 0, 'r', 0, 'P', 0, 'o', 0, 'i', 0, 'n', 0, 't', 0, ' ', 0, 'D', 0, 'o', 0, 'c', 0, 'u', 0, 'm', 0, 'e', 0, 'n', 0, 't', 0}
	oleMSI                = []byte{0x84, 0x10, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}
	oleOutlookMessage     = []byte{'_', 0, '_', 0, 's', 0, 'u', 0, 'b', 0, 's', 0, 't', 0, 'g', 0, '1', 0, '.', 0, '0', 0, '_', 0}
	oleVisioDocument      = []byte{'V', 0, 'i', 0, 's', 0, 'i', 0, 'o', 0, 'D', 0, 'o', 0, 'c', 0, 'u', 0, 'm', 0, 'e', 0, 'n', 0, 't', 0}
)

func DetectOLE(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte{0xd0, 0xcf, 0x11, 0xe0, 0xa1, 0xb1, 0x1a, 0xe1}) {
		return nil
	}

	limit := min(b.Len(), 4096)
	data := b[:limit]

	if bytes.Contains(data, oleWordDocument) {
		return &types.Metadata{Kind: types.KindOLECompoundDocument, Type: types.TypeMicrosoftWordDocumentDOC}
	}

	if bytes.Contains(data, oleWorkbook) || bytes.Contains(data, oleBook) {
		return &types.Metadata{Kind: types.KindOLECompoundDocument, Type: types.TypeMicrosoftExcelWorkbookXLS}
	}

	if bytes.Contains(data, olePowerPointDocument) {
		return &types.Metadata{Kind: types.KindOLECompoundDocument, Type: types.TypeMicrosoftPowerPointPresentationPPT}
	}

	if bytes.Contains(data, oleMSI) {
		return &types.Metadata{Kind: types.KindOLECompoundDocument, Type: types.TypeMicrosoftInstallerMSI}
	}

	if bytes.Contains(data, oleOutlookMessage) {
		return &types.Metadata{Kind: types.KindOLECompoundDocument, Type: types.TypeMicrosoftOutlookMessageMSG}
	}

	if bytes.Contains(data, oleVisioDocument) {
		return &types.Metadata{Kind: types.KindOLECompoundDocument, Type: types.TypeMicrosoftVisioDrawingVSD}
	}

	return &types.Metadata{
		Kind: types.KindOLECompoundDocument,
	}
}
