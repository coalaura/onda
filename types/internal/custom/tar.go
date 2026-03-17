package custom

import "github.com/coalaura/onda/types"

func DetectTar(b types.Buffer) *types.Metadata {
	var names []string

	for offset := 0; offset+512 <= b.Len(); offset += 512 {
		if b.Has(offset+257, []byte("ustar\x00")) || b.Has(offset+257, []byte("ustar  ")) {
			var nameLen int

			for nameLen < 100 && b[offset+nameLen] != 0 {
				nameLen++
			}

			if nameLen > 0 {
				names = append(names, string(b[offset:offset+nameLen]))
			}
		}
	}

	if len(names) == 0 {
		return nil
	}

	for _, name := range names {
		if name == "package/package.json" || name == "package.json" {
			return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeNpmPackageTarball}
		}

		if name == "oci-layout" || name == "index.json" || name == "manifest.json" {
			return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeOCIImageLayoutTar}
		}

		if name == "PKG-INFO" || name == "setup.py" || name == "pyproject.toml" {
			return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypePythonSourceDistributionSDist}
		}

		if name == "info/index.json" {
			return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeCondaPackage}
		}
	}

	return &types.Metadata{Kind: types.KindTARArchive}
}
