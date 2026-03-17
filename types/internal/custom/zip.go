package custom

import (
	"bytes"

	"github.com/coalaura/onda/types"
)

func DetectZIPContainer(b types.Buffer) *types.Metadata {
	if b.Len() < 4 || string(b[:4]) != "PK\x03\x04" {
		return nil
	}

	var (
		hasManifest      bool
		hasDex           bool
		hasAppxManifest  bool
		hasAppxSignature bool
		hasSketchMeta    bool
		hasSketchUser    bool
		hasSketchDoc     bool
		hasWord          bool
		hasExcel         bool
		hasPowerPoint    bool
		firstFile        = true
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
		if !ok || nameLen == 0 || i+30+int(nameLen) > b.Len() {
			i += 4
			continue
		}

		name := b[i+30 : i+30+int(nameLen)]

		// Exact match using string cast (zero allocation)
		if firstFile && string(name) == "mimetype" {
			compression, _ := b.U16LE(i + 8)
			extraLen, _ := b.U16LE(i + 28)
			dataLen, _ := b.U32LE(i + 18)
			dataStart := i + 30 + int(nameLen) + int(extraLen)
			dataEnd := dataStart + int(dataLen)

			if compression == 0 && dataStart >= 0 && int(dataLen) >= 0 && dataEnd >= dataStart && dataEnd <= b.Len() {
				switch string(b[dataStart:dataEnd]) {
				case "application/epub+zip":
					return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeEPUBDocument}
				case "application/vnd.oasis.opendocument.text":
					return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOpenDocumentTextODT}
				case "application/vnd.oasis.opendocument.spreadsheet":
					return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOpenDocumentSpreadsheetODS}
				case "application/vnd.oasis.opendocument.presentation":
					return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOpenDocumentPresentationODP}
				case "image/openraster":
					return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOpenRasterImageORA}
				}
			}
		}

		firstFile = false

		// Fast, zero-allocation ASCII case-insensitive matching
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
		} else if matchASCII(name, "doc.kml") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeKMZArchive}
		} else if hasSuffixASCII(name, ".dist-info/wheel") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypePythonWheelWHL}
		} else if matchASCII(name, "manifest.json") && bytes.Contains(b, []byte("xapk_version")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidPackageXAPK}
		} else if matchASCII(name, "install.rdf") || (matchASCII(name, "manifest.json") && bytes.Contains(b, []byte("browser_specific_settings"))) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeFirefoxExtensionXPI}
		} else if matchASCII(name, "meta-inf/manifest.mf") {
			if bytes.Contains(b, []byte("WEB-INF/web.xml")) {
				return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeJavaWebArchiveWAR}
			}

			if bytes.Contains(b, []byte("META-INF/application.xml")) {
				return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeJavaEnterpriseArchiveEAR}
			}

			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeJavaArchiveJAR}
		} else if hasPrefixASCII(name, "payload/") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeIOSApplicationArchiveIPA}
		} else if matchASCII(name, "bundleconfig.pb") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidAppBundleAAB}
		} else if matchASCII(name, "toc.pb") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidSplitAPKS}
		} else if bytes.Contains(name, []byte("apex_manifest")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidSystemPackageAPEX}
		} else if matchASCII(name, "extension.vsixmanifest") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeVisualStudioExtensionVSIX}
		} else if matchASCII(name, "3d/3dmodel.model") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.Type3MFDocument}
		} else if hasSuffixASCII(name, ".nuspec") && bytes.Contains(b, []byte("package/services/metadata/core-properties")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeNuGetPackageNUPKG}
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

		i += 30 + int(nameLen)
	}

	limitSearch := min(b.Len(), 32768)
	searchArea := b[:limitSearch]

	// Word
	if bytes.Contains(searchArea, []byte("application/vnd.ms-word.document.macroEnabled.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordMacroEnabledDocumentDOCM}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.openxmlformats-officedocument.wordprocessingml.template.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordTemplateDOTX}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.ms-word.template.macroEnabledTemplate.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordMacroEnabledTemplateDOTM}
	}

	if hasWord {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordDocumentDOCX}
	}

	// Excel
	if bytes.Contains(searchArea, []byte("application/vnd.ms-excel.sheet.macroEnabled.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelMacroEnabledWorkbookXLSM}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.openxmlformats-officedocument.spreadsheetml.template.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelTemplateXLTX}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.ms-excel.template.macroEnabled.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelMacroEnabledTemplateXLTM}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.ms-excel.addin.macroEnabled.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelAddInXLAM}
	}

	if hasExcel {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelSpreadsheetXLSX}
	}

	// PowerPoint
	if bytes.Contains(searchArea, []byte("application/vnd.ms-powerpoint.presentation.macroEnabled.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointMacroEnabledPresentationPPTM}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.openxmlformats-officedocument.presentationml.template.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointTemplatePOTX}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.ms-powerpoint.template.macroEnabled.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointMacroEnabledTemplatePOTM}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.openxmlformats-officedocument.presentationml.slideshow.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointSlideshowPPSX}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.ms-powerpoint.slideshow.macroEnabled.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointMacroEnabledSlideshowPPSM}
	}

	if bytes.Contains(searchArea, []byte("application/vnd.ms-powerpoint.addin.macroEnabled.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointAddInPPAM}
	}

	if hasPowerPoint {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointPresentationPPTX}
	}

	if hasManifest {
		if hasDex {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidPackageAPK}
		}

		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidArchiveAAR}
	}

	if hasAppxManifest {
		if hasAppxSignature {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMSIXPackage}
		}

		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAPPXPackage}
	}

	if hasSketchDoc && (hasSketchMeta || hasSketchUser) {
		if bytes.Contains(searchArea, []byte("com.bohemiancoding.sketch")) || bytes.Contains(searchArea, []byte("com.sketch3")) {
			return &types.Metadata{Kind: types.KindSketchDocument, Type: types.TypeSketchDocument}
		}
	}

	return &types.Metadata{Kind: types.KindZIPArchive}
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
