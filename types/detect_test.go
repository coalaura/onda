package types_test

import (
	"errors"
	"testing"

	"github.com/coalaura/onda/types"
	_ "github.com/coalaura/onda/types/all"
)

func TestDetectJavaClassNotMachOUniversal(t *testing.T) {
	t.Parallel()

	data := []byte{0xca, 0xfe, 0xba, 0xbe, 0x00, 0x00, 0x00, 0x34, 0x00, 0x10}

	meta, err := types.Detect("Example.class", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindJavaClass {
		t.Fatalf("expected KindJavaClass, got %v", meta.Kind)
	}
}

func TestDetectMetafileAvoidsFalsePositive(t *testing.T) {
	t.Parallel()

	data := []byte{0x01, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00}

	_, err := types.Detect("random.bin", data)
	if !errors.Is(err, types.ErrUnknownFormat) {
		t.Fatalf("expected ErrUnknownFormat, got %v", err)
	}
}

func TestDetectEMF(t *testing.T) {
	t.Parallel()

	data := make([]byte, 88)
	data[0] = 0x01
	data[4] = 88
	copy(data[40:], []byte{0x20, 0x45, 0x4d, 0x46})

	meta, err := types.Detect("drawing.emf", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindMetafileImage || meta.Type != types.TypeEnhancedMetafileEMF {
		t.Fatalf("unexpected metadata: %+v", *meta)
	}
}

func TestDetectNetpbmRequiresDelimiter(t *testing.T) {
	t.Parallel()

	_, err := types.Detect("bad.ppm", []byte("P6NOT"))
	if !errors.Is(err, types.ErrUnknownFormat) {
		t.Fatalf("expected ErrUnknownFormat, got %v", err)
	}
}

func TestDetectNetpbmWithDelimiter(t *testing.T) {
	t.Parallel()

	meta, err := types.Detect("good.ppm", []byte("P6\n255\n"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindNetpbmImage || meta.Type != types.TypePPMBinary {
		t.Fatalf("unexpected metadata: %+v", *meta)
	}
}

func TestDetectISOCompatibleBrand(t *testing.T) {
	t.Parallel()

	data := []byte{
		0x00, 0x00, 0x00, 0x18,
		'f', 't', 'y', 'p',
		'i', 's', 'o', 'm',
		0x00, 0x00, 0x00, 0x00,
		'a', 'v', 'i', 'f',
		'm', 'i', 'f', '1',
	}

	meta, err := types.Detect("image.heif", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindISOBaseMedia || meta.Type != types.TypeAVIFImage {
		t.Fatalf("unexpected metadata: %+v", *meta)
	}
}

func TestDetectDocx(t *testing.T) {
	t.Parallel()

	data := makeZipLocalFile("[Content_Types].xml", nil)
	data = append(data, []byte("word/document.xml")...)

	meta, err := types.Detect("file.docx", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindZIPArchive || meta.Type != types.TypeMicrosoftWordDocumentDOCX {
		t.Fatalf("unexpected metadata: %+v", *meta)
	}
}

func TestDetectEPUB(t *testing.T) {
	t.Parallel()

	data := makeZipLocalFile("mimetype", []byte("application/epub+zip"))

	meta, err := types.Detect("book.epub", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindZIPArchive || meta.Type != types.TypeEPUBDocument {
		t.Fatalf("unexpected metadata: %+v", *meta)
	}
}

func TestDetectJar(t *testing.T) {
	t.Parallel()

	data := makeZipLocalFile("META-INF/MANIFEST.MF", nil)

	meta, err := types.Detect("app.jar", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindZIPArchive || meta.Type != types.TypeJavaArchiveJAR {
		t.Fatalf("unexpected metadata: %+v", *meta)
	}
}

func TestDetectOggOpus(t *testing.T) {
	t.Parallel()

	data := append([]byte("OggS"), make([]byte, 64)...)
	copy(data[32:], []byte("OpusHead"))

	meta, err := types.Detect("voice.ogg", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindOggContainer || meta.Type != types.TypeOpusAudio {
		t.Fatalf("unexpected metadata: %+v", *meta)
	}
}

func TestDetectMPEGTS(t *testing.T) {
	t.Parallel()

	data := make([]byte, 3*188)
	data[0] = 0x47
	data[188] = 0x47
	data[376] = 0x47

	meta, err := types.Detect("stream.ts", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindMPEGTransportStream || meta.Type != types.TypeTS {
		t.Fatalf("unexpected metadata: %+v", *meta)
	}
}

func TestDetectOLEWordDoc(t *testing.T) {
	t.Parallel()

	data := make([]byte, 512)
	copy(data, []byte{0xd0, 0xcf, 0x11, 0xe0, 0xa1, 0xb1, 0x1a, 0xe1})
	copy(data[128:], utf16LE("WordDocument"))

	meta, err := types.Detect("legacy.doc", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindOLECompoundDocument || meta.Type != types.TypeMicrosoftWordDocumentDOC {
		t.Fatalf("unexpected metadata: %+v", *meta)
	}
}

func TestDetectAAB(t *testing.T) {
	t.Parallel()

	data := makeZipLocalFile("BundleConfig.pb", nil)

	meta, err := types.Detect("app.aab", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindZIPArchive || meta.Type != types.TypeAndroidAppBundleAAB {
		t.Fatalf("unexpected metadata: %+v", *meta)
	}
}

func TestDetectWindowsEventLog(t *testing.T) {
	t.Parallel()

	meta, err := types.Detect("security.evtx", []byte("ElfFile\x00rest"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindWindowsEventLog {
		t.Fatalf("unexpected metadata: %+v", *meta)
	}
}

func TestDetectGIMPXCF(t *testing.T) {
	t.Parallel()

	meta, err := types.Detect("image.xcf", []byte("gimp xcf v003"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindGIMPXCFImage {
		t.Fatalf("unexpected metadata: %+v", *meta)
	}
}

func TestDetectDocm(t *testing.T) {
	t.Parallel()

	data := makeZipLocalFile("[Content_Types].xml", nil)
	data = append(data, []byte("application/vnd.ms-word.document.macroEnabled.main+xml")...)

	meta, err := types.Detect("file.docm", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindZIPArchive || meta.Type != types.TypeMicrosoftWordMacroEnabledDocumentDOCM {
		t.Fatalf("unexpected metadata: %+v", *meta)
	}
}

func TestDetectXlsm(t *testing.T) {
	t.Parallel()

	data := makeZipLocalFile("[Content_Types].xml", nil)
	data = append(data, []byte("application/vnd.ms-excel.sheet.macroEnabled.main+xml")...)

	meta, err := types.Detect("sheet.xlsm", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindZIPArchive || meta.Type != types.TypeMicrosoftExcelMacroEnabledWorkbookXLSM {
		t.Fatalf("unexpected metadata: %+v", *meta)
	}
}

func TestDetectPptm(t *testing.T) {
	t.Parallel()

	data := makeZipLocalFile("[Content_Types].xml", nil)
	data = append(data, []byte("application/vnd.ms-powerpoint.presentation.macroEnabled.main+xml")...)

	meta, err := types.Detect("slides.pptm", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindZIPArchive || meta.Type != types.TypeMicrosoftPowerPointMacroEnabledPresentationPPTM {
		t.Fatalf("unexpected metadata: %+v", *meta)
	}
}

func TestDetectWar(t *testing.T) {
	t.Parallel()

	data := makeZipLocalFile("META-INF/MANIFEST.MF", nil)
	data = append(data, []byte("WEB-INF/web.xml")...)

	meta, err := types.Detect("app.war", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindZIPArchive || meta.Type != types.TypeJavaWebArchiveWAR {
		t.Fatalf("unexpected metadata: %+v", *meta)
	}
}

func TestDetectNuGet(t *testing.T) {
	t.Parallel()

	data := makeZipLocalFile("package.nuspec", nil)
	data = append(data, []byte("package/services/metadata/core-properties")...)

	meta, err := types.Detect("pkg.nupkg", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindZIPArchive || meta.Type != types.TypeNuGetPackageNUPKG {
		t.Fatalf("unexpected metadata: %+v", *meta)
	}
}

func TestDetectJMOD(t *testing.T) {
	t.Parallel()

	meta, err := types.Detect("java.base.jmod", []byte("JMOD\x00\x00"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if meta.Kind != types.KindJavaModule || meta.Type != types.TypeJMOD {
		t.Fatalf("unexpected metadata: %+v", *meta)
	}
}

func makeZipLocalFile(name string, data []byte) []byte {
	buf := make([]byte, 30+len(name)+len(data))

	copy(buf[0:], []byte{'P', 'K', 3, 4})
	buf[8] = 0
	buf[9] = 0
	buf[18] = byte(len(data))
	buf[19] = byte(len(data) >> 8)
	buf[20] = byte(len(data) >> 16)
	buf[21] = byte(len(data) >> 24)
	buf[22] = buf[18]
	buf[23] = buf[19]
	buf[24] = buf[20]
	buf[25] = buf[21]
	buf[26] = byte(len(name))
	buf[27] = byte(len(name) >> 8)

	copy(buf[30:], []byte(name))
	copy(buf[30+len(name):], data)

	return buf
}

func utf16LE(s string) []byte {
	b := make([]byte, len(s)*2)

	for i := 0; i < len(s); i++ {
		b[i*2] = s[i]
	}

	return b
}
