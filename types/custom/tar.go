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

		prefixEnd := bytes.IndexByte(header[345:500], 0)
		if prefixEnd == -1 {
			prefixEnd = 155
		}

		if nameEnd > 0 {
			name := header[:nameEnd]
			prefix := header[345 : 345+prefixEnd]

			switch {
			case tarPathEqual(prefix, name, "package/package.json"), tarPathEqual(prefix, name, "package.json"):
				return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeNpmPackage}
			case tarPathEqual(prefix, name, "oci-layout"), tarPathEqual(prefix, name, "index.json"), tarPathEqual(prefix, name, "manifest.json"):
				return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeOCIImageLayout}
			case tarPathEqual(prefix, name, "PKG-INFO"), tarPathEqual(prefix, name, "setup.py"), tarPathEqual(prefix, name, "pyproject.toml"):
				return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypePythonSourceDistribution}
			case tarPathEqual(prefix, name, "info/index.json"):
				return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeCondaPackage}
			case tarPathEqual(prefix, name, ".PKGINFO"):
				return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeArchLinuxPackage}
			case tarPathEqual(prefix, name, "Vagrantfile"):
				return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeVagrantBox}
			case tarPathEqual(prefix, name, "install/doinst.sh"):
				return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeSlackwarePackage}
			case tarPathEqual(prefix, name, "ComicInfo.xml"), tarPathEqual(prefix, name, "comicinfo.xml"):
				return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeComicBook}
			case tarPathEqual(prefix, name, "metadata"), tarPathEqual(prefix, name, "deploy"):
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

	var unsignedSum uint64
	var signedSum int64

	for i, value := range header {
		if i >= 148 && i < 156 {
			value = ' '
		}

		unsignedSum += uint64(value)
		signedSum += int64(int8(value))
	}

	return unsignedSum == stored || signedSum >= 0 && uint64(signedSum) == stored
}

func parseTarOctal(field []byte) (uint64, bool) {
	if len(field) > 0 && field[0]&0x80 != 0 {
		if field[0]&0x40 != 0 {
			return 0, false
		}

		value := uint64(field[0] & 0x3f)

		for _, part := range field[1:] {
			if value > (^uint64(0)-uint64(part))/256 {
				return 0, false
			}

			value = value*256 + uint64(part)
		}

		return value, true
	}

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

func tarPathEqual(prefix []byte, name []byte, target string) bool {
	if len(prefix) == 0 {
		return string(name) == target
	}

	return len(prefix)+1+len(name) == len(target) && string(prefix) == target[:len(prefix)] && target[len(prefix)] == '/' && string(name) == target[len(prefix)+1:]
}
