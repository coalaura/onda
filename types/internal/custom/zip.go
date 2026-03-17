package custom

import (
	"bytes"

	"github.com/coalaura/onda/types"
)

func DetectZIPContainer(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte{'P', 'K', 3, 4}) {
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

	for i := 0; i <= limit; {
		idx := bytes.Index(b[i:], []byte{'P', 'K', 3, 4})
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

		if firstFile && bytes.EqualFold(name, []byte("mimetype")) {
			compression, _ := b.U16LE(i + 8)
			extraLen, _ := b.U16LE(i + 28)
			dataLen, _ := b.U32LE(i + 18)
			dataStart := i + 30 + int(nameLen) + int(extraLen)
			dataEnd := dataStart + int(dataLen)

			if compression == 0 && dataStart >= 0 && dataEnd <= b.Len() {
				data := b[dataStart:dataEnd]

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

		firstFile = false

		if bytes.EqualFold(name, []byte("androidmanifest.xml")) {
			hasManifest = true
		} else if bytes.HasSuffix(name, []byte(".dex")) || bytes.EqualFold(name, []byte("classes.dex")) {
			hasDex = true
		} else if bytes.EqualFold(name, []byte("appxmanifest.xml")) {
			hasAppxManifest = true
		} else if bytes.EqualFold(name, []byte("appxsignature.p7x")) {
			hasAppxSignature = true
		} else if bytes.EqualFold(name, []byte("document.json")) {
			hasSketchDoc = true
		} else if bytes.EqualFold(name, []byte("meta.json")) {
			hasSketchMeta = true
		} else if bytes.EqualFold(name, []byte("user.json")) {
			hasSketchUser = true
		} else if bytes.EqualFold(name, []byte("doc.kml")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeKMZArchive}
		} else if bytes.HasSuffix(name, []byte(".dist-info/wheel")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypePythonWheelWHL}
		} else if bytes.EqualFold(name, []byte("manifest.json")) && bytes.Contains(b, []byte("xapk_version")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidPackageXAPK}
		} else if bytes.EqualFold(name, []byte("install.rdf")) || (bytes.EqualFold(name, []byte("manifest.json")) && bytes.Contains(b, []byte("browser_specific_settings"))) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeFirefoxExtensionXPI}
		} else if bytes.EqualFold(name, []byte("meta-inf/manifest.mf")) {
			if bytes.Contains(b, []byte("WEB-INF/web.xml")) {
				return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeJavaWebArchiveWAR}
			}

			if bytes.Contains(b, []byte("META-INF/application.xml")) {
				return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeJavaEnterpriseArchiveEAR}
			}

			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeJavaArchiveJAR}
		} else if bytes.HasPrefix(name, []byte("payload/")) || bytes.HasPrefix(name, []byte("Payload/")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeIOSApplicationArchiveIPA}
		} else if bytes.EqualFold(name, []byte("bundleconfig.pb")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidAppBundleAAB}
		} else if bytes.EqualFold(name, []byte("toc.pb")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidSplitAPKS}
		} else if bytes.Contains(name, []byte("apex_manifest")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeAndroidSystemPackageAPEX}
		} else if bytes.EqualFold(name, []byte("extension.vsixmanifest")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeVisualStudioExtensionVSIX}
		} else if bytes.HasSuffix(name, []byte(".nuspec")) && bytes.Contains(b, []byte("package/services/metadata/core-properties")) {
			return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeNuGetPackageNUPKG}
		}

		if len(name) >= 5 && bytes.EqualFold(name[:5], []byte("word/")) {
			hasWord = true
		} else if len(name) >= 3 && bytes.EqualFold(name[:3], []byte("xl/")) {
			hasExcel = true
		} else if len(name) >= 4 && bytes.EqualFold(name[:4], []byte("ppt/")) {
			hasPowerPoint = true
		}

		i += 30 + int(nameLen)
	}

	// Word
	if bytes.Contains(b, []byte("application/vnd.ms-word.document.macroEnabled.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordMacroEnabledDocumentDOCM}
	}

	if bytes.Contains(b, []byte("application/vnd.openxmlformats-officedocument.wordprocessingml.template.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordTemplateDOTX}
	}

	if bytes.Contains(b, []byte("application/vnd.ms-word.template.macroEnabledTemplate.main+xml")) {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordMacroEnabledTemplateDOTM}
	}

	if hasWord {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftWordDocumentDOCX}
	}

	// Excel
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

	if hasExcel {
		return &types.Metadata{Kind: types.KindZIPArchive, Type: types.TypeMicrosoftExcelSpreadsheetXLSX}
	}

	// PowerPoint
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
		if bytes.Contains(b, []byte("com.bohemiancoding.sketch")) || bytes.Contains(b, []byte("com.sketch3")) {
			return &types.Metadata{Kind: types.KindSketchDocument, Type: types.TypeSketchDocument}
		}
	}

	return &types.Metadata{Kind: types.KindZIPArchive}
}
