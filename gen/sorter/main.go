package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strings"
)

var ignoredTypes = map[string]bool{
	// Structural / Generic
	"TypeNone":      true,
	"TypeEmpty":     true,
	"TypeContainer": true,
	"TypeWrapper":   true,
	"TypeLegacy":    true,
	"TypeSpecial":   true,

	// Endianness / Bitness
	"Type32Bit":                  true,
	"Type32BitBigEndian":         true,
	"Type32BitLittleEndian":      true,
	"Type64Bit":                  true,
	"Type64BitBigEndian":         true,
	"Type64BitLittleEndian":      true,
	"TypeBigEndian":              true,
	"TypeLittleEndian":           true,
	"TypeBinaryBigEndian":        true,
	"TypeBinaryLittleEndian":     true,
	"TypeNanosecondBigEndian":    true,
	"TypeNanosecondLittleEndian": true,
	"TypeByteSwapped":            true,

	// Versions
	"TypeVersion035": true,
	"TypeVersion036": true,
	"TypeVersion037": true,
	"TypeVersion038": true,
	"TypeVersion039": true,
	"TypeVersion040": true,
	"TypeVersion041": true,
	"TypeVersion1":   true,
	"TypeVersion2":   true,
	"TypeVersion3":   true,
	"TypeVersion4":   true,
	"TypeVersion5":   true,
	"TypeVersion7":   true,
	"TypeVersion8":   true,
	"TypeGIF87a":     true,
	"TypeGIF89a":     true,

	// Compression / Chunk States
	"TypeCodestream":     true,
	"TypeLZ4Legacy":      true,
	"TypeLZMACompressed": true,
	"TypeZlibCompressed": true,
	"TypeUncompressed":   true,
	"TypeSkinnableFrame": true,
	"TypeCAT":            true,
	"TypeLIST":           true,

	// Filesystem Entries
	"TypeBlockDevice":     true,
	"TypeCharacterDevice": true,
	"TypeDirectory":       true,
	"TypeNamedPipe":       true,
	"TypeSocket":          true,
	"TypeSymbolicLink":    true,

	// Specific Internal Markers / Architectures
	"TypeB4":              true,
	"TypeCE":              true,
	"TypeCOWD":            true,
	"TypeKDM":             true,
	"TypeIMMM":            true,
	"TypePE32ARM":         true,
	"TypePE32ARM64":       true,
	"TypePE32ARMv7":       true,
	"TypePE32Itanium":     true,
	"TypePE32PlusARM64":   true,
	"TypePE32PlusUnknown": true,
	"TypePE32PlusX8664":   true,
	"TypePE32Unknown":     true,
	"TypePE32X86":         true,
	"TypePE32X8664":       true,
	"TypeFAT12":           true,
	"TypeFAT16":           true,
	"TypeFAT32":           true,

	// Text / Encodings
	"TypeASCII":           true,
	"TypeUTF8":            true,
	"TypeUTF16":           true,
	"TypeUTF16LE":         true,
	"TypeUTF16BE":         true,
	"TypeUTF32LE":         true,
	"TypeUTF32BE":         true,
	"TypeLatin1":          true,
	"TypeWindows1252":     true,
	"TypeNewASCII":        true,
	"TypeNewASCIIWithCRC": true,
	"TypeOldASCII":        true,

	// Source Code / Scripts
	"TypeBashScript":                 true,
	"TypeBatchScript":                true,
	"TypeCMakeScript":                true,
	"TypeCPPSource":                  true,
	"TypeCSharpSource":               true,
	"TypeCSource":                    true,
	"TypeCSS":                        true,
	"TypeDartSource":                 true,
	"TypeDockerfile":                 true,
	"TypeDockerComposeConfiguration": true,
	"TypeGoSource":                   true,
	"TypeGraphQL":                    true,
	"TypeGraphQLSchema":              true,
	"TypeINIConfiguration":           true,
	"TypeJavaSource":                 true,
	"TypeJavaScriptSource":           true,
	"TypeJSONCSource":                true,
	"TypeJSON5Source":                true,
	"TypeKotlinSource":               true,
	"TypeLuaScript":                  true,
	"TypeMakefile":                   true,
	"TypeMarkdownDocument":           true,
	"TypeObjectiveCSource":           true,
	"TypePHPScript":                  true,
	"TypePerlScript":                 true,
	"TypePowerShellScript":           true,
	"TypePropertiesFile":             true,
	"TypePythonScript":               true,
	"TypeRScript":                    true,
	"TypeRubyScript":                 true,
	"TypeRustSource":                 true,
	"TypeScalaSource":                true,
	"TypeSQLScript":                  true,
	"TypeSwiftSource":                true,
	"TypeTOMLConfiguration":          true,
	"TypeTerraformConfiguration":     true,
	"TypeTerraformModule":            true,
	"TypeTypeScriptSource":           true,
	"TypeYAMLConfiguration":          true,
	"TypeElixirScript":               true,
	"TypeHaskellSource":              true,
	"TypeVimScript":                  true,
	"TypeFishScript":                 true,
	"TypeZshScript":                  true,
	"TypeNixExpression":              true,
	"TypeGuileScript":                true,
	"TypeEmacsLispScript":            true,
	"TypeClojureScript":              true,
	"TypeFSharpSource":               true,
	"TypeShellcheckDirective":        true,
	"TypeZigSource":                  true,
	"TypeSoliditySource":             true,
	"TypeNimSource":                  true,

	// Templates / Components
	"TypeVueComponent":       true,
	"TypeReactComponent":     true,
	"TypeAngularComponent":   true,
	"TypeSvelteComponent":    true,
	"TypeDjangoTemplate":     true,
	"TypeJinjaTemplate":      true,
	"TypeGoTemplate":         true,
	"TypeHandlebarsTemplate": true,

	// Interface Definitions
	"TypeProtocolBuffer":       true,
	"TypeThriftInterface":      true,
	"TypeOpenAPISpecification": true,

	// Hardware Description Languages
	"TypeVHDLSource":          true,
	"TypeVerilogSource":       true,
	"TypeSystemVerilogSource": true,

	// Proof Assistants / Theorem Provers
	"TypeCoqSource":   true,
	"TypeLeanSource":  true,
	"TypeIdrisSource": true,
	"TypeAgdaSource":  true,

	// Documentation Formats
	"TypeLaTeXDocument":    true,
	"TypeTeXDocument":      true,
	"TypeReStructuredText": true,
	"TypeAsciiDoc":         true,
	"TypeOrgMode":          true,

	// Data Formats
	"TypeCSVData":       true,
	"TypeTSVData":       true,
	"TypeLogData":       true,
	"TypeHTTPLog":       true,
	"TypeApacheLog":     true,
	"TypeNginxLog":      true,
	"TypeDiffPatch":     true,
	"TypeGitDiff":       true,
	"TypeCommitMessage": true,
	"TypeGeoJSON":       true,
	"TypeTopoJSON":      true,

	// Configuration Files
	"TypeSSHConfig": true,
	"TypeGitConfig": true,

	// High Entropy / Binary
	"TypeHighEntropy": true,
	"TypeEncrypted":   true,
	"TypeBinary":      true,
}

