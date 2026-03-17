package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func main() {
	log.Println("Sorting ids.go...")
	count := sortIDs("./ids.go")

	log.Println("Sorting format registrations...")
	sortRegistrations("./internal")

	log.Println("Updating README.md...")
	updateReadme(count)
}

func updateReadme(count int) {
	path := "../README.md"

	content, err := os.ReadFile(path)
	if err != nil {
		log.Printf("failed to read README.md: %v", err)
		return
	}

	rounded := (count / 10) * 10

	re := regexp.MustCompile(`over \*\*\d+\+ file formats\*\*`)
	newText := fmt.Sprintf("over **%d+ file formats**", rounded)

	updated := re.ReplaceAll(content, []byte(newText))

	if !bytes.Equal(content, updated) {
		if err := os.WriteFile(path, updated, 0644); err != nil {
			log.Printf("failed to write README.md: %v", err)
		} else {
			log.Printf("Updated README.md format count to %d+", rounded)
		}
	}
}

func sortIDs(path string) int {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("failed to parse %s: %v", path, err)
	}

	kindsMap := make(map[string]string)
	typesMap := make(map[string]string)

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
					if strings.HasPrefix(name.Name, "Kind") {
						if _, exists := kindsMap[name.Name]; !exists {
							kindsMap[name.Name] = `""`
						}
					} else if strings.HasPrefix(name.Name, "Type") {
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

	var kinds []string

	for k := range kindsMap {
		if k != "KindUnknown" {
			kinds = append(kinds, k)
		}
	}

	sort.Slice(kinds, func(i, j int) bool {
		return strings.ToLower(kinds[i]) < strings.ToLower(kinds[j])
	})

	kinds = append([]string{"KindUnknown"}, kinds...)

	var typesList []string

	for t := range typesMap {
		if t != "TypeNone" {
			typesList = append(typesList, t)
		}
	}

	sort.Slice(typesList, func(i, j int) bool {
		return strings.ToLower(typesList[i]) < strings.ToLower(typesList[j])
	})

	typesList = append([]string{"TypeNone"}, typesList...)

	var buf bytes.Buffer

	buf.WriteString("package types\n\n")
	buf.WriteString("type KindID uint16\n")
	buf.WriteString("type TypeID uint16\n\n")

	buf.WriteString("const (\n")

	for i, k := range kinds {
		if i == 0 {
			fmt.Fprintf(&buf, "\t%s KindID = iota\n", k)
		} else {
			fmt.Fprintf(&buf, "\t%s\n", k)
		}
	}

	buf.WriteString(")\n\nconst (\n")

	for i, t := range typesList {
		if i == 0 {
			fmt.Fprintf(&buf, "\t%s TypeID = iota\n", t)
		} else {
			fmt.Fprintf(&buf, "\t%s\n", t)
		}
	}

	buf.WriteString(")\n\nvar kindNames = [...]string{\n")

	for _, k := range kinds {
		fmt.Fprintf(&buf, "\t%s: %s,\n", k, kindsMap[k])
	}

	buf.WriteString("}\n\nvar typeNames = [...]string{\n")

	for _, t := range typesList {
		fmt.Fprintf(&buf, "\t%s: %s,\n", t, typesMap[t])
	}

	buf.WriteString("}\n\n")

	buf.WriteString(`func (k KindID) String() string {
	if int(k) >= 0 && int(k) < len(kindNames) {
		name := kindNames[k]
		if name != "" {
			return name
		}
	}

	return "Unknown"
}

func (t TypeID) String() string {
	if int(t) >= 0 && int(t) < len(typeNames) {
		return typeNames[t]
	}

	return ""
}
`)

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("format error in ids.go: %v", err)
	}

	if err := os.WriteFile(path, formatted, 0644); err != nil {
		log.Fatalf("write error for ids.go: %v", err)
	}

	ignoredTypes := map[string]bool{
		"TypeNone":                   true,
		"TypeEmpty":                  true,
		"TypeContainer":              true,
		"TypeWrapper":                true,
		"TypeSpanned":                true,
		"TypeLittleEndian":           true,
		"TypeBigEndian":              true,
		"Type32BitBigEndian":         true,
		"Type32BitLittleEndian":      true,
		"Type64Bit":                  true,
		"Type64BitBigEndian":         true,
		"Type64BitLittleEndian":      true,
		"TypeBinaryBigEndian":        true,
		"TypeBinaryLittleEndian":     true,
		"TypeNanosecondBigEndian":    true,
		"TypeNanosecondLittleEndian": true,
		"TypeByteSwapped":            true,
		"TypeCodestream":             true,
		"TypeStreamVersion7":         true,
		"TypeStreamVersion8":         true,
	}

	var validTypes int

	for _, t := range typesList {
		if !ignoredTypes[t] {
			validTypes++
		}
	}

	// len(kinds) includes KindUnknown, so we subtract 1
	return (len(kinds) - 1) + validTypes
}

func sortRegistrations(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("failed to read dir %s: %v", dir, err)
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
		log.Fatalf("failed to parse %s: %v", path, err)
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
		sigs         []ast.Stmt
		maskedSigs   []ast.Stmt
		customSigs   []ast.Stmt
		weakSigs     []ast.Stmt
		fallbackSigs []ast.Stmt
		otherSigs    []ast.Stmt
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
		case "Register":
			customSigs = append(customSigs, stmt)
		case "RegisterWeak":
			weakSigs = append(weakSigs, stmt)
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
	sortStmts(customSigs)
	sortStmts(weakSigs)
	sortStmts(fallbackSigs)
	sortStmts(otherSigs)

	originalSrc, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	start := fset.Position(initFunc.Body.Pos()).Offset
	end := fset.Position(initFunc.Body.End()).Offset

	var bodyBuf bytes.Buffer

	bodyBuf.WriteString("{\n")

	groups := [][]ast.Stmt{sigs, maskedSigs, customSigs, weakSigs, fallbackSigs, otherSigs}
	firstGroup := true

	for _, group := range groups {
		if len(group) == 0 {
			continue
		}

		if !firstGroup {
			bodyBuf.WriteString("\n")
		}

		for _, s := range group {
			bodyBuf.WriteString("\t" + stmtToString(fset, s) + "\n")
		}

		firstGroup = false
	}

	bodyBuf.WriteString("}")

	newSrc := append(originalSrc[:start], bodyBuf.Bytes()...)
	newSrc = append(newSrc, originalSrc[end+1:]...)

	finalSrc, err := format.Source(newSrc)
	if err != nil {
		fmt.Println(string(newSrc))
		log.Fatalf("failed final format for %s: %v", path, err)
	}

	if err := os.WriteFile(path, finalSrc, 0644); err != nil {
		log.Fatalf("failed to write %s: %v", path, err)
	}
}

func stmtToString(fset *token.FileSet, stmt ast.Stmt) string {
	var buf bytes.Buffer

	format.Node(&buf, fset, stmt)

	return buf.String()
}
