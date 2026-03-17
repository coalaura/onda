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

	names := extractZipFilenames(b)
	if len(names) == 0 {
		return &types.Metadata{Kind: types.KindZIPArchive}
	}

	if strings.ToLower(names[0]) == "mimetype" {
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

			if bytes.Equal(data, []byte("image/openraster")) {
				return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeOpenRasterImageORA}
			}
		}
	}

	var (
		hasManifest      bool
		hasDex           bool
		hasAppxManifest  bool
		hasAppxBlockMap  bool
		hasAppxSignature bool
		hasSketchMeta    bool
		hasSketchUser    bool
		hasSketchDoc     bool
		hasWord          bool
		hasExcel         bool
		hasPowerPoint    bool
	)

	for _, name := range names {
		lower := strings.ToLower(name)

		if lower == "androidmanifest.xml" {
			hasManifest = true
		}

		if strings.HasSuffix(lower, ".dex") || lower == "classes.dex" {
			hasDex = true
		}

		if lower == "appxmanifest.xml" {
			hasAppxManifest = true
		}

		if lower == "appxblockmap.xml" {
			hasAppxBlockMap = true
		}

		if lower == "appxsignature.p7x" {
			hasAppxSignature = true
		}

		if lower == "document.json" {
			hasSketchDoc = true
		}

		if lower == "meta.json" {
			hasSketchMeta = true
		}

		if lower == "user.json" {
			hasSketchUser = true
		}

		if strings.HasPrefix(lower, "word/") {
			hasWord = true
		}

		if strings.HasPrefix(lower, "xl/") {
			hasExcel = true
		}

		if strings.HasPrefix(lower, "ppt/") {
			hasPowerPoint = true
		}

		if lower == "doc.kml" {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeKMZArchive}
		}

		if strings.HasSuffix(lower, ".dist-info/wheel") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypePythonWheelWHL}
		}

		if lower == "manifest.json" && bytes.Contains(b, []byte("xapk_version")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidPackageXAPK}
		}

		if lower == "install.rdf" || (lower == "manifest.json" && bytes.Contains(b, []byte("browser_specific_settings"))) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeFirefoxExtensionXPI}
		}

		if lower == "meta-inf/manifest.mf" {
			if bytes.Contains(b, []byte("WEB-INF/web.xml")) {
				return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeJavaWebArchiveWAR}
			}
			if bytes.Contains(b, []byte("META-INF/application.xml")) {
				return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeJavaEnterpriseArchiveEAR}
			}
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeJavaArchiveJAR}
		}

		if strings.HasPrefix(lower, "payload/") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeIOSApplicationArchiveIPA}
		}

		if lower == "bundleconfig.pb" {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidAppBundleAAB}
		}

		if lower == "toc.pb" {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidSplitAPKS}
		}

		if strings.Contains(lower, "apex_manifest") {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidSystemPackageAPEX}
		}

		if lower == "extension.vsixmanifest" {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeVisualStudioExtensionVSIX}
		}

		if strings.HasSuffix(lower, ".nuspec") && bytes.Contains(b, []byte("package/services/metadata/core-properties")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeNuGetPackageNUPKG}
		}
	}

	if hasWord {
		if bytes.Contains(b, []byte("application/vnd.ms-word.document.macroEnabled.main+xml")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordMacroEnabledDocumentDOCM}
		}

		if bytes.Contains(b, []byte("application/vnd.openxmlformats-officedocument.wordprocessingml.template.main+xml")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordTemplateDOTX}
		}

		if bytes.Contains(b, []byte("application/vnd.ms-word.template.macroEnabledTemplate.main+xml")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordMacroEnabledTemplateDOTM}
		}

		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordDocumentDOCX}
	}

	if hasExcel {
		if bytes.Contains(b, []byte("application/vnd.ms-excel.sheet.macroEnabled.main+xml")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelMacroEnabledWorkbookXLSM}
		}

		if bytes.Contains(b, []byte("application/vnd.openxmlformats-officedocument.spreadsheetml.template.main+xml")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelTemplateXLTX}
		}

		if bytes.Contains(b, []byte("application/vnd.ms-excel.template.macroEnabled.main+xml")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelMacroEnabledTemplateXLTM}
		}

		if bytes.Contains(b, []byte("application/vnd.ms-excel.addin.macroEnabled.main+xml")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelAddInXLAM}
		}

		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelSpreadsheetXLSX}
	}

	if hasPowerPoint {
		if bytes.Contains(b, []byte("application/vnd.ms-powerpoint.presentation.macroEnabled.main+xml")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointMacroEnabledPresentationPPTM}
		}

		if bytes.Contains(b, []byte("application/vnd.openxmlformats-officedocument.presentationml.template.main+xml")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointTemplatePOTX}
		}

		if bytes.Contains(b, []byte("application/vnd.ms-powerpoint.template.macroEnabled.main+xml")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointMacroEnabledTemplatePOTM}
		}

		if bytes.Contains(b, []byte("application/vnd.openxmlformats-officedocument.presentationml.slideshow.main+xml")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointSlideshowPPSX}
		}

		if bytes.Contains(b, []byte("application/vnd.ms-powerpoint.slideshow.macroEnabled.main+xml")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointMacroEnabledSlideshowPPSM}
		}

		if bytes.Contains(b, []byte("application/vnd.ms-powerpoint.addin.macroEnabled.main+xml")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointAddInPPAM}
		}

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

		if hasAppxBlockMap {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAPPXPackage}
		}

		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAPPXPackage}
	}

	if hasSketchDoc && (hasSketchMeta || hasSketchUser) {
		if bytes.Contains(b, []byte("com.bohemiancoding.sketch")) || bytes.Contains(b, []byte("com.sketch3")) {
			return &types.Metadata{Kind: types.KindSketchDocument, Type: types.TypeSketchDocument}
		}
	}

	return &types.Metadata{Kind: types.KindZIPArchive}
}

func extractZipFilenames(b types.Buffer) []string {
	var names []string

	limit := b.Len() - 30

	for i := 0; i <= limit; i++ {
		if b[i] == 'P' && b[i+1] == 'K' && b[i+2] == 3 && b[i+3] == 4 {
			nameLen, ok := b.U16LE(i + 26)
			if !ok || nameLen == 0 || i+30+int(nameLen) > b.Len() {
				continue
			}

			names = append(names, string(b[i+30:i+30+int(nameLen)]))

			i += 29 + int(nameLen)
		}
	}

	return names
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
