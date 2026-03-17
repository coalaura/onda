package custom

import (
	"bytes"

	"github.com/coalaura/onda/types"
)

func DetectTar(b types.Buffer) *types.Metadata {
	var isTar bool

	for offset := 0; offset+512 <= b.Len(); offset += 512 {
		if b.Len() >= offset+262 && (string(b[offset+257:offset+262]) == "ustar" || string(b[offset+257:offset+264]) == "ustar  ") {
			isTar = true

			nameEnd := bytes.IndexByte(b[offset:offset+100], 0)
			if nameEnd == -1 {
				nameEnd = 100
			}

			if nameEnd > 0 {
				name := string(b[offset : offset+nameEnd])

				switch name {
				case "package/package.json", "package.json":
					return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeNpmPackageTarball}
				case "oci-layout", "index.json", "manifest.json":
					return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeOCIImageLayoutTar}
				case "PKG-INFO", "setup.py", "pyproject.toml":
					return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypePythonSourceDistributionSDist}
				case "info/index.json":
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
