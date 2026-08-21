package custom

import (
	"bytes"

	"github.com/coalaura/wtf/types"
)

func DetectTar(b types.Buffer) *types.Metadata {
	var isTar bool

	limit := min(b.Len(), maxScanSize)

	for offset := 0; offset+512 <= limit; {
		header := b[offset : offset+512]
		if string(header[257:262]) != "ustar" || !validTarChecksum(header) {
			break
		}

		isTar = true

		nameEnd := bytes.IndexByte(header[:100], 0)
		if nameEnd == -1 {
			nameEnd = 100
		}

		if nameEnd > 0 {
			switch string(header[:nameEnd]) {
			case "package/package.json", "package.json":
				return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeNpmPackage}
			case "oci-layout", "index.json", "manifest.json":
				return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeOCIImageLayout}
			case "PKG-INFO", "setup.py", "pyproject.toml":
				return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypePythonSourceDistribution}
			case "info/index.json":
				return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeCondaPackage}
			case ".PKGINFO":
				return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeArchLinuxPackage}
			case "Vagrantfile":
				return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeVagrantBox}
			case "install/doinst.sh":
				return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeSlackwarePackage}
			case "ComicInfo.xml", "comicinfo.xml":
				return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeComicBook}
			case "metadata", "deploy":
				return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeFlatpak}
			}
		}

		size, ok := parseTarOctal(header[124:136])
		if !ok {
			break
		}

		next := uint64(offset+512) + ((size + 511) / 512 * 512)
		if next > uint64(limit) {
			break
		}

		offset = int(next)
	}

	if isTar {
		return &types.Metadata{Kind: types.KindTARArchive}
	}

	return nil
}

func validTarChecksum(header []byte) bool {
	stored, ok := parseTarOctal(header[148:156])
	if !ok {
		return false
	}

	var sum uint64

	for i, value := range header {
		if i >= 148 && i < 156 {
			value = ' '
		}

		sum += uint64(value)
	}

	return sum == stored
}

func parseTarOctal(field []byte) (uint64, bool) {
	field = bytes.Trim(field, " \x00")
	if len(field) == 0 {
		return 0, false
	}

	var value uint64

	for _, digit := range field {
		if digit < '0' || digit > '7' {
			return 0, false
		}

		value = value*8 + uint64(digit-'0')
	}

	return value, true
}
