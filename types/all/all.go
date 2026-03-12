package all

import (
	// Executables
	_ "github.com/coalaura/onda/types/dex"
	_ "github.com/coalaura/onda/types/elf"
	_ "github.com/coalaura/onda/types/java"
	_ "github.com/coalaura/onda/types/macho"
	_ "github.com/coalaura/onda/types/pe"
	_ "github.com/coalaura/onda/types/wasm"

	// Images
	_ "github.com/coalaura/onda/types/avif"
	_ "github.com/coalaura/onda/types/bmp"
	_ "github.com/coalaura/onda/types/dds"
	_ "github.com/coalaura/onda/types/gif"
	_ "github.com/coalaura/onda/types/hdr"
	_ "github.com/coalaura/onda/types/heif"
	_ "github.com/coalaura/onda/types/icns"
	_ "github.com/coalaura/onda/types/ico"
	_ "github.com/coalaura/onda/types/jp2"
	_ "github.com/coalaura/onda/types/jpeg"
	_ "github.com/coalaura/onda/types/png"
	_ "github.com/coalaura/onda/types/psd"
	_ "github.com/coalaura/onda/types/qoi"
	_ "github.com/coalaura/onda/types/tga"
	_ "github.com/coalaura/onda/types/tiff"
	_ "github.com/coalaura/onda/types/webp"

	// Audio
	_ "github.com/coalaura/onda/types/aac"
	_ "github.com/coalaura/onda/types/aiff"
	_ "github.com/coalaura/onda/types/amr"
	_ "github.com/coalaura/onda/types/au"
	_ "github.com/coalaura/onda/types/flac"
	_ "github.com/coalaura/onda/types/midi"
	_ "github.com/coalaura/onda/types/mp3"
	_ "github.com/coalaura/onda/types/ogg"
	_ "github.com/coalaura/onda/types/wav"

	// Video
	_ "github.com/coalaura/onda/types/avi"
	_ "github.com/coalaura/onda/types/flv"
	_ "github.com/coalaura/onda/types/mkv"
	_ "github.com/coalaura/onda/types/mp4"
	_ "github.com/coalaura/onda/types/mpeg"
	_ "github.com/coalaura/onda/types/webm"

	// Binaries
	_ "github.com/coalaura/onda/types/java"
	_ "github.com/coalaura/onda/types/wasm"

	// Documents / data
	_ "github.com/coalaura/onda/types/ole"
	_ "github.com/coalaura/onda/types/pdf"
	_ "github.com/coalaura/onda/types/pgp"
	_ "github.com/coalaura/onda/types/rtf"
	_ "github.com/coalaura/onda/types/sqlite"

	// Fonts
	_ "github.com/coalaura/onda/types/otf"
	_ "github.com/coalaura/onda/types/ttf"
	_ "github.com/coalaura/onda/types/woff"
	_ "github.com/coalaura/onda/types/woff2"

	// Archives / compression
	_ "github.com/coalaura/onda/types/7z"
	_ "github.com/coalaura/onda/types/ar"
	_ "github.com/coalaura/onda/types/bzip2"
	_ "github.com/coalaura/onda/types/cab"
	_ "github.com/coalaura/onda/types/gzip"
	_ "github.com/coalaura/onda/types/lz4"
	_ "github.com/coalaura/onda/types/rar"
	_ "github.com/coalaura/onda/types/tar"
	_ "github.com/coalaura/onda/types/xz"
	_ "github.com/coalaura/onda/types/zip"
	_ "github.com/coalaura/onda/types/zstd"

	// Capture / data
	_ "github.com/coalaura/onda/types/pcap"
	_ "github.com/coalaura/onda/types/pcapng"

	// Disk images
	_ "github.com/coalaura/onda/types/qcow2"
	_ "github.com/coalaura/onda/types/vdi"
	_ "github.com/coalaura/onda/types/vhd"
)
