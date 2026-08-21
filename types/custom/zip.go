package custom

import (
	"bytes"
	"io"

	"github.com/coalaura/wtf/types"
)

const (
	maxEOCDSize       = 22 + 65535
	maxCentralSize    = maxScanSize * 8
	maxSyntheticSize  = maxScanSize * 16
	centralHeaderSize = 46
	localHeaderSize   = 30
)

func DetectZIPContainer(b types.Buffer) *types.Metadata {
	if b.Len() < 4 || string(b[:4]) != "PK\x03\x04" {
		return nil
	}

	meta := detectZIPEntries(b[:min(b.Len(), maxScanSize)])

	refined := detectZIPReaderAt(bytes.NewReader(b), int64(b.Len()))
	if refined != nil {
		return preferZIPMetadata(meta, refined)
	}

	return meta
}

func preferZIPMetadata(current, refined *types.Metadata) *types.Metadata {
	switch refined.Type {
	case types.TypeMicrosoftWordDocument:
		switch current.Type {
		case types.TypeMicrosoftWordMacroEnabledDocument, types.TypeMicrosoftWordTemplate, types.TypeMicrosoftWordMacroEnabledTemplate:
			return current
		}
	case types.TypeMicrosoftExcelWorkbook:
		switch current.Type {
		case types.TypeMicrosoftExcelMacroEnabledWorkbook, types.TypeMicrosoftExcelTemplate, types.TypeMicrosoftExcelMacroEnabledTemplate, types.TypeMicrosoftExcelAddIn:
			return current
		}
	case types.TypeMicrosoftPowerPointPresentation:
		switch current.Type {
		case types.TypeMicrosoftPowerPointMacroEnabledPresentation, types.TypeMicrosoftPowerPointTemplate, types.TypeMicrosoftPowerPointMacroEnabledTemplate, types.TypeMicrosoftPowerPointSlideshow, types.TypeMicrosoftPowerPointMacroEnabledSlideshow, types.TypeMicrosoftPowerPointAddIn:
			return current
		}
	}

	return refined
}

