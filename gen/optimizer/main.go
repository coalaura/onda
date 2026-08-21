package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/scanner"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Sig struct {
	Kind   string
	Type   string
	Offset int
	Magic  []byte
	Mask   []byte
	IsMask bool
	IsWeak bool
}

type Pair struct {
	Kind string
	Type string
}

type SignatureKey struct {
	Offset int
	Magic  string
	Mask   string
}

type Generator struct {
	StrongSigs []Sig
	WeakSigs   []Sig
	Custom     []string
	Weak       []string
	Fallback   []string
}

type ScannedToken struct {
	Offset int
	Token  token.Token
	Text   string
}

func main() {
	fmt.Println("Parsing signatures and detectors for Radix Trie optimization...")
	gen := parseSignatures("./internal")

	fmt.Println("Checking for duplicate signatures...")
	checkDuplicateSignatures(gen.StrongSigs, gen.WeakSigs)

	fmt.Println("Checking for unused Kind/Type IDs...")
	checkUnusedIDs("./ids_kind.go", "./ids_type.go", "./internal", "./custom")

	fmt.Println("Generating optimized jump table (gen_signatures.go)...")
	generateOptimizedCode(gen, "./gen_signatures.go")

	fmt.Println("Merging custom detectors (gen_detectors.go)...")
	mergeCustomDetectors("./custom", "./gen_detectors.go")

	fmt.Println("Generating format list (gen_formats.go)...")
	generateFormatList(gen, "./gen_formats.go")
}

func parseSignatures(dir string) *Generator {
	gen := &Generator{}

	fset := token.NewFileSet()

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("failed to read dir: %v\n", err)

		os.Exit(1)
	}

	for _, entry := range files {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}

		path := filepath.Join(dir, entry.Name())

		node, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			fmt.Printf("failed to parse %s: %v\n", path, err)

			os.Exit(1)
		}

		for _, decl := range node.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Name.Name != "init" {
				continue
			}

			for _, stmt := range fn.Body.List {
				exprStmt, ok := stmt.(*ast.ExprStmt)
				if !ok {
					continue
				}

				call, ok := exprStmt.X.(*ast.CallExpr)
				if !ok {
					continue
				}

				sel, ok := call.Fun.(*ast.SelectorExpr)
				if !ok {
					continue
				}

				pkg, ok := sel.X.(*ast.Ident)
				if !ok || pkg.Name != "types" {
					continue
				}

				funcName := sel.Sel.Name
				pos := fset.Position(call.Pos())

				switch funcName {
				case "RegisterSignature", "RegisterMaskedSignature", "RegisterWeakSignature", "RegisterWeakMaskedSignature":
					isMask := funcName == "RegisterMaskedSignature" || funcName == "RegisterWeakMaskedSignature"

					expectedArgs := 4
					if isMask {
						expectedArgs = 5
					}

					if len(call.Args) != expectedArgs {
						failRegistration(pos, fmt.Sprintf("%s requires %d arguments", funcName, expectedArgs))
					}

					kindSelector, ok := extractSelector(call.Args[0])
					if !ok || !strings.HasPrefix(kindSelector, "types.Kind") {
						failRegistration(pos, "invalid Kind identifier")
					}

					typeSelector, ok := extractSelector(call.Args[1])
					if !ok || !strings.HasPrefix(typeSelector, "types.Type") {
						failRegistration(pos, "invalid Type identifier")
					}

					offset, ok := extractInt(call.Args[2])
					if !ok || offset < 0 {
						failRegistration(pos, "offset must be a non-negative integer literal")
					}

					magic, ok := extractBytes(call.Args[3])
					if !ok || len(magic) == 0 {
						failRegistration(pos, "magic must be a non-empty byte literal")
					}

					var mask []byte

					if isMask {
						mask, ok = extractBytes(call.Args[4])
						if !ok || len(mask) != len(magic) {
							failRegistration(pos, "mask must be a byte literal with the same length as magic")
						}
					}

					sig := Sig{
						Kind:   strings.TrimPrefix(kindSelector, "types."),
						Type:   strings.TrimPrefix(typeSelector, "types."),
						Offset: offset,
						Magic:  magic,
						Mask:   mask,
						IsMask: isMask,
						IsWeak: funcName == "RegisterWeakSignature" || funcName == "RegisterWeakMaskedSignature",
					}

					if sig.IsWeak {
						gen.WeakSigs = append(gen.WeakSigs, sig)
					} else {
						gen.StrongSigs = append(gen.StrongSigs, sig)
					}
				case "Register", "RegisterWeak", "RegisterFallback":
					if len(call.Args) != 1 {
						failRegistration(pos, fmt.Sprintf("%s requires 1 argument", funcName))
					}

					argCall, ok := call.Args[0].(*ast.CallExpr)
					if !ok || len(argCall.Args) != 1 {
						failRegistration(pos, "detector must be wrapped in types.DetectFunc")
					}

					wrapper, ok := extractSelector(argCall.Fun)
					if !ok || wrapper != "types.DetectFunc" {
						failRegistration(pos, "detector must be wrapped in types.DetectFunc")
					}

					detSel, ok := argCall.Args[0].(*ast.SelectorExpr)
					if !ok {
						failRegistration(pos, "invalid detector function")
					}

					detPkg, ok := detSel.X.(*ast.Ident)
					if !ok || detPkg.Name != "custom" {
						failRegistration(pos, "detector function must be from the custom package")
					}

					funcNameStr := detSel.Sel.Name

					switch funcName {
					case "Register":
						gen.Custom = append(gen.Custom, funcNameStr)
					case "RegisterWeak":
						gen.Weak = append(gen.Weak, funcNameStr)
					case "RegisterFallback":
						gen.Fallback = append(gen.Fallback, funcNameStr)
					}
				}
			}
		}
	}

	return gen
}

