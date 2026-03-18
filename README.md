![img](banner.png)

onda is a tiny, "hardware-accelerated" file sniffer for Go.

## Blazing Fast

onda achieves sub-millisecond detection times through extreme mechanical sympathy. Instead of allocating memory, iterating slices or using locks at runtime, onda uses a custom Ahead-of-Time (AOT) compiler to generate a deeply nested Radix Trie (Prefix Tree) in pure Go.

The Go compiler flattens this tree into a highly optimized jump table in assembly. The CPU branch predictor routes file signatures in nanoseconds-resulting in **zero-allocation startup**, **zero runtime locks** and $O(1)$ time complexity for 95% of files.

```bash
$ time onda onda
onda
  └─ Executable and Linkable Format
     Type: ELF64 Little-Endian

real	0m0.002s
user	0m0.000s
sys 	0m0.002s
```
*(Benchmark ran on an AMD Ryzen 7 7840U / CachyOS Linux)*

## Features

- **Hardware-Accelerated Hot Path**: $O(1)$ magic-byte detection via AOT-compiled jump tables.
- **Zero-Cost Abstraction**: Static signatures are stripped from the runtime binary, saving memory and `init()` overhead.
- **Smart Fallbacks**: Includes custom structural parsers for complex formats (SVG, ZIP, XML, Text) that lack fixed headers.
- **Versatile**: Works as both a standalone CLI and a lightweight Go package.

## Installation

Prebuilt releases are available [here](https://github.com/coalaura/onda/releases). You can bootstrap **onda** with a single command. This script will detect your OS and CPU (`amd64`/`arm64`), download the correct binary and install it to `/usr/local/bin/onda`.

```bash
curl -sL https://src.ws2.sh/onda/install.sh | sh
```

Or install it via Go:

```bash
go install github.com/coalaura/onda@latest
```

## Usage

```bash
onda [flags] <file>
```

**Flags:**
- `-p`, `--porcelain`: Print easily parseable output (tab-separated: `Kind\tType`)
- `-v`, `--version`: Print version information
- `-h`, `--help`: Print this help message

**Examples:**

```bash
onda sample.png
onda --porcelain sample.png
```

## Go package

```go
package main

import (
	"fmt"

	"github.com/coalaura/onda/types"
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

onda currently detects over **510+ file formats** across various categories. Instead of listing every single format, we maintain a robust, ever-growing library of signatures and custom detectors.

- **Archive/package/filesystem:** 7-Zip, APFS, APK, Btrfs, Bzip2, CAB, Debian, exFAT, ext2/3/4, Gzip, HFS+, JAR, LUKS, LZ4, NTFS, RAR, RPM, SquashFS, TAR, XFS, XZ, ZIP, Zstandard and many more.
- **Audio/tracker:** AAC, AIFF, FLAC, MIDI, MP3, Ogg, Opus, Vorbis, WAV, WavPack and more.
- **Image/textuinternalre/icon:** AVIF, BMP, DDS, GIF, HEIF, ICO, JPEG, JPEG XL, PNG, PSD, SVG, TIFF, WebP and more.
- **Video/container:** AVI, FLV, Matroska (MKV), MP4, MPEG, QuickTime, WebM and more.
- **Document/data:** Apache Arrow/Parquet, DICOM, EPUB, HDF5, Microsoft Office (DOCX, XLSX, PPTX), PDF, RTF, SQLite, XML and more.
- **Font:** EOT, OpenType, TrueType, WOFF, WOFF2.
- **Executable/system/disk:** AppImage, Dalvik (DEX), ELF, Java Class, LLVM, Mach-O, Nintendo ROMs, PE, PCAP, QCOW, VHD, VMDK, WebAssembly and more.
- **Text fallback:** ASCII and UTF-8 text detection.
