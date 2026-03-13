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
		case bytes.Contains(b, []byte("application/vnd.ms-word.document.macroEnabled.main+xml")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordMacroEnabledDocumentDOCM}
		case bytes.Contains(b, []byte("application/vnd.openxmlformats-officedocument.wordprocessingml.template.main+xml")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordTemplateDOTX}
		case bytes.Contains(b, []byte("application/vnd.ms-word.template.macroEnabledTemplate.main+xml")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordMacroEnabledTemplateDOTM}
		case bytes.Contains(b, []byte("application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordDocumentDOCX}
		case bytes.Contains(b, []byte("application/vnd.ms-excel.sheet.macroEnabled.main+xml")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelMacroEnabledWorkbookXLSM}
		case bytes.Contains(b, []byte("application/vnd.openxmlformats-officedocument.spreadsheetml.template.main+xml")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelTemplateXLTX}
		case bytes.Contains(b, []byte("application/vnd.ms-excel.template.macroEnabled.main+xml")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelMacroEnabledTemplateXLTM}
		case bytes.Contains(b, []byte("application/vnd.ms-excel.addin.macroEnabled.main+xml")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelAddInXLAM}
		case bytes.Contains(b, []byte("application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelSpreadsheetXLSX}
		case bytes.Contains(b, []byte("application/vnd.ms-powerpoint.presentation.macroEnabled.main+xml")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointMacroEnabledPresentationPPTM}
		case bytes.Contains(b, []byte("application/vnd.openxmlformats-officedocument.presentationml.template.main+xml")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointTemplatePOTX}
		case bytes.Contains(b, []byte("application/vnd.ms-powerpoint.template.macroEnabled.main+xml")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointMacroEnabledTemplatePOTM}
		case bytes.Contains(b, []byte("application/vnd.openxmlformats-officedocument.presentationml.slideshow.main+xml")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointSlideshowPPSX}
		case bytes.Contains(b, []byte("application/vnd.ms-powerpoint.slideshow.macroEnabled.main+xml")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointMacroEnabledSlideshowPPSM}
		case bytes.Contains(b, []byte("application/vnd.ms-powerpoint.addin.macroEnabled.main+xml")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointAddInPPAM}
		case bytes.Contains(b, []byte("application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointPresentationPPTX}
		case bytes.Contains(b, []byte("word/")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordDocumentDOCX}
		case bytes.Contains(b, []byte("xl/")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelSpreadsheetXLSX}
		case bytes.Contains(b, []byte("ppt/")):
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftPowerPointPresentationPPTX}
		}
	}

	if bytes.Contains(b, []byte(".nuspec")) && bytes.Contains(b, []byte("package/services/metadata/core-properties")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeNuGetPackageNUPKG}
	}

	if lowerName == "bundleconfig.pb" || bytes.Contains(b, []byte("BundleConfig.pb")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidAppBundleAAB}
	}

	if lowerName == "toc.pb" || bytes.Contains(b, []byte("toc.pb")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidSplitAPKS}
	}

	if lowerName == "extension.vsixmanifest" || bytes.Contains(b, []byte("extension.vsixmanifest")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeVisualStudioExtensionVSIX}
	}

	if lowerName == "appxmanifest.xml" || bytes.Contains(b, []byte("AppxManifest.xml")) {
		if bytes.Contains(b, []byte("AppxBlockMap.xml")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAPPXPackage}
		}

		if bytes.Contains(b, []byte("AppxSignature.p7x")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMSIXPackage}
		}

		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAPPXPackage}
	}

	if lowerName == "doc.kml" {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeKMZArchive}
	}

	if strings.HasSuffix(lowerName, ".dist-info/wheel") || bytes.Contains(b, []byte(".dist-info/WHEEL")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypePythonWheelWHL}
	}

	if lowerName == "manifest.json" && bytes.Contains(b, []byte("xapk_version")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidPackageXAPK}
	}

	if lowerName == "install.rdf" || bytes.Contains(b, []byte("manifest.json")) && bytes.Contains(b, []byte("browser_specific_settings")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeFirefoxExtensionXPI}
	}

	if lowerName == "androidmanifest.xml" && bytes.Contains(b, []byte("classes.jar")) && !bytes.Contains(b, []byte("classes.dex")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidArchiveAAR}
	}

	if lowerName == "androidmanifest.xml" || bytes.Contains(b, []byte("classes.dex")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidPackageAPK}
	}

	if lowerName == "meta-inf/manifest.mf" || bytes.Contains(b, []byte("META-INF/MANIFEST.MF")) {
		if bytes.Contains(b, []byte("WEB-INF/web.xml")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeJavaWebArchiveWAR}
		}

		if bytes.Contains(b, []byte("META-INF/application.xml")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeJavaEnterpriseArchiveEAR}
		}

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