func checkDuplicateSignatures(strongSigs, weakSigs []Sig) {
	seen := make(map[SignatureKey]Sig)
	hasDuplicates := false

	check := func(sigs []Sig, category string) {
		for _, sig := range sigs {
			mask := sig.Mask
			if len(mask) == 0 {
				mask = bytes.Repeat([]byte{0xff}, len(sig.Magic))
			}

			magic := make([]byte, len(sig.Magic))

			for i := range sig.Magic {
				magic[i] = sig.Magic[i] & mask[i]
			}

			key := SignatureKey{
				Offset: sig.Offset,
				Magic:  string(magic),
				Mask:   string(mask),
			}

			if prev, exists := seen[key]; exists {
				fmt.Printf("[%s] Conflict at Offset: %d | Magic: %q | Mask: %q\n"+
					"  -> Existing:  %s.%s\n"+
					"  -> Duplicate: %s.%s\n",
					category, sig.Offset, string(sig.Magic), string(sig.Mask),
					prev.Kind, prev.Type,
					sig.Kind, sig.Type,
				)

				hasDuplicates = true
			} else {
				seen[key] = sig
			}
		}
	}

	check(strongSigs, "strong signatures")
	check(weakSigs, "weak signatures")

	if hasDuplicates {
		fmt.Println("Found duplicate signatures (they evaluate to the exact same bytes and will fight each other). Please fix them.")

		os.Exit(1)
	}
}

