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

	if meta.Name != "Java Class" {
		t.Fatalf("expected Java Class, got %q", meta.Name)
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

	if meta.Name != "Metafile Image" || meta.Type != "Enhanced Metafile (EMF)" {
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

	if meta.Name != "Netpbm Image" || meta.Type != "PPM binary" {
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

	if meta.Name != "ISO Base Media" || meta.Type != "AVIF Image" {
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

	if meta.Name != "ZIP Archive" || meta.Type != "Microsoft Word Document (DOCX)" {
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

	if meta.Name != "ZIP Archive" || meta.Type != "EPUB Document" {
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

	if meta.Name != "ZIP Archive" || meta.Type != "Java Archive (JAR)" {
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

	if meta.Name != "Ogg Container" || meta.Type != "Opus Audio" {
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

	if meta.Name != "MPEG Transport Stream" || meta.Type != "TS" {
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
