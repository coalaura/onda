//go:build ignore

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
	"sort"
	"strings"
)

func main() {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, "ids.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("failed to parse ids.go: %v", err)
	}

	kindsMap := make(map[string]string)
	typesMap := make(map[string]string)

	// Extract existing constants and string mappings from the AST
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

	// Sort kinds alphabetically, keeping KindUnknown at the top
	var kinds []string

	for k := range kindsMap {
		if k != "KindUnknown" {
			kinds = append(kinds, k)
		}
	}

	sort.Strings(kinds)

	kinds = append([]string{"KindUnknown"}, kinds...)

	// Sort types alphabetically, keeping TypeNone at the top
	var typesList []string

	for t := range typesMap {
		if t != "TypeNone" {
			typesList = append(typesList, t)
		}
	}

	sort.Strings(typesList)

	typesList = append([]string{"TypeNone"}, typesList...)

	// Generate the new ids.go file
	var buf bytes.Buffer

	buf.WriteString(`//go:generate go run sort.go

package types

type KindID uint16
type TypeID uint16

const (
`)

	for i, k := range kinds {
		if i == 0 {
			buf.WriteString(fmt.Sprintf("\t%s KindID = iota\n", k))
		} else {
			buf.WriteString(fmt.Sprintf("\t%s\n", k))
		}
	}

	buf.WriteString(")\n\nconst (\n")

	for i, t := range typesList {
		if i == 0 {
			buf.WriteString(fmt.Sprintf("\t%s TypeID = iota\n", t))
		} else {
			buf.WriteString(fmt.Sprintf("\t%s\n", t))
		}
	}

	buf.WriteString(")\n\nvar kindNames = [...]string{\n")

	for _, k := range kinds {
		buf.WriteString(fmt.Sprintf("\t%s: %s,\n", k, kindsMap[k]))
	}

	buf.WriteString("}\n\nvar typeNames = [...]string{\n")

	for _, t := range typesList {
		buf.WriteString(fmt.Sprintf("\t%s: %s,\n", t, typesMap[t]))
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

	// Format the generated source code
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("format error: %v\n%s", err, buf.String())
	}

	// Overwrite ids.go
	if err := os.WriteFile("ids.go", formatted, 0644); err != nil {
		log.Fatalf("write error: %v", err)
	}
}
