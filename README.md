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
7-Zip archive, AR archive, Android app bundle (AAB), Android archive (AAR), Android package (APK), Android split APK set (APKS), Android package (XAPK), APPX package, AppleDouble file, AppleSingle file, Btrfs filesystem, Bzip2 archive, Cabinet archive, Conda package (tar), CPIO archive (new ASCII, new ASCII with CRC, old ASCII, binary big-endian, binary little-endian), Debian package, ext filesystem (ext2, ext3, ext4), Firefox extension (XPI), Git index, Git pack, Gzip archive, iOS application archive (IPA), Java archive (JAR), Java enterprise archive (EAR), Java web archive (WAR), KMZ archive, LHA archive, LZ4 frame, LZIP archive, LZOP archive, MacBinary, MSIX package, NuGet package (NUPKG), npm package tarball, OCI image layout (tar), Python source distribution (sdist), Python wheel (WHL), RAR archive (RAR4, RAR5), RPM package, RubyGem package, Sketch document, Snappy framed data, SquashFS filesystem, StuffIt archive, TAR archive, Unix compress archive, Visual Studio extension (VSIX), Windows Imaging Format, XAR archive, XFS filesystem, XZ archive, ZIP archive (standard, empty, spanned), Zstandard archive.

#### Audio/tracker:
AAC audio (ADTS), AC-3, E-AC-3, AIFF audio (AIFF, AIFC), AMR audio, AMR-WB audio, AU audio, CAF audio, CD audio (CDA), DSF audio, DTS audio, FLAC audio, Impulse Tracker module, FastTracker module, MIDI sequence, Monkey's Audio, MP3 audio (ID3 tagged, MPEG audio frame), MPEG Layer II, MPEG Layer III, Musepack audio (stream version 7, stream version 8), Ogg container, Opus audio, QCP audio, Scream Tracker module, Speex audio, TAK audio, TTA audio, VOC audio, Vorbis audio, WAV audio, WavPack audio, MPEG-4 audio (M4A family).

#### Image/texture/icon:
AVIF image (single image, sequence), BMP image, BPG image, Canon RAW image (CR2, CR3), DDS image, DjVu document, Enhanced Metafile image, Farbfeld image, Fujifilm RAW image (RAF), GIF image (GIF87a, GIF89a), GIMP XCF image, glTF binary (GLB), HEIF image, ICNS icon, ICC profile, JPEG image, JPEG 2000 image, JPEG XL image (codestream, container), JNG image, JPEG XR image (little-endian, big-endian), KTX texture (KTX, KTX2), Netpbm image (PBM ASCII, PGM ASCII, PPM ASCII, PBM binary, PGM binary, PPM binary, PAM), Nikon RAW image (NEF), Olympus RAW image (ORF), OpenEXR image, OpenRaster image (ORA), Panasonic RAW image (RW2), Pentax RAW image (PEF), PCX image, PNG image, Photoshop document (PSD, PSB), QOI image, Radiance HDR image, Sony RAW image (ARW/SR2), Sun raster image, SVG image, TIFF image (little-endian, big-endian), Adobe DNG image, WebP image, Windows Metafile image, Windows icon, Windows cursor.

#### Video/container:
3GPP media, 3G2 media, ASF container, AVI video, EBML container, FLV video, F4V video, IVF video, ISO Base Media file, M3U8 playlist, M4V video, Matroska container, MP4 video, MPEG program stream, MPEG transport stream (TS, M2TS, M2TS/BDAV), QuickTime movie, RealMedia, RIFF container, Theora video, WebM container, WTV video.

#### Document/data/database:
Apache Arrow file, Apache Parquet, Apple binary property list, Apple iWork document, Avro object container, BAM data, CRAM data, DICOM medical image, EPUB document, FITS astronomical image, GRIB data, HDF5 data, iCalendar (ICS), KeePass database (KDBX), Microsoft Excel workbook (XLS), Microsoft Excel spreadsheet (XLSX), Microsoft Excel macro-enabled workbook/template/add-in (XLSM, XLTM, XLAM), Microsoft Outlook email folder (PST/OST), Microsoft PowerPoint presentation (PPT, PPTX), Microsoft PowerPoint macro-enabled presentation/template/slideshow/add-in (PPTM, POTM, PPSM, PPAM), Microsoft Word document (DOC, DOCX), Microsoft Word template/macro-enabled variants (DOTX, DOCM, DOTM), MOBI document, NetCDF data, OLE compound document, OpenDocument presentation (ODP), OpenDocument spreadsheet (ODS), OpenDocument text (ODT), ORC columnar data, PDF document, PEM certificate, PEM private key, PGP message, PostScript document, Rich Text Format document, SQLite database, Torrent file, vCard (VCF), WARC file, Windows registry hive, XML document.

#### Font:
EOT font, OpenType font, TrueType font, TrueType collection, WOFF font, WOFF2 font.

#### Executable/system/disk:
Android boot image, Android oat, Android odex, Android sparse image, AppImage, Blender file, CRX browser extension (v2/v3), Dalvik executable (DEX 035, DEX 036, DEX 037, DEX 038, DEX 039, DEX 040, DEX 041), Executable and Linkable Format (ELF/ELF32/ELF64, little-endian/big-endian), ISO 9660 image, Java class, Java KeyStore, Java module (JMOD), LLVM bitcode (raw, wrapper), Lua bytecode, Mach-O binary (32-bit/64-bit, little-endian/big-endian), Mach-O universal binary (32-bit, 64-bit), NES ROM, PCAP capture (little-endian, big-endian, nanosecond little-endian, nanosecond big-endian), PCAPNG capture, PKCS#12, Portable Executable (PE32/PE32+), Python bytecode, QCOW disk image (QCOW2), Shebang script, Shockwave Flash (uncompressed, zlib compressed, lzma compressed), U-Boot image, VHD disk image, VirtualBox disk image, VHDX disk image, VMware disk image (VMDK), WebAssembly module, Windows event log (EVTX), Windows shortcut.

#### Text fallback:
ASCII text, UTF-8 text.
