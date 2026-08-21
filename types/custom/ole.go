package custom

import (
	"bytes"
	"io"

	"github.com/coalaura/wtf/types"
)

const (
	maxDIFATSectors   = 16
	maxDirectoryBytes = maxScanSize
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

	meta := detectOLEReaderAt(bytes.NewReader(b), int64(b.Len()))
	if meta != nil {
		return meta
	}

	meta = detectOLESubtype(b)
	if meta != nil {
		return meta
	}

	return &types.Metadata{
		Kind: types.KindOLECompoundDocument,
	}
}

func detectOLESubtype(b types.Buffer) *types.Metadata {
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

	return nil
}

func containsOLE(b types.Buffer, magic []byte) bool {
	prefixEnd := min(b.Len(), maxScanSize)
	if bytes.Contains(b[:prefixEnd], magic) {
		return true
	}

	suffixStart := max(prefixEnd, b.Len()-maxScanSize)

	return bytes.Contains(b[suffixStart:], magic)
}

func detectOLEReaderAt(r io.ReaderAt, size int64) *types.Metadata {
	header := make([]byte, 512)
	if size < int64(len(header)) || !readAtExactly(r, header, 0) {
		return nil
	}

	b := types.Buffer(header)
	if !b.Has(0, []byte{0xd0, 0xcf, 0x11, 0xe0, 0xa1, 0xb1, 0x1a, 0xe1}) {
		return nil
	}

	majorVersion, _ := b.U16LE(26)
	byteOrder, _ := b.U16LE(28)
	sectorShift, _ := b.U16LE(30)

	if byteOrder != 0xfffe || majorVersion == 3 && sectorShift != 9 || majorVersion == 4 && sectorShift != 12 || majorVersion != 3 && majorVersion != 4 {
		return nil
	}

	sectorSize := int64(1) << sectorShift
	if size < sectorSize {
		return nil
	}

	totalSectors := uint32((size - sectorSize) / sectorSize)

	numFATSectors, _ := b.U32LE(44)
	firstDirectorySector, _ := b.U32LE(48)
	firstDIFATSector, _ := b.U32LE(68)
	numDIFATSectors, _ := b.U32LE(72)

	if numFATSectors == 0 || numFATSectors > maxScanSize/4 || numDIFATSectors > maxDIFATSectors || firstDirectorySector >= totalSectors {
		return nil
	}

	fatSectors := make([]uint32, 0, numFATSectors)

	for i := 0; i < 109 && len(fatSectors) < int(numFATSectors); i++ {
		sector, _ := b.U32LE(76 + i*4)
		if sector == 0xffffffff {
			continue
		}

		if sector >= totalSectors {
			return nil
		}

		fatSectors = append(fatSectors, sector)
	}

	difatSector := firstDIFATSector
	difatSeen := make(map[uint32]bool)

	for i := uint32(0); i < numDIFATSectors && len(fatSectors) < int(numFATSectors); i++ {
		if difatSector >= totalSectors || difatSeen[difatSector] {
			return nil
		}

		difatSeen[difatSector] = true

		sectorData := make([]byte, int(sectorSize))

		if !readAtExactly(r, sectorData, (int64(difatSector)+1)*sectorSize) {
			return nil
		}

		sectorBuffer := types.Buffer(sectorData)

		for offset := 0; offset+4 < len(sectorData) && len(fatSectors) < int(numFATSectors); offset += 4 {
			sector, _ := sectorBuffer.U32LE(offset)
			if sector == 0xffffffff {
				continue
			}

			if sector >= totalSectors {
				return nil
			}

			fatSectors = append(fatSectors, sector)
		}

		difatSector, _ = sectorBuffer.U32LE(len(sectorData) - 4)
	}

	if len(fatSectors) != int(numFATSectors) {
		return nil
	}

	fatCache := make(map[uint32][]byte)

	nextSector := func(sector uint32) (uint32, bool) {
		entriesPerSector := uint32(sectorSize / 4)

		fatIndex := sector / entriesPerSector
		if fatIndex >= uint32(len(fatSectors)) {
			return 0, false
		}

		fatSector := fatSectors[fatIndex]

		data := fatCache[fatSector]
		if data == nil {
			data = make([]byte, int(sectorSize))
			if !readAtExactly(r, data, (int64(fatSector)+1)*sectorSize) {
				return 0, false
			}

			fatCache[fatSector] = data
		}

		next, ok := types.Buffer(data).U32LE(int(sector%entriesPerSector) * 4)
		return next, ok
	}

	directory := make([]byte, 0, min(maxDirectoryBytes, int(sectorSize)*4))
	directorySeen := make(map[uint32]bool)
	sector := firstDirectorySector

	for sector != 0xfffffffe {
		if sector >= totalSectors || directorySeen[sector] || len(directory)+int(sectorSize) > maxDirectoryBytes {
			return nil
		}

		directorySeen[sector] = true

		start := len(directory)

		directory = append(directory, make([]byte, int(sectorSize))...)

		if !readAtExactly(r, directory[start:], (int64(sector)+1)*sectorSize) {
			return nil
		}

		var ok bool

		sector, ok = nextSector(sector)
		if !ok || sector >= 0xfffffff8 && sector != 0xfffffffe {
			return nil
		}
	}

	for offset := 0; offset+128 <= len(directory); offset += 128 {
		entry := types.Buffer(directory[offset : offset+128])

		nameLength, _ := entry.U16LE(64)
		if nameLength == 0 {
			continue
		}

		if nameLength < 2 || nameLength > 64 || nameLength&1 != 0 || entry[nameLength-2] != 0 || entry[nameLength-1] != 0 {
			return nil
		}
	}

	if meta := detectOLESubtype(types.Buffer(directory)); meta != nil {
		return meta
	}

	return &types.Metadata{Kind: types.KindOLECompoundDocument}
}
