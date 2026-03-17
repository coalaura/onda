package custom

import (
	"bytes"

	"github.com/coalaura/onda/types"
)

func DetectTar(b types.Buffer) *types.Metadata {
	var isTar bool

	for offset := 0; offset+512 <= b.Len(); offset += 512 {
		if b.Has(offset+257, []byte("ustar\x00")) || b.Has(offset+257, []byte("ustar  ")) {
			isTar = true

			nameEnd := bytes.IndexByte(b[offset:offset+100], 0)
			if nameEnd == -1 {
				nameEnd = 100
			}

			if nameEnd > 0 {
				name := b[offset : offset+nameEnd]

				if bytes.Equal(name, []byte("package/package.json")) || bytes.Equal(name, []byte("package.json")) {
					return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeNpmPackageTarball}
				}

				if bytes.Equal(name, []byte("oci-layout")) || bytes.Equal(name, []byte("index.json")) || bytes.Equal(name, []byte("manifest.json")) {
					return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeOCIImageLayoutTar}
				}

				if bytes.Equal(name, []byte("PKG-INFO")) || bytes.Equal(name, []byte("setup.py")) || bytes.Equal(name, []byte("pyproject.toml")) {
					return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypePythonSourceDistributionSDist}
				}

				if bytes.Equal(name, []byte("info/index.json")) {
					return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeCondaPackage}
				}
			}
		}
	}

	if isTar {
		return &types.Metadata{Kind: types.KindTARArchive}
	}

	return nil
}
