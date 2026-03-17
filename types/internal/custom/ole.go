package custom

import "github.com/coalaura/onda/types"

var (
	oleWordDocument       = utf16LE("WordDocument")
	oleWorkbook           = utf16LE("Workbook")
	oleBook               = utf16LE("Book")
	olePowerPointDocument = utf16LE("PowerPoint Document")
)

func DetectOLE(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte{0xd0, 0xcf, 0x11, 0xe0, 0xa1, 0xb1, 0x1a, 0xe1}) {
		return nil
	}

	if hasBytes(b, oleWordDocument) {
		return &types.Metadata{Kind: types.KindOLECompoundDocument, Type: types.TypeMicrosoftWordDocumentDOC}
	}

	if hasBytes(b, oleWorkbook) || hasBytes(b, oleBook) {
		return &types.Metadata{Kind: types.KindOLECompoundDocument, Type: types.TypeMicrosoftExcelWorkbookXLS}
	}

	if hasBytes(b, olePowerPointDocument) {
		return &types.Metadata{Kind: types.KindOLECompoundDocument, Type: types.TypeMicrosoftPowerPointPresentationPPT}
	}

	return &types.Metadata{
		Kind: types.KindOLECompoundDocument,
	}
}

func hasBytes(b types.Buffer, needle []byte) bool {
	if len(needle) == 0 || len(needle) > b.Len() {
		return false
	}

	limit := min(b.Len(), 4096) - len(needle)

	for i := 0; i <= limit; i += 8 {
		if b.Has(i, needle) {
			return true
		}
	}

	return false
}

func utf16LE(s string) []byte {
	if len(s) == 0 {
		return nil
	}

	b := make([]byte, len(s)*2)

	for i := 0; i < len(s); i++ {
		b[i*2] = s[i]
		b[i*2+1] = 0
	}

	return b
}
