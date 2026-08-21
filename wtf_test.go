package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/coalaura/wtf/types"
)

type testCase struct {
	file string
	kind types.KindID
	typ  types.TypeID
}

func TestConsistency(t *testing.T) {
	typeNames := getIdentifiers(t, "types/ids_type.go", "Type")
	kindNames := getIdentifiers(t, "types/ids_kind.go", "Kind")

	for i := types.TypeID(1); i < types.MaxTypes; i++ {
		if i.String() == "" {
			name := "Unknown Identifier"
			if int(i) < len(typeNames) {
				name = typeNames[i]
			}

			t.Errorf("Type %s (%d) is missing a label in typeNames array", name, i)
		}
	}

	for i := types.KindID(1); i < types.MaxKinds; i++ {
		if i.String() == "Unknown" {
			name := "Unknown Identifier"
			if int(i) < len(kindNames) {
				name = kindNames[i]
			}

			t.Errorf("Kind %s (%d) is missing a label in kindNames array", name, i)
		}
	}
}

func TestDetectFixtures(t *testing.T) {
	cases := []testCase{
		{file: "3g2.3g2", kind: types.KindISOBaseMedia, typ: types.Type3G2Video},
		{file: "3gp.3gp", kind: types.KindISOBaseMedia, typ: types.Type3GPPVideo},
		{file: "7z.7z", kind: types.Kind7ZipArchive, typ: types.TypeNone},
		{file: "aifc.aif", kind: types.KindIFFContainer, typ: types.TypeAIFCAudio},
		{file: "android.apk", kind: types.KindZIPArchive, typ: types.TypeAndroidPackage},
		{file: "archive.zst", kind: types.KindZstandardArchive, typ: types.TypeNone},
		{file: "arw.arw", kind: types.KindTIFFImage, typ: types.TypeSonyRAW},
		{file: "audio.flac", kind: types.KindFLACAudio, typ: types.TypeNone},
		{file: "avi.avi", kind: types.KindRIFFContainer, typ: types.TypeAVIVideo},
		{file: "avif.avif", kind: types.KindISOBaseMedia, typ: types.TypeAVIFImageSequence},
		{file: "big_endian.tif", kind: types.KindTIFFImage, typ: types.TypeBigEndian},
		{file: "bigtiff.tif", kind: types.KindTIFFImage, typ: types.TypeBigTIFF},
		{file: "bundle.aab", kind: types.KindZIPArchive, typ: types.TypeAndroidAppBundle},
		{file: "comic.cbz", kind: types.KindZIPArchive, typ: types.TypeComicBook},
		{file: "cr2.cr2", kind: types.KindCanonRAWImage, typ: types.TypeHE},
		{file: "cursor.cur", kind: types.KindWindowsCursorImage, typ: types.TypeNone},
		{file: "data.json", kind: types.KindJSONDocument, typ: types.TypeNone},
		{file: "dicom.dcm", kind: types.KindDICOMMedicalImage, typ: types.TypeNone},
		{file: "dng.dng", kind: types.KindTIFFImage, typ: types.TypeAdobeDNG},
		{file: "document.doc", kind: types.KindOLECompoundDocument, typ: types.TypeMicrosoftWordDocument},
		{file: "dp3_icc.icc", kind: types.KindICCProfile, typ: types.TypeNone},
		{file: "elf_x64", kind: types.KindExecutableAndLinkableFormat, typ: types.Type64BitLittleEndian},
		{file: "elf_x86", kind: types.KindExecutableAndLinkableFormat, typ: types.Type32BitLittleEndian},
		{file: "epub.epub", kind: types.KindZIPArchive, typ: types.TypeEPUBDocument},
		{file: "excel.xls", kind: types.KindOLECompoundDocument, typ: types.TypeMicrosoftExcelWorkbook},
		{file: "excel.xlsx", kind: types.KindZIPArchive, typ: types.TypeMicrosoftExcelWorkbook},
		{file: "gif87a.gif", kind: types.KindGIFImage, typ: types.TypeGIF87a},
		{file: "gif89a.gif", kind: types.KindGIFImage, typ: types.TypeGIF89a},
		{file: "gz.gz", kind: types.KindGzipArchive, typ: types.TypeNone},
		{file: "heif.heif", kind: types.KindISOBaseMedia, typ: types.TypeHEIFImage},
		{file: "ico.ico", kind: types.KindWindowsIconImage, typ: types.TypeNone},
		{file: "installer.msi", kind: types.KindOLECompoundDocument, typ: types.TypeMicrosoftInstaller},
		{file: "jar.jar", kind: types.KindZIPArchive, typ: types.TypeJavaArchive},
		{file: "java.class", kind: types.KindJavaClassBytecode, typ: types.TypeNone},
		{file: "jpeg.jpeg", kind: types.KindJPEGImage, typ: types.TypeNone},
		{file: "jpegxl.jxl", kind: types.KindJPEGXLImage, typ: types.TypeCodestream},
		{file: "little_endian.tif", kind: types.KindTIFFImage, typ: types.TypeLittleEndian},
		{file: "m4a.m4a", kind: types.KindISOBaseMedia, typ: types.TypeMPEG4Audio},
		{file: "macho", kind: types.KindMachOBinary, typ: types.Type64BitLittleEndian},
		{file: "mkv.mkv", kind: types.KindEBMLContainer, typ: types.TypeMatroskaVideo},
		{file: "module.wasm", kind: types.KindWebAssemblyModule, typ: types.TypeNone},
		{file: "mp4.mp4", kind: types.KindISOBaseMedia, typ: types.TypeMP4Video},
		{file: "mpeg2.ts", kind: types.KindMPEGTransportStream, typ: types.TypeNone},
		{file: "mpeg3.mp3", kind: types.KindMPEGAudio, typ: types.TypeMP3ID3Tagged},
		{file: "nef.nef", kind: types.KindTIFFImage, typ: types.TypeNikonRAW},
		{file: "nifti.nii", kind: types.KindNIfTIMedicalImage, typ: types.TypeNone},
		{file: "ogg.ogg", kind: types.KindOggContainer, typ: types.TypeVorbisAudio},
		{file: "open-doc.odt", kind: types.KindZIPArchive, typ: types.TypeOpenDocumentText},
		{file: "open-spread.ods", kind: types.KindZIPArchive, typ: types.TypeOpenDocumentSpreadsheet},
		{file: "opus.opus", kind: types.KindOggContainer, typ: types.TypeOpusAudio},
		{file: "package.deb", kind: types.KindDebianPackage, typ: types.TypeNone},
		{file: "pcap.pcap", kind: types.KindPCAPCapture, typ: types.TypeLittleEndian},
		{file: "pcapng.pcapng", kind: types.KindPCAPNGCapture, typ: types.TypeNone},
		{file: "pcx.pcx", kind: types.KindPCXImage, typ: types.TypeNone},
		{file: "pdf.pdf", kind: types.KindPDFDocument, typ: types.TypeNone},
		{file: "pe_x64", kind: types.KindPortableExecutable, typ: types.TypePE32PlusX8664},
		{file: "pe_x86", kind: types.KindPortableExecutable, typ: types.TypePE32X86},
		{file: "png.png", kind: types.KindPNGImage, typ: types.TypeNone},
		{file: "powerpoint.ppt", kind: types.KindOLECompoundDocument, typ: types.TypeMicrosoftPowerPointPresentation},
		{file: "powerpoint.pptx", kind: types.KindZIPArchive, typ: types.TypeMicrosoftPowerPointPresentation},
		{file: "pptm.pptm", kind: types.KindZIPArchive, typ: types.TypeMicrosoftPowerPointMacroEnabledPresentation},
		{file: "quicktime.mov", kind: types.KindISOBaseMedia, typ: types.TypeQuickTimeMovie},
		{file: "rar4.rar", kind: types.KindRARArchive, typ: types.TypeLegacy},
		{file: "rar5.rar", kind: types.KindRARArchive, typ: types.TypeVersion5},
		{file: "sqlite.db-shm", kind: types.KindSQLite3SharedMemory, typ: types.TypeLittleEndian},
		{file: "sqlite.db-wal", kind: types.KindSQLite3WriteAheadLog, typ: types.TypeLittleEndian},
		{file: "sqlite.db", kind: types.KindSQLiteDatabase, typ: types.TypeNone},
		{file: "svg.svg", kind: types.KindSVGImage, typ: types.TypeNone},
		{file: "tar.tar", kind: types.KindTARArchive, typ: types.TypeNone},
		{file: "torrent.torrent", kind: types.KindTorrentDocument, typ: types.TypeNone},
		{file: "universal.macho", kind: types.KindMachOUniversalBinary, typ: types.TypeNone},
		{file: "war.war", kind: types.KindZIPArchive, typ: types.TypeJavaWebArchive},
		{file: "wav.wav", kind: types.KindRIFFContainer, typ: types.TypeWAVAudio},
		{file: "webm.webm", kind: types.KindEBMLContainer, typ: types.TypeWebMVideo},
		{file: "webp.webp", kind: types.KindRIFFContainer, typ: types.TypeWebPImage},
		{file: "word.docx", kind: types.KindZIPArchive, typ: types.TypeMicrosoftWordDocument},
		{file: "xlsm.xlsm", kind: types.KindZIPArchive, typ: types.TypeMicrosoftExcelMacroEnabledWorkbook},
		{file: "zip.zip", kind: types.KindZIPArchive, typ: types.TypeNone},
	}

	for _, tc := range cases {
		t.Run(tc.file, func(t *testing.T) {
			path := filepath.Join("test", tc.file)

			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("read fixture: %v", err)
			}

			meta, err := types.Detect(tc.file, data)
			if err != nil {
				t.Fatalf("detect: %v", err)
			}

			if meta == nil {
				t.Fatal("detect returned nil metadata")
			}

			if meta.Kind != tc.kind {
				t.Fatalf("kind mismatch: got %v (%q), want %v (%q)", meta.Kind, meta.Kind.String(), tc.kind, tc.kind.String())
			}

			if meta.Type != tc.typ {
				t.Fatalf("type mismatch: got %v (%q), want %v (%q)", meta.Type, meta.Type.String(), tc.typ, tc.typ.String())
			}
		})
	}
}

func getIdentifiers(t *testing.T, filename, prefix string) []string {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		t.Fatalf("Failed to parse %s: %v", filename, err)
	}

	var names []string

	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.CONST {
			continue
		}

		for _, spec := range genDecl.Specs {
			vs, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			for _, name := range vs.Names {
				if strings.HasPrefix(name.Name, prefix) {
					names = append(names, name.Name)
				}
			}
		}
	}

	return names
}