func checkUnusedIDs(idsKindPath string, idsTypePath string, internalDir string, customDir string) {
	definedKinds := parseDefinedIDs(idsKindPath, "Kind")
	definedTypes := parseDefinedIDs(idsTypePath, "Type")

	usedKinds := make(map[string]bool)
	usedTypes := make(map[string]bool)

	collectUsedIDs := func(dir string) {
		fset := token.NewFileSet()

		files, err := os.ReadDir(dir)
		if err != nil {
			return
		}

		for _, entry := range files {
			if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
				continue
			}

			path := filepath.Join(dir, entry.Name())

			node, err := parser.ParseFile(fset, path, nil, 0)
			if err != nil {
				continue
			}

			ast.Inspect(node, func(n ast.Node) bool {
				if sel, ok := n.(*ast.SelectorExpr); ok {
					if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "types" {
						name := sel.Sel.Name
						if strings.HasPrefix(name, "Kind") {
							usedKinds[name] = true
						} else if strings.HasPrefix(name, "Type") {
							usedTypes[name] = true
						}
					}
				}

				return true
			})
		}
	}

	collectUsedIDs(internalDir)
	collectUsedIDs(customDir)

	collectUsedIDs(".")
	collectUsedIDs("..")

	var (
		unusedKinds []string
		unusedTypes []string
	)

	for k := range definedKinds {
		if k == "KindUnknown" {
			continue
		}

		if !usedKinds[k] {
			unusedKinds = append(unusedKinds, k)
		}
	}

	for t := range definedTypes {
		if t == "TypeNone" {
			continue
		}

		if !usedTypes[t] {
			unusedTypes = append(unusedTypes, t)
		}
	}

	sort.Strings(unusedKinds)
	sort.Strings(unusedTypes)

	if len(unusedKinds) > 0 {
		fmt.Println("UNUSED KIND IDs:")

		for _, k := range unusedKinds {
			fmt.Printf("  %s\n", k)
		}
	}

	if len(unusedTypes) > 0 {
		fmt.Println("UNUSED TYPE IDs:")

		for _, t := range unusedTypes {
			fmt.Printf("  %s\n", t)
		}
	}

	if len(unusedKinds) > 0 || len(unusedTypes) > 0 {
		fmt.Printf("Found %d unused Kind IDs and %d unused Type IDs\n", len(unusedKinds), len(unusedTypes))

		os.Exit(1)
	}
}

func parseDefinedIDs(path string, prefix string) map[string]bool {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		fmt.Printf("failed to parse %s: %v", path, err)
	}

	result := make(map[string]bool)

	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.CONST {
			continue
		}

		for _, spec := range genDecl.Specs {
			vs, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			for _, name := range vs.Names {
				if strings.HasPrefix(name.Name, prefix) {
					result[name.Name] = true
				}
			}
		}
	}

	return result
}

func generateOptimizedCode(gen *Generator, outPath string) {
	var buf strings.Builder

	buf.WriteString("// Code generated by optimizer, DO NOT EDIT.\n")
	buf.WriteString("package types\n\n")

	buf.WriteString("func detectOptimized(b Buffer) *Metadata {\n")
	buf.WriteString("\tif len(b) == 0 {\n\t\treturn nil\n\t}\n\n")

	// 1. Strong Custom Detectors (Containers, complex formats)
	sort.Strings(gen.Custom)

	for _, c := range gen.Custom {
		fmt.Fprintf(&buf, "\tif meta := %s(b); meta != nil {\n\t\treturn meta\n\t}\n\n", c)
	}

	// 2. Strong Signatures Radix Tree
	generateRadixForest(gen.StrongSigs, &buf)

	// 3. Weak Custom Detectors (JSON, XML, etc.)
	sort.Strings(gen.Weak)
	for _, c := range gen.Weak {
		fmt.Fprintf(&buf, "\tif meta := %s(b); meta != nil {\n\t\treturn meta\n\t}\n\n", c)
	}

	// 4. Weak Signatures Radix Tree
	generateRadixForest(gen.WeakSigs, &buf)

	// 5. Fallback Detectors (Plain Text)
	sort.Strings(gen.Fallback)

	for _, c := range gen.Fallback {
		fmt.Fprintf(&buf, "\tif meta := %s(b); meta != nil {\n\t\treturn meta\n\t}\n\n", c)
	}

	buf.WriteString("\treturn nil\n")
	buf.WriteString("}\n")

	formatted, err := format.Source([]byte(buf.String()))
	if err != nil {
		fmt.Printf("format error: %v\n%s\n", err, buf.String())

		os.Exit(1)
	}

	if err := os.WriteFile(outPath, formatted, 0644); err != nil {
		fmt.Printf("write error: %v\n", err)

		os.Exit(1)
	}
}