var removeEntries = []string{
	"KindCALSImage",
	"KindDolphinRVZ",
	"KindEFax",
	"KindElectronicArtsBigArchive",
	"KindExpertWitnessFormat",
	"KindFlowCytometryStandard",
	"KindGDSII",
	"KindHealthLevel7",
	"KindIdRoQVideo",
	"KindLuraDocument",
	"KindMacOSEventsLog",
	"KindMicrosoftInfo",
	"KindMicrosoftWinMobileNoteDocument",
	"KindNationalImageryTransmission",
	"KindNationalTransferFormat",
	"KindNintendoBCFNT",
	"KindNintendoBFLYT",
	"KindNintendoBRRES",
	"KindNintendoTPL",
	"KindOpenStreetMapPBF",
	"KindPacketSnifferXCP",
	"KindPathWayMap",
	"KindPowerISODAA",
	"KindQuickReport",
	"KindRagePackageFormat",
	"KindRealMedia",
	"KindRuntimeSoftwareImage",
	"KindTargetExpress",
	"KindUniversalDiskFormat",
	"KindVBScriptEncoded",
	"KindVCard",
	"KindValvePak",
	"KindVideoCD",
	"KindWalkmanMP3",
	"KindWebArchive",
	"KindWebVTTSubtitle",
	"KindWhereIsItCatalog",
	"KindWiiGameCubeTHP",
	"KindWinPharoahCapture",
	"KindWinPharoahFilter",
	"KindWindowsResource",
	"KindWindowsShortcut",
	"KindWwisePackage",
	"KindXPCOMLibraries",
	"KindZBackup",
	"TypeAC3",
	"TypeADTS",
	"TypeB4",
	"TypeBA2",
	"TypeBethesdaArchive2",
	"TypeBlockedGNUZipFormat",
	"TypeBrotherPEC",
	"TypeBrotherPES",
	"TypeCAT",
	"TypeCE",
	"TypeCOWD",
	"TypeDSAKey",
	"TypeEAC3",
	"TypeECKey",
	"TypeHDPhotoImage",
	"TypeHighEfficiency",
	"TypeIMMM",
	"TypeInternalWAD",
	"TypeKDB",
	"TypeKDBX",
	"TypeKDM",
	"TypeLIST",
	"TypeM2TSBDAV",
	"TypeOBJ",
	"TypePSB",
	"TypePSD",
	"TypePatchWAD",
	"TypePortableArbitraryMap",
	"TypeQCOW1",
	"TypeQCOW2",
	"TypeRPF0",
	"TypeRPF2",
	"TypeRPF3",
	"TypeRPF4",
	"TypeRPF6",
	"TypeRPF7",
	"TypeRSAKey",
	"TypeVMDK",
	"TypeVMDKDescription",
}

