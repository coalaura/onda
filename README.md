# onda

onda is a tiny file type detector for Go.

## Features

- Fast magic-byte/signature based detection
- Supports many common formats (images, archives, audio, video, executables, documents, disk images)
- Includes custom parsers for formats that need structural checks (not just fixed headers)
- Works as both a CLI and a Go package

## Installation

Prebuilt releases are available [here](https://github.com/coalaura/onda/releases) or build it yourself.

## Usage

```bash
onda <file>
```

Example:

```bash
onda sample.png
```

## Go package

```go
package main

import (
	"fmt"

	"github.com/coalaura/onda/types"
	_ "github.com/coalaura/onda/types/all"
)

func main() {
	meta, err := types.Detect("sample.png", []byte{0x89, 'P', 'N', 'G', 0x0d, 0x0a, 0x1a, 0x0a})
	if err != nil {
		fmt.Println("unknown")

		return
	}

	fmt.Println(meta.Kind.String(), meta.Type.String())
}
```

## Supported types

#### Archive/package/filesystem:
7-Zip archive, AR archive, Android package (APK), Bzip2 archive, Cabinet archive, CPIO archive (new ASCII, new ASCII with CRC, old ASCII, binary big-endian, binary little-endian), Debian package, Gzip archive, iOS application archive (IPA), Java archive (JAR), LZ4 frame, LZIP archive, LZOP archive, RAR archive (RAR4, RAR5), RPM package, Snappy framed data, SquashFS filesystem, TAR archive, Unix compress archive, Windows Imaging Format, XAR archive, XZ archive, ZIP archive (standard, empty, spanned), Zstandard archive.

#### Audio/tracker:
AAC audio (ADTS), AIFF audio (AIFF, AIFC), AMR audio, AMR-WB audio, AU audio, CAF audio, DSF audio, FLAC audio, Impulse Tracker module, FastTracker module, MIDI sequence, Monkey's Audio, MP3 audio (ID3 tagged, MPEG audio frame), Musepack audio (stream version 7, stream version 8), Ogg container, Opus audio, Scream Tracker module, Speex audio, TAK audio, VOC audio, Vorbis audio, WAV audio, WavPack audio, MPEG-4 audio (M4A family).

#### Image/texture/icon:
AVIF image (single image, sequence), BMP image, BPG image, DDS image, Enhanced Metafile image, Farbfeld image, GIF image (GIF87a, GIF89a), HEIF image, ICNS icon, JPEG image, JPEG 2000 image, JPEG XL image (codestream, container), JPEG XR image (little-endian, big-endian), KTX texture (KTX, KTX2), Netpbm image (PBM ASCII, PGM ASCII, PPM ASCII, PBM binary, PGM binary, PPM binary, PAM), OpenEXR image, PNG image, Photoshop document (PSD, PSB), QOI image, Radiance HDR image, Sun raster image, TIFF image (little-endian, big-endian), WebP image, Windows Metafile image, Windows icon, Windows cursor.

#### Video/container:
3GPP media, ASF container, AVI video, EBML container, FLV video, IVF video, ISO Base Media file, Matroska container, MP4 video, MPEG program stream, MPEG transport stream (TS, M2TS), QuickTime movie, RIFF container, Theora video, WebM container.

#### Document/data/database:
Apache Arrow file, Apache Parquet, Apple binary property list, Avro object container, DICOM medical image, EPUB document, FITS astronomical image, HDF5 data, KeePass database (KDBX), Microsoft Excel spreadsheet (XLSX), Microsoft PowerPoint presentation (PPTX), Microsoft Word document (DOCX), OLE compound document, OpenDocument presentation (ODP), OpenDocument spreadsheet (ODS), OpenDocument text (ODT), ORC columnar data, PDF document, PGP message, PostScript document, Rich Text Format document, SQLite database, Windows registry hive.

#### Font:
OpenType font, TrueType font, TrueType collection, WOFF font, WOFF2 font.

#### Executable/system/disk:
Android oat, Android odex, Android sparse image, CRX browser extension, Dalvik executable (DEX 035, DEX 036, DEX 037, DEX 038, DEX 039, DEX 040, DEX 041), Executable and Linkable Format (ELF32/ELF64, little-endian/big-endian), ISO 9660 image, Java class, Java KeyStore, LLVM bitcode (raw, wrapper), Mach-O binary (32-bit/64-bit, little-endian/big-endian), Mach-O universal binary (32-bit, 64-bit), NES ROM, PCAP capture (little-endian, big-endian, nanosecond little-endian, nanosecond big-endian), PCAPNG capture, Portable Executable (PE32/PE32+), QCOW disk image (QCOW2), Shockwave Flash (uncompressed, zlib compressed, lzma compressed), VirtualBox disk image, VHDX disk image, VMware disk image (VMDK), WebAssembly module, Windows shortcut.