func mergeCustomDetectors(dir string, outPath string) {
	fset := token.NewFileSet()

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("failed to read custom dir: %v\n", err)

		os.Exit(1)
	}

	importSet := make(map[string]bool)

	var decls []string

	for _, entry := range files {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}

		path := filepath.Join(dir, entry.Name())

		node, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			fmt.Printf("failed to parse %s: %v\n", path, err)

			os.Exit(1)
		}

		for _, imp := range node.Imports {
			p := imp.Path.Value
			if p != `"github.com/coalaura/wtf/types"` && p != `"github.com/coalaura/wtf/types/internal/custom"` {
				var buf bytes.Buffer

				format.Node(&buf, fset, imp)

				importSet[buf.String()] = true
			}
		}

		for _, decl := range node.Decls {
			if genDecl, isGen := decl.(*ast.GenDecl); isGen {
				if genDecl.Tok == token.IMPORT {
					continue
				}

				var buf bytes.Buffer

				format.Node(&buf, fset, decl)

				decls = append(decls, buf.String())
			} else if fn, isFn := decl.(*ast.FuncDecl); isFn {
				if fn.Name.Name == "init" {
					continue
				}

				var buf bytes.Buffer

				format.Node(&buf, fset, decl)

				decls = append(decls, buf.String())
			}
		}
	}

	var out strings.Builder

	out.WriteString("// Code generated by optimizer, DO NOT EDIT.\n")
	out.WriteString("package types\n\n")

	if len(importSet) > 0 {
		out.WriteString("import (\n")

		var imps []string

		for imp := range importSet {
			imps = append(imps, imp)
		}

		sort.Strings(imps)

		for _, imp := range imps {
			out.WriteByte('\t')
			out.WriteString(imp)
			out.WriteByte('\n')
		}

		out.WriteString(")\n\n")
	}

	for _, d := range decls {
		d = stripPackageSelectors(d)

		out.WriteString(d)
		out.WriteString("\n\n")
	}

	formatted, err := format.Source([]byte(out.String()))
	if err != nil {
		fmt.Printf("format error in detectors: %v\n%s\n", err, out.String())

		os.Exit(1)
	}

	if err := os.WriteFile(outPath, formatted, 0644); err != nil {
		fmt.Printf("write error: %v\n", err)

		os.Exit(1)
	}
}

func stripPackageSelectors(source string) string {
	fset := token.NewFileSet()

	file := fset.AddFile("", -1, len(source))

	var scan scanner.Scanner

	scan.Init(file, []byte(source), nil, scanner.ScanComments)

	var tokens []ScannedToken

	for {
		position, tok, text := scan.Scan()
		if tok == token.EOF {
			break
		}

		tokens = append(tokens, ScannedToken{Offset: file.Offset(position), Token: tok, Text: text})
	}

	remove := make(map[int]int)

	for i := 0; i+2 < len(tokens); i++ {
		current := tokens[i]
		if current.Token != token.IDENT || current.Text != "types" && current.Text != "custom" || tokens[i+1].Token != token.PERIOD || tokens[i+2].Token != token.IDENT || tokens[i+2].Text[0] < 'A' || tokens[i+2].Text[0] > 'Z' {
			continue
		}

		remove[current.Offset] = tokens[i+1].Offset + 1
	}

	var out strings.Builder

	for offset := 0; offset < len(source); {
		if end, ok := remove[offset]; ok {
			offset = end
			continue
		}

		out.WriteByte(source[offset])

		offset++
	}

	return out.String()
}