func detectZIPEntries(b types.Buffer) *types.Metadata {
	var (
		hasManifest         bool
		hasDex              bool
		hasAppxManifest     bool
		hasAppxSignature    bool
		hasSketchMeta       bool
		hasSketchUser       bool
		hasSketchDoc        bool
		hasWord             bool
		hasExcel            bool
		hasPowerPoint       bool
		hasWordVBA          bool
		hasExcelVBA         bool
		hasPowerPointVBA    bool
		hasManifestMF       bool
		hasWebXML           bool
		hasAppXML           bool
		hasMinecraftMeta    bool
		hasFabricMod        bool
		hasForgeMod         bool
		hasLottieManifest   bool
		hasLottieAnimations bool
		hasPyTorchData      bool
		hasPyTorchVersion   bool
		hasComicInfo        bool
		hasProcreateDoc     bool
		hasCeltxScript      bool
		hasStudioOneSong    bool
		hasBitwigProject    bool
		hasPSVitaEboot      bool
		hasOsuBeatmap       bool
		hasOsuSkin          bool
		firstFile           = true
	)

	limit := b.Len() - 30
	zipMagic := []byte{'P', 'K', 3, 4}

	for i := 0; i <= limit; {
		idx := bytes.Index(b[i:], zipMagic)
		if idx == -1 {
			break
		}

		i += idx

		nameLen, ok := b.U16LE(i + 26)
		extraLen, extraOK := b.U16LE(i + 28)
		compressedSize, sizeOK := b.U32LE(i + 18)
		flags, flagsOK := b.U16LE(i + 6)

		dataStart := i + 30 + int(nameLen) + int(extraLen)

		if !ok || !extraOK || !sizeOK || !flagsOK || nameLen == 0 || dataStart < i || dataStart > b.Len() {
			i += 4
			continue
		}

		name := b[i+30 : i+30+int(nameLen)]

		if firstFile && string(name) == "mimetype" {
			compression, _ := b.U16LE(i + 8)

			dataEnd := dataStart + int(compressedSize)

			if compression == 0 && dataEnd >= dataStart && dataEnd <= b.Len() {
				switch string(b[dataStart:dataEnd]) {
				case "application/epub+zip":
					return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeEPUBDocument}
				case "application/vnd.oasis.opendocument.text":
					return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOpenDocumentText}
				case "application/vnd.oasis.opendocument.spreadsheet":
					return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOpenDocumentSpreadsheet}
				case "application/vnd.oasis.opendocument.presentation":
					return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOpenDocumentPresentation}
				case "application/vnd.oasis.opendocument.graphics":
					return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOpenDocumentGraphics}
				case "application/vnd.oasis.opendocument.database":
					return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOpenDocumentDatabase}
				case "application/vnd.oasis.opendocument.formula":
					return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOpenDocumentFormula}
				case "application/vnd.oasis.opendocument.chart":
					return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOpenDocumentChart}
				case "application/vnd.oasis.opendocument.image":
					return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOpenDocumentImage}
				case "application/x-krita":
					return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeKritaDocument}
				case "image/openraster":
					return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOpenRasterImage}
				}
			}
		}

		firstFile = false

		if matchASCII(name, "androidmanifest.xml") {
			hasManifest = true
		} else if hasSuffixASCII(name, ".dex") || matchASCII(name, "classes.dex") {
			hasDex = true
		} else if matchASCII(name, "appxmanifest.xml") {
			hasAppxManifest = true
		} else if matchASCII(name, "appxsignature.p7x") {
			hasAppxSignature = true
		} else if matchASCII(name, "document.json") {
			hasSketchDoc = true
		} else if matchASCII(name, "meta.json") {
			hasSketchMeta = true
		} else if matchASCII(name, "user.json") {
			hasSketchUser = true
		} else if matchASCII(name, "modules") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeNone}
		} else if matchASCII(name, "meta-inf/manifest.mf") {
			hasManifestMF = true
		} else if matchASCII(name, "web-inf/web.xml") {
			hasWebXML = true
		} else if matchASCII(name, "meta-inf/application.xml") {
			hasAppXML = true
		} else if matchASCII(name, "pack.mcmeta") {
			hasMinecraftMeta = true
		} else if matchASCII(name, "fabric.mod.json") {
			hasFabricMod = true
		} else if matchASCII(name, "mcmod.info") || matchASCII(name, "meta-inf/mods.toml") {
			hasForgeMod = true
		} else if matchASCII(name, "model.weights.h5") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeKerasModel}
		} else if hasSuffixASCII(name, "/data.pkl") || matchASCII(name, "data.pkl") {
			hasPyTorchData = true
		} else if hasSuffixASCII(name, "/version") || matchASCII(name, "version") {
			hasPyTorchVersion = true
		} else if matchASCII(name, "main.lua") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeLOVEGame}
		} else if matchASCII(name, "doc.kml") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeKMZArchive}
		} else if hasSuffixASCII(name, ".dist-info/wheel") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypePythonWheel}
		} else if matchASCII(name, "manifest.json") {
			hasLottieManifest = true

			if bytes.Contains(b, []byte("xapk_version")) {
				return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidPackageX}
			}

			if bytes.Contains(b, []byte("browser_specific_settings")) {
				return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeFirefoxExtension}
			}
		} else if matchASCII(name, "install.rdf") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeFirefoxExtension}
		} else if hasPrefixASCII(name, "animations/") {
			hasLottieAnimations = true
		} else if hasPrefixASCII(name, "payload/") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeIOSApplicationArchive}
		} else if matchASCII(name, "bundleconfig.pb") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidAppBundle}
		} else if matchASCII(name, "toc.pb") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidSplitAPKSet}
		} else if bytes.Contains(name, []byte("apex_manifest")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidSystemPackage}
		} else if matchASCII(name, "extension.vsixmanifest") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeVisualStudioExtension}
		} else if matchASCII(name, "3d/3dmodel.model") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.Type3MFDocument}
		} else if hasSuffixASCII(name, ".nuspec") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeNuGetPackage}
		} else if matchASCII(name, "comicinfo.xml") {
			hasComicInfo = true
		} else if matchASCII(name, "document.archive") {
			hasProcreateDoc = true
		} else if matchASCII(name, "script.html") || matchASCII(name, "project.rdf") {
			hasCeltxScript = true
		} else if matchASCII(name, "song.xml") || matchASCII(name, "project.xml") {
			hasStudioOneSong = true
		} else if matchASCII(name, "project.json") {
			hasBitwigProject = true
		} else if matchASCII(name, "eboot.bin") || matchASCII(name, "sce_sys/param.sfo") {
			hasPSVitaEboot = true
		} else if hasSuffixASCII(name, ".osu") {
			hasOsuBeatmap = true
		} else if matchASCII(name, "skin.ini") {
			hasOsuSkin = true
		} else if matchASCII(name, "metadata/buildversionhistory.plist") || hasPrefixASCII(name, "index/document.iwa") {
			return &types.Metadata{Kind: types.KindAppleiWorkDocument}
		}

		if hasPrefixASCII(name, "word/") {
			hasWord = true
		} else if hasPrefixASCII(name, "xl/") {
			hasExcel = true
		} else if hasPrefixASCII(name, "ppt/") {
			hasPowerPoint = true
		}

		if matchASCII(name, "word/vbaproject.bin") {
			hasWordVBA = true
		} else if matchASCII(name, "xl/vbaproject.bin") {
			hasExcelVBA = true
		} else if matchASCII(name, "ppt/vbaproject.bin") {
			hasPowerPointVBA = true
		}

		if flags&0x08 != 0 || uint64(dataStart)+uint64(compressedSize) > uint64(b.Len()) {
			break
		}

		i = dataStart + int(compressedSize)
	}

	limitSearch := min(b.Len(), 32768)
	searchArea := b[:limitSearch]

	// Macro-enabled detection via vbaProject.bin
	if hasWordVBA {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordMacroEnabledDocument}
	}

	if hasExcelVBA {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelMacroEnabledWorkbook}
	}

	if hasPowerPointVBA {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointMacroEnabledPresentation}
	}

	// Word
	if bytes.Contains(searchArea, []byte("application/vnd.ms-word.document.macroEnabled.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordMacroEnabledDocument}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.openxmlformats-officedocument.wordprocessingml.template.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordTemplate}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.ms-word.template.macroEnabledTemplate.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordMacroEnabledTemplate}
	}

	if hasWord {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordDocument}
	}

	// Excel
	if bytes.Contains(searchArea, []byte("application/vnd.ms-excel.sheet.macroEnabled.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelMacroEnabledWorkbook}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.openxmlformats-officedocument.spreadsheetml.template.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelTemplate}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.ms-excel.template.macroEnabled.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelMacroEnabledTemplate}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.ms-excel.addin.macroEnabled.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelAddIn}
	}

	if hasExcel {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelWorkbook}
	}

	// PowerPoint
	if bytes.Contains(searchArea, []byte("application/vnd.ms-powerpoint.presentation.macroEnabled.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointMacroEnabledPresentation}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.openxmlformats-officedocument.presentationml.template.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointTemplate}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.ms-powerpoint.template.macroEnabled.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointMacroEnabledTemplate}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.openxmlformats-officedocument.presentationml.slideshow.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointSlideshow}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.ms-powerpoint.slideshow.macroEnabled.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointMacroEnabledSlideshow}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.ms-powerpoint.addin.macroEnabled.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointAddIn}
	}

	if hasPowerPoint {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointPresentation}
	}

	// XPS
	if bytes.Contains(searchArea, []byte("application/vnd.ms-xpsdocument.document+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeXMLPaperSpecification}
	}

	if hasManifest {
		if hasDex {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidPackage}
		}

		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidArchive}
	}

	if hasAppxManifest {
		if hasAppxSignature {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMSIXPackage}
		}

		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAPPXPackage}
	}

	if hasSketchDoc && (hasSketchMeta || hasSketchUser) {
		if bytes.Contains(searchArea, []byte("com.bohemiancoding.sketch")) || bytes.Contains(searchArea, []byte("com.sketch3")) {
			return &types.Metadata{Kind: types.KindSketchDocument}
		}
	}

	if hasManifestMF {
		if hasWebXML {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeJavaWebArchive}
		}

		if hasAppXML {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeJavaEnterpriseArchive}
		}

		if hasFabricMod {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeFabricMod}
		}

		if hasForgeMod {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeForgeMod}
		}

		if hasMinecraftMeta {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMinecraftResourcePack}
		}

		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeJavaArchive}
	}

	if hasLottieManifest && hasLottieAnimations {
		return &types.Metadata{Kind: types.KindLottieAnimation}
	}

	if hasPyTorchData && hasPyTorchVersion {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypePyTorchModel}
	}

	if hasComicInfo {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeComicBook}
	}

	if hasProcreateDoc {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeProcreate}
	}

	if hasCeltxScript {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeCeltx}
	}

	if hasStudioOneSong {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeStudioOne}
	}

	if hasBitwigProject {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeBitwig}
	}

	if hasPSVitaEboot {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypePSVita}
	}

	if hasOsuBeatmap {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOsuBeatmap}
	}

	if hasOsuSkin {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOsuSkin}
	}

	return &types.Metadata{Kind: types.KindZIPArchive}
}