func main() {
	fmt.Println("Sorting ids_kind.go and ids_type.go...")
	count := sortIDs("./ids_kind.go", "./ids_type.go")

	fmt.Println("Sorting format registrations...")
	sortRegistrations("./internal")

	fmt.Println("Updating README.md...")
	updateReadme(count)
}

func updateReadme(count int) {
	path := "../README.md"

	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("failed to read README.md: %v\n", err)

		os.Exit(1)
	}

	rounded := (count / 10) * 10

	re := regexp.MustCompile(`over \*\*\d+\+ file formats\*\*`)
	newText := fmt.Sprintf("over **%d+ file formats**", rounded)

	updated := re.ReplaceAll(content, []byte(newText))

	if !bytes.Equal(content, updated) {
		if err := os.WriteFile(path, updated, 0644); err != nil {
			fmt.Printf("failed to write README.md: %v\n", err)

			os.Exit(1)
		} else {
			fmt.Printf("Updated README.md format count to %d+\n", rounded)
		}
	}
}

func sortIDs(kindPath, typePath string) int {
	kindsMap := make(map[string]string)
	typesMap := make(map[string]string)

	parseFile := func(path string) {
		fset := token.NewFileSet()

		node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			fmt.Printf("failed to parse %s: %v\n", path, err)

			os.Exit(1)
		}

		for _, decl := range node.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			if genDecl.Tok == token.CONST {
				for _, spec := range genDecl.Specs {
					vs, ok := spec.(*ast.ValueSpec)
					if !ok {
						continue
					}

					for _, name := range vs.Names {
						if strings.HasPrefix(name.Name, "Kind") || name.Name == "MaxKinds" {
							if _, exists := kindsMap[name.Name]; !exists {
								kindsMap[name.Name] = `""`
							}
						} else if strings.HasPrefix(name.Name, "Type") || name.Name == "MaxTypes" {
							if _, exists := typesMap[name.Name]; !exists {
								typesMap[name.Name] = `""`
							}
						}
					}
				}
			} else if genDecl.Tok == token.VAR {
				for _, spec := range genDecl.Specs {
					vs, ok := spec.(*ast.ValueSpec)
					if !ok {
						continue
					}

					for _, name := range vs.Names {
						if name.Name == "kindNames" || name.Name == "typeNames" {
							compLit, ok := vs.Values[0].(*ast.CompositeLit)
							if !ok {
								continue
							}

							for _, elt := range compLit.Elts {
								kv, ok := elt.(*ast.KeyValueExpr)
								if !ok {
									continue
								}

								keyIdent, ok := kv.Key.(*ast.Ident)
								if !ok {
									continue
								}

								valLit, ok := kv.Value.(*ast.BasicLit)
								if !ok {
									continue
								}

								if name.Name == "kindNames" {
									kindsMap[keyIdent.Name] = valLit.Value
								} else {
									typesMap[keyIdent.Name] = valLit.Value
								}
							}
						}
					}
				}
			}
		}
	}

	parseFile(kindPath)
	parseFile(typePath)

	var (
		kinds       []string
		hasMaxKinds bool
	)

	for k := range kindsMap {
		if k == "KindUnknown" || slices.Contains(removeEntries, k) {
			continue
		}

		if k == "MaxKinds" {
			hasMaxKinds = true

			continue
		}

		kinds = append(kinds, k)
	}

	sort.Slice(kinds, func(i, j int) bool {
		return strings.ToLower(kinds[i]) < strings.ToLower(kinds[j])
	})

	kinds = append([]string{"KindUnknown"}, kinds...)

	if hasMaxKinds {
		kinds = append(kinds, "MaxKinds")
	}

	var (
		typesList   []string
		hasMaxTypes bool
	)

	for t := range typesMap {
		if t == "TypeNone" || slices.Contains(removeEntries, t) {
			continue
		}

		if t == "MaxTypes" {
			hasMaxTypes = true

			continue
		}

		typesList = append(typesList, t)
	}

	sort.Slice(typesList, func(i, j int) bool {
		return strings.ToLower(typesList[i]) < strings.ToLower(typesList[j])
	})

	typesList = append([]string{"TypeNone"}, typesList...)

	if hasMaxTypes {
		typesList = append(typesList, "MaxTypes")
	}

	var kindBuf bytes.Buffer

	kindBuf.WriteString("package types\n\n")
	kindBuf.WriteString("type KindID uint16\n\n")
	kindBuf.WriteString("const (\n")

	for i, k := range kinds {
		if i == 0 {
			fmt.Fprintf(&kindBuf, "\t%s KindID = iota\n", k)
		} else {
			fmt.Fprintf(&kindBuf, "\t%s\n", k)
		}
	}

	kindBuf.WriteString(")\n\nvar kindNames = [...]string{\n")

	for _, k := range kinds {
		if k == "MaxKinds" {
			continue
		}

		fmt.Fprintf(&kindBuf, "\t%s: %s,\n", k, kindsMap[k])
	}

	kindBuf.WriteString("}\n\n")
	kindBuf.WriteString(`func (k KindID) String() string {
	if int(k) >= 0 && int(k) < len(kindNames) {
		name := kindNames[k]
		if name != "" {
			return name
		}
	}

	return "Unknown"
}
`)

	formattedKind, err := format.Source(kindBuf.Bytes())
	if err != nil {
		fmt.Printf("format error in %s: %v\n", kindPath, err)

		os.Exit(1)
	}

	if err := os.WriteFile(kindPath, formattedKind, 0644); err != nil {
		fmt.Printf("write error for %s: %v\n", kindPath, err)

		os.Exit(1)
	}

	var typeBuf bytes.Buffer

	typeBuf.WriteString("package types\n\n")
	typeBuf.WriteString("type TypeID uint16\n\n")
	typeBuf.WriteString("const (\n")

	for i, t := range typesList {
		if i == 0 {
			fmt.Fprintf(&typeBuf, "\t%s TypeID = iota\n", t)
		} else {
			fmt.Fprintf(&typeBuf, "\t%s\n", t)
		}
	}

	typeBuf.WriteString(")\n\nvar typeNames = [...]string{\n")

	for _, t := range typesList {
		if t == "MaxTypes" {
			continue
		}
		fmt.Fprintf(&typeBuf, "\t%s: %s,\n", t, typesMap[t])
	}

	typeBuf.WriteString("}\n\n")
	typeBuf.WriteString(`func (t TypeID) String() string {
	if int(t) >= 0 && int(t) < len(typeNames) {
		return typeNames[t]
	}

	return ""
}
`)

	formattedType, err := format.Source(typeBuf.Bytes())
	if err != nil {
		fmt.Printf("format error in %s: %v\n", typePath, err)

		os.Exit(1)
	}

	if err := os.WriteFile(typePath, formattedType, 0644); err != nil {
		fmt.Printf("write error for %s: %v\n", typePath, err)

		os.Exit(1)
	}

	var validTypes int

	for _, t := range typesList {
		if t == "MaxTypes" {
			continue
		}

		if !ignoredTypes[t] {
			validTypes++
		}
	}

	kindCount := len(kinds) - 1 // subtract KindUnknown
	if hasMaxKinds {
		kindCount-- // subtract MaxKinds
	}

	return kindCount + validTypes
}

