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

	if meta.Name != "Java class" {
		t.Fatalf("expected Java class, got %q", meta.Name)
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

	if meta.Name != "Enhanced Metafile image" {
		t.Fatalf("expected Enhanced Metafile image, got %q", meta.Name)
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

	if meta.Name != "Netpbm image" || meta.Type != "PPM binary" {
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

	if meta.Name != "AVIF image" {
		t.Fatalf("expected AVIF image, got %q", meta.Name)
	}
}