func generateFormatList(gen *Generator, outPath string) {
	sigTypes := make(map[string][]string)

	addSigType := func(kind, typ string) {
		if typ == "" || typ == "TypeNone" {
			return
		}

		sigTypes[kind] = append(sigTypes[kind], typ)
	}

	for _, sig := range gen.StrongSigs {
		addSigType(sig.Kind, sig.Type)
	}

	for _, sig := range gen.WeakSigs {
		addSigType(sig.Kind, sig.Type)
	}

	customTypes := extractTypesFromCustom("./custom")

	for kind, types := range customTypes {
		sigTypes[kind] = append(sigTypes[kind], types...)
	}

	for kind := range sigTypes {
		seen := make(map[string]bool)

		var unique []string

		for _, t := range sigTypes[kind] {
			if !seen[t] {
				seen[t] = true

				unique = append(unique, t)
			}
		}

		sort.Strings(unique)

		sigTypes[kind] = unique
	}

	var buf strings.Builder

	buf.WriteString("// Code generated by optimizer, DO NOT EDIT.\n")
	buf.WriteString("package types\n\n")

	buf.WriteString("var subtypeMap = map[KindID][]TypeID{\n")

	var kinds []string

	for k := range sigTypes {
		kinds = append(kinds, k)
	}

	sort.Strings(kinds)

	for _, kind := range kinds {
		fmt.Fprintf(&buf, "\t%s: {", kind)

		for i, t := range sigTypes[kind] {
			if i > 0 {
				buf.WriteString(", ")
			}

			buf.WriteString(t)
		}

		buf.WriteString("},\n")
	}

	buf.WriteString("}\n\n")

	buf.WriteString("type FormatInfo struct {\n")
	buf.WriteString("\tKind string\n")
	buf.WriteString("\tType string\n")
	buf.WriteString("}\n\n")

	buf.WriteString("func ListFormats() []FormatInfo {\n")
	buf.WriteString("\tlist := make([]FormatInfo, 0, 256)\n\n")

	buf.WriteString("\tfor i := range KindID(len(kindNames)) {\n")
	buf.WriteString("\t\tname := kindNames[i]\n")
	buf.WriteString("\t\tif name == \"\" || name == \"Unknown\" {\n")
	buf.WriteString("\t\t\tcontinue\n")
	buf.WriteString("\t\t}\n\n")

	buf.WriteString("\t\tlist = append(list, FormatInfo{Kind: name})\n\n")

	buf.WriteString("\t\tif types, ok := subtypeMap[i]; ok {\n")
	buf.WriteString("\t\t\tfor _, t := range types {\n")
	buf.WriteString("\t\t\t\tif tn := t.String(); tn != \"\" {\n")
	buf.WriteString("\t\t\t\t\tlist = append(list, FormatInfo{Kind: name, Type: tn})\n")
	buf.WriteString("\t\t\t\t}\n")
	buf.WriteString("\t\t\t}\n")
	buf.WriteString("\t\t}\n")
	buf.WriteString("\t}\n\n")

	buf.WriteString("\treturn list\n")
	buf.WriteString("}\n")

	formatted, err := format.Source([]byte(buf.String()))
	if err != nil {
		fmt.Printf("format error: %v\n%s\n", err, buf.String())

		os.Exit(1)
	}

	if err := os.WriteFile(outPath, formatted, 0644); err != nil {
		fmt.Printf("write error: %v\n", err)

		os.Exit(1)
	}
}

func extractTypesFromCustom(dir string) map[string][]string {
	result := make(map[string][]string)

	files, err := os.ReadDir(dir)
	if err != nil {
		return result
	}

	kindRe := regexp.MustCompile(`Kind:\s*types\.(Kind\w+)`)
	typeRe := regexp.MustCompile(`Type:\s*types\.(Type\w+)`)
	dynamicTypeRe := regexp.MustCompile(`Type:\s*[a-z]\w*`)
	typeIDRe := regexp.MustCompile(`types\.(Type\w+)`)

	for _, entry := range files {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}

		path := filepath.Join(dir, entry.Name())

		content, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		dynamicKinds := make(map[string]bool)

		var i int

		for i < len(content) {
			idx := bytes.Index(content[i:], []byte("&types.Metadata{"))
			if idx == -1 {
				break
			}

			start := i + idx + len("&types.Metadata{")
			i = start

			depth := 1
			end := start

			for end < len(content) && depth > 0 {
				switch content[end] {
				case '{':
					depth++
				case '}':
					depth--
				}

				end++
			}

			block := content[start:end]

			kindMatch := kindRe.FindSubmatch(block)
			typeMatch := typeRe.FindSubmatch(block)

			if len(kindMatch) > 1 && len(typeMatch) > 1 {
				kind := string(kindMatch[1])
				typ := string(typeMatch[1])

				result[kind] = append(result[kind], typ)
			} else if len(kindMatch) > 1 && dynamicTypeRe.Match(block) {
				dynamicKinds[string(kindMatch[1])] = true
			}
		}

		for kind := range dynamicKinds {
			for _, match := range typeIDRe.FindAllSubmatch(content, -1) {
				typ := string(match[1])
				if typ != "TypeID" && typ != "TypeNone" {
					result[kind] = append(result[kind], typ)
				}
			}
		}
	}

	return result
}