func sortRegistrations(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("failed to read dir %s: %v\n", dir, err)

		os.Exit(1)
	}

	for _, entry := range files {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}

		path := filepath.Join(dir, entry.Name())

		processFormatFile(path)
	}
}

func processFormatFile(path string) {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("failed to parse %s: %v\n", path, err)

		os.Exit(1)
	}

	var initFunc *ast.FuncDecl

	for _, decl := range node.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok && fn.Name.Name == "init" {
			initFunc = fn

			break
		}
	}

	if initFunc == nil || len(initFunc.Body.List) == 0 {
		return
	}

	var (
		sigs           []ast.Stmt
		maskedSigs     []ast.Stmt
		weakSigs       []ast.Stmt
		weakMaskedSigs []ast.Stmt
		customSigs     []ast.Stmt
		weakSigsCustom []ast.Stmt
		fallbackSigs   []ast.Stmt
		otherSigs      []ast.Stmt
	)

	for _, stmt := range initFunc.Body.List {
		exprStmt, ok := stmt.(*ast.ExprStmt)
		if !ok {
			otherSigs = append(otherSigs, stmt)

			continue
		}

		callExpr, ok := exprStmt.X.(*ast.CallExpr)
		if !ok {
			otherSigs = append(otherSigs, stmt)

			continue
		}

		selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			otherSigs = append(otherSigs, stmt)

			continue
		}

		switch selExpr.Sel.Name {
		case "RegisterSignature":
			sigs = append(sigs, stmt)
		case "RegisterMaskedSignature":
			maskedSigs = append(maskedSigs, stmt)
		case "RegisterWeakSignature":
			weakSigs = append(weakSigs, stmt)
		case "RegisterWeakMaskedSignature":
			weakMaskedSigs = append(weakMaskedSigs, stmt)
		case "Register":
			customSigs = append(customSigs, stmt)
		case "RegisterWeak":
			weakSigsCustom = append(weakSigsCustom, stmt)
		case "RegisterFallback":
			fallbackSigs = append(fallbackSigs, stmt)
		default:
			otherSigs = append(otherSigs, stmt)
		}
	}

	sortStmts := func(stmts []ast.Stmt) {
		sort.SliceStable(stmts, func(i, j int) bool {
			strI := strings.ToLower(stmtToString(fset, stmts[i]))
			strJ := strings.ToLower(stmtToString(fset, stmts[j]))
			return strI < strJ
		})
	}

	sortStmts(sigs)
	sortStmts(maskedSigs)
	sortStmts(weakSigs)
	sortStmts(weakMaskedSigs)
	sortStmts(customSigs)
	sortStmts(weakSigsCustom)
	sortStmts(fallbackSigs)
	sortStmts(otherSigs)

	originalSrc, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("failed to read %s: %v\n", path, err)

		os.Exit(1)
	}

	start := fset.Position(initFunc.Body.Pos()).Offset
	end := fset.Position(initFunc.Body.End()).Offset

	var bodyBuf bytes.Buffer

	bodyBuf.WriteString("{\n")

	groups := [][]ast.Stmt{sigs, maskedSigs, weakSigs, weakMaskedSigs, customSigs, weakSigsCustom, fallbackSigs, otherSigs}

	firstGroup := true

	for _, group := range groups {
		if len(group) == 0 {
			continue
		}

		if !firstGroup {
			bodyBuf.WriteString("\n")
		}

		for _, s := range group {
			bodyBuf.WriteByte('\t')
			bodyBuf.WriteString(stmtToString(fset, s))
			bodyBuf.WriteByte('\n')
		}

		firstGroup = false
	}

	bodyBuf.WriteString("}")

	newSrc := append(originalSrc[:start], bodyBuf.Bytes()...)
	newSrc = append(newSrc, originalSrc[end+1:]...)

	finalSrc, err := format.Source(newSrc)
	if err != nil {
		fmt.Printf("failed final format for %s: %v\n", path, err)

		os.Exit(1)
	}

	if err := os.WriteFile(path, finalSrc, 0644); err != nil {
		fmt.Printf("failed to write %s: %v\n", path, err)

		os.Exit(1)
	}
}

func stmtToString(fset *token.FileSet, stmt ast.Stmt) string {
	var buf bytes.Buffer

	format.Node(&buf, fset, stmt)

	return buf.String()
}