func detectZIPReaderAt(r io.ReaderAt, size int64) *types.Metadata {
	if size < 22 {
		return nil
	}

	tailSize := min(size, int64(maxEOCDSize))
	tail := make([]byte, int(tailSize))

	if !readAtExactly(r, tail, size-tailSize) {
		return nil
	}

	eocd := -1
	searchEnd := len(tail)

	for searchEnd >= 4 {
		candidate := bytes.LastIndex(tail[:searchEnd], []byte{'P', 'K', 5, 6})
		if candidate < 0 {
			break
		}

		if candidate+22 <= len(tail) {
			commentLength, _ := types.Buffer(tail).U16LE(candidate + 20)
			if candidate+22+int(commentLength) == len(tail) {
				eocd = candidate

				break
			}
		}

		searchEnd = candidate
	}

	if eocd < 0 {
		return nil
	}

	tailBuffer := types.Buffer(tail)
	disk, _ := tailBuffer.U16LE(eocd + 4)
	centralDisk, _ := tailBuffer.U16LE(eocd + 6)
	entriesOnDisk, _ := tailBuffer.U16LE(eocd + 8)
	entryCount, _ := tailBuffer.U16LE(eocd + 10)
	centralSize, _ := tailBuffer.U32LE(eocd + 12)
	centralOffset, _ := tailBuffer.U32LE(eocd + 16)
	eocdOffset := size - tailSize + int64(eocd)

	if disk != 0 || centralDisk != 0 || entriesOnDisk != entryCount || entryCount == 0 || entryCount == 0xffff || centralSize == 0xffffffff || centralOffset == 0xffffffff || centralSize > maxCentralSize || int64(centralOffset)+int64(centralSize) != eocdOffset {
		return nil
	}

	central := make([]byte, int(centralSize))
	if !readAtExactly(r, central, int64(centralOffset)) {
		return nil
	}

	entryData := make([]byte, 0, min(maxSyntheticSize, int(centralSize)))

	var (
		offset           int
		firstLocalOffset uint32
		firstName        []byte
	)

	for i := 0; i < int(entryCount); i++ {
		if offset+centralHeaderSize > len(central) || !bytes.Equal(central[offset:offset+4], []byte{'P', 'K', 1, 2}) {
			return nil
		}

		record := types.Buffer(central[offset:])

		nameLength, _ := record.U16LE(28)
		extraLength, _ := record.U16LE(30)
		commentLength, _ := record.U16LE(32)
		localOffset, _ := record.U32LE(42)

		recordSize := centralHeaderSize + int(nameLength) + int(extraLength) + int(commentLength)

		if nameLength == 0 || recordSize > len(central)-offset || len(entryData)+localHeaderSize+int(nameLength) > maxSyntheticSize {
			return nil
		}

		name := central[offset+centralHeaderSize : offset+centralHeaderSize+int(nameLength)]

		if i == 0 {
			firstLocalOffset = localOffset
			firstName = name
		}

		start := len(entryData)

		entryData = append(entryData, make([]byte, localHeaderSize)...)

		entryData[start] = 'P'
		entryData[start+1] = 'K'
		entryData[start+2] = 3
		entryData[start+3] = 4
		entryData[start+26] = byte(nameLength)
		entryData[start+27] = byte(nameLength >> 8)

		entryData = append(entryData, name...)

		offset += recordSize
	}

	if offset != len(central) || int64(firstLocalOffset) >= int64(centralOffset) {
		return nil
	}

	local := make([]byte, localHeaderSize+len(firstName))

	if !readAtExactly(r, local, int64(firstLocalOffset)) || !bytes.Equal(local[:4], []byte{'P', 'K', 3, 4}) {
		return nil
	}

	localBuffer := types.Buffer(local)

	localNameLength, _ := localBuffer.U16LE(26)

	if int(localNameLength) != len(firstName) || !bytes.Equal(local[localHeaderSize:localHeaderSize+len(firstName)], firstName) {
		return nil
	}

	meta := detectZIPEntries(types.Buffer(entryData))
	if meta.Kind == types.KindZIPArchive && meta.Type == types.TypeNone {
		return nil
	}

	return meta
}

func readAtExactly(r io.ReaderAt, data []byte, offset int64) bool {
	n, err := r.ReadAt(data, offset)
	return n == len(data) && (err == nil || err == io.EOF)
}

// matchASCII compares a byte slice to a lowercase string without allocating.
func matchASCII(b []byte, lower string) bool {
	if len(b) != len(lower) {
		return false
	}

	for i := range b {
		c := b[i]
		if c >= 'A' && c <= 'Z' {
			c += 32 // convert to lowercase
		}

		if c != lower[i] {
			return false
		}
	}

	return true
}

func hasSuffixASCII(b []byte, lowerSuffix string) bool {
	if len(b) < len(lowerSuffix) {
		return false
	}

	return matchASCII(b[len(b)-len(lowerSuffix):], lowerSuffix)
}

func hasPrefixASCII(b []byte, lowerPrefix string) bool {
	if len(b) < len(lowerPrefix) {
		return false
	}

	return matchASCII(b[:len(lowerPrefix)], lowerPrefix)
}