func generateRadixForest(sigs []Sig, buf *strings.Builder) {
	byOffset := make(map[int][]Sig)

	for _, s := range sigs {
		byOffset[s.Offset] = append(byOffset[s.Offset], s)
	}

	var offsets []int

	for off := range byOffset {
		offsets = append(offsets, off)
	}

	sort.Ints(offsets)

	for _, off := range offsets {
		group := byOffset[off]

		generateRadixNode(group, 0, off, buf, "\t")

		buf.WriteString("\n")
	}
}

func generateRadixNode(sigs []Sig, depth int, offset int, buf *strings.Builder, indent string) {
	var (
		exact     []Sig
		remaining []Sig
	)

	for _, s := range sigs {
		if len(s.Magic) == depth {
			exact = append(exact, s)
		} else {
			remaining = append(remaining, s)
		}
	}

	if len(remaining) > 0 {
		var unswitchable []Sig

		byByte := make(map[byte][]Sig)

		for _, s := range remaining {
			if s.IsMask && s.Mask[depth] != 0xff {
				unswitchable = append(unswitchable, s)
			} else {
				byByte[s.Magic[depth]] = append(byByte[s.Magic[depth]], s)
			}
		}

		sort.SliceStable(unswitchable, func(i, j int) bool {
			if len(unswitchable[i].Magic) != len(unswitchable[j].Magic) {
				return len(unswitchable[i].Magic) > len(unswitchable[j].Magic)
			}

			if unswitchable[i].Kind != unswitchable[j].Kind {
				return unswitchable[i].Kind < unswitchable[j].Kind
			}

			return unswitchable[i].Type < unswitchable[j].Type
		})

		for _, s := range unswitchable {
			emitIfHas(s, buf, indent)
		}

		if len(byByte) > 0 {
			var keys []int

			for k := range byByte {
				keys = append(keys, int(k))
			}

			sort.Ints(keys)

			if len(keys) == 1 {
				bVal := byte(keys[0])
				group := byByte[bVal]

				if len(group) == 1 {
					emitIfHas(group[0], buf, indent)
				} else {
					fmt.Fprintf(buf, "%sif len(b) > %d && b[%d] == %#02x {\n", indent, offset+depth, offset+depth, bVal)

					l := lcp(group, depth+1)

					generateRadixNode(group, l, offset, buf, indent+"\t")

					fmt.Fprintf(buf, "%s}\n", indent)
				}
			} else {
				fmt.Fprintf(buf, "%sif len(b) > %d {\n", indent, offset+depth)
				indent += "\t"
				fmt.Fprintf(buf, "%sswitch b[%d] {\n", indent, offset+depth)

				for _, k := range keys {
					bVal := byte(k)
					group := byByte[bVal]

					fmt.Fprintf(buf, "%scase %#02x:\n", indent, bVal)

					if len(group) == 1 {
						emitIfHas(group[0], buf, indent+"\t")
					} else {
						l := lcp(group, depth+1)
						generateRadixNode(group, l, offset, buf, indent+"\t")
					}
				}

				fmt.Fprintf(buf, "%s}\n", indent)
				indent = indent[:len(indent)-1]
				fmt.Fprintf(buf, "%s}\n", indent)
			}
		}
	}

	sort.SliceStable(exact, func(i, j int) bool {
		if exact[i].Kind != exact[j].Kind {
			return exact[i].Kind < exact[j].Kind
		}

		return exact[i].Type < exact[j].Type
	})

	for _, s := range exact {
		emitIfHas(s, buf, indent)
	}
}

func lcp(sigs []Sig, startDepth int) int {
	if len(sigs) == 0 {
		return startDepth
	}

	minLen := len(sigs[0].Magic)

	for _, s := range sigs {
		if len(s.Magic) < minLen {
			minLen = len(s.Magic)
		}
	}

	l := startDepth

	for ; l < minLen; l++ {
		b := sigs[0].Magic[l]

		for _, s := range sigs {
			if s.IsMask && s.Mask[l] != 0xff {
				return l
			}

			if s.Magic[l] != b {
				return l
			}
		}
	}

	return l
}

