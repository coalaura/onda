package custom

import "github.com/coalaura/onda/types"

func DetectTar(b types.Buffer) *types.Metadata {
	if b.Has(257, []byte("ustar\x00")) || b.Has(257, []byte("ustar  ")) {
		if b.Has(0, []byte("package/package.json")) || b.Has(0, []byte("package.json")) {
			return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeNpmPackageTarball}
		}

		if b.Has(0, []byte("oci-layout")) || b.Has(0, []byte("index.json")) || b.Has(0, []byte("manifest.json")) {
			return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeOCIImageLayoutTar}
		}

		if b.Has(0, []byte("PKG-INFO")) || b.Has(0, []byte("setup.py")) || b.Has(0, []byte("pyproject.toml")) {
			return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypePythonSourceDistributionSDist}
		}

		if b.Has(0, []byte("info/index.json")) {
			return &types.Metadata{Kind: types.KindTARArchive, Type: types.TypeCondaPackage}
		}

		return &types.Metadata{
			Kind: types.KindTARArchive,
		}
	}

	return nil
}