func emitIfHas(c Sig, buf *strings.Builder, indent string) {
	if c.IsMask {
		fmt.Fprintf(buf, "%sif b.HasMask(%d, %q, %q) {\n", indent, c.Offset, string(c.Magic), string(c.Mask))
	} else {
		end := c.Offset + len(c.Magic)

		if len(c.Magic) == 1 {
			fmt.Fprintf(buf, "%sif len(b) >= %d && b[%d] == %#02x {\n", indent, end, c.Offset, c.Magic[0])
		} else if c.Offset == 0 {
			fmt.Fprintf(buf, "%sif len(b) >= %d && string(b[:%d]) == %q {\n", indent, end, end, string(c.Magic))
		} else {
			fmt.Fprintf(buf, "%sif len(b) >= %d && string(b[%d:%d]) == %q {\n", indent, end, c.Offset, end, string(c.Magic))
		}
	}

	var fields []string

	fields = append(fields, "Kind: "+c.Kind)

	if c.Type != "TypeNone" && c.Type != "" {
		fields = append(fields, "Type: "+c.Type)
	}

	if c.IsWeak {
		fields = append(fields, "Confidence: ConfidenceMedium")
	}

	fmt.Fprintf(buf, "%s\treturn &Metadata{%s}\n", indent, strings.Join(fields, ", "))
	fmt.Fprintf(buf, "%s}\n", indent)
}

func failRegistration(pos token.Position, message string) {
	fmt.Printf("%s: %s\n", pos, message)

	os.Exit(1)
}

func extractSelector(expr ast.Expr) (string, bool) {
	if sel, ok := expr.(*ast.SelectorExpr); ok {
		if x, ok := sel.X.(*ast.Ident); ok {
			return x.Name + "." + sel.Sel.Name, true
		}
	}

	if id, ok := expr.(*ast.Ident); ok {
		return id.Name, true
	}

	return "", false
}

func extractInt(expr ast.Expr) (int, bool) {
	if bl, ok := expr.(*ast.BasicLit); ok && bl.Kind == token.INT {
		v, err := strconv.ParseInt(bl.Value, 0, 64)
		if err != nil || int64(int(v)) != v {
			return 0, false
		}

		return int(v), true
	}

	return 0, false
}

func extractBytes(expr ast.Expr) ([]byte, bool) {
	if call, ok := expr.(*ast.CallExpr); ok {
		arrayType, isArray := call.Fun.(*ast.ArrayType)
		if isArray && arrayType.Len == nil {
			byteType, isByte := arrayType.Elt.(*ast.Ident)
			if isByte && byteType.Name == "byte" && len(call.Args) == 1 {
				if bl, ok := call.Args[0].(*ast.BasicLit); ok && bl.Kind == token.STRING {
					s, err := strconv.Unquote(bl.Value)
					if err == nil {
						return []byte(s), true
					}
				}
			}
		}
	}

	if comp, ok := expr.(*ast.CompositeLit); ok {
		arrayType, isArray := comp.Type.(*ast.ArrayType)
		if !isArray || arrayType.Len != nil {
			return nil, false
		}

		byteType, isByte := arrayType.Elt.(*ast.Ident)
		if !isByte || byteType.Name != "byte" {
			return nil, false
		}

		var b []byte

		for _, elt := range comp.Elts {
			bl, ok := elt.(*ast.BasicLit)
			if !ok {
				return nil, false
			}

			switch bl.Kind {
			case token.INT:
				v, err := strconv.ParseUint(bl.Value, 0, 8)
				if err != nil {
					return nil, false
				}

				b = append(b, byte(v))
			case token.CHAR:
				s, err := strconv.Unquote(bl.Value)
				if err != nil || len(s) != 1 {
					return nil, false
				}

				b = append(b, s[0])
			default:
				return nil, false
			}
		}

		return b, true
	}

	return nil, false
}
