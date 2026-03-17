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
}

func main() {
	log.Println("Parsing signatures for Radix Trie optimization...")
	sigs := parseSignatures("./internal")

	log.Println("Generating optimized jump table...")
	generateOptimizedCode(sigs, "./optimized.go")
}

func parseSignatures(dir string) []Sig {
	var sigs []Sig

	fset := token.NewFileSet()

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("failed to read dir: %v", err)
	}

	for _, entry := range files {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}

		path := filepath.Join(dir, entry.Name())

		node, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			log.Fatalf("failed to parse %s: %v", path, err)
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

				isMask := sel.Sel.Name == "RegisterMaskedSignature"
				isSig := sel.Sel.Name == "RegisterSignature"

				if !isSig && !isMask {
					continue
				}

				kind := strings.TrimPrefix(extractSelector(call.Args[0]), "types.")
				typ := strings.TrimPrefix(extractSelector(call.Args[1]), "types.")
				offset := extractInt(call.Args[2])
				magic := extractBytes(call.Args[3])

				var mask []byte

				if isMask {
					mask = extractBytes(call.Args[4])
				}

				sigs = append(sigs, Sig{
					Kind:   kind,
					Type:   typ,
					Offset: offset,
					Magic:  magic,
					Mask:   mask,
					IsMask: isMask,
				})
			}
		}
	}

	return sigs
}

func generateOptimizedCode(sigs []Sig, outPath string) {
	byOffset := make(map[int][]Sig)

	for _, s := range sigs {
		byOffset[s.Offset] = append(byOffset[s.Offset], s)
	}

	var offsets []int

	for off := range byOffset {
		offsets = append(offsets, off)
	}

	sort.Ints(offsets)

	var buf bytes.Buffer

	buf.WriteString("// code generated, don't edit\n")
	buf.WriteString("package types\n\n")
	buf.WriteString("func detectOptimized(b Buffer) *Metadata {\n")
	buf.WriteString("\tif b.Len() == 0 {\n\t\treturn nil\n\t}\n\n")

	for _, off := range offsets {
		group := byOffset[off]

		generateRadixNode(group, 0, off, &buf, "\t")

		buf.WriteString("\n")
	}

	buf.WriteString("\treturn nil\n")
	buf.WriteString("}\n")

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("format error: %v\n%s", err, buf.String())
	}

	if err := os.WriteFile(outPath, formatted, 0644); err != nil {
		log.Fatalf("write error: %v", err)
	}
}

func generateRadixNode(sigs []Sig, depth int, offset int, buf *bytes.Buffer, indent string) {
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

		sort.Slice(unswitchable, func(i, j int) bool {
			return len(unswitchable[i].Magic) > len(unswitchable[j].Magic)
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

			if len(keys) == 1 && len(byByte[byte(keys[0])]) == 1 {
				emitIfHas(byByte[byte(keys[0])][0], buf, indent)
			} else {
				fmt.Fprintf(buf, "%sif b.Len() > %d {\n", indent, offset+depth)

				indent += "\t"

				fmt.Fprintf(buf, "%s_ = b[%d] // BCE hint\n", indent, offset+depth)
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

func emitIfHas(c Sig, buf *bytes.Buffer, indent string) {
	if c.IsMask {
		fmt.Fprintf(buf, "%sif b.HasMask(%d, %#v, %#v) {\n", indent, c.Offset, c.Magic, c.Mask)
	} else {
		fmt.Fprintf(buf, "%sif b.Has(%d, %#v) {\n", indent, c.Offset, c.Magic)
	}

	if c.Type == "TypeNone" || c.Type == "" {
		fmt.Fprintf(buf, "%s\treturn &Metadata{Kind: %s}\n", indent, c.Kind)
	} else {
		fmt.Fprintf(buf, "%s\treturn &Metadata{Kind: %s, Type: %s}\n", indent, c.Kind, c.Type)
	}

	fmt.Fprintf(buf, "%s}\n", indent)
}

func extractSelector(expr ast.Expr) string {
	if sel, ok := expr.(*ast.SelectorExpr); ok {
		if x, ok := sel.X.(*ast.Ident); ok {
			return x.Name + "." + sel.Sel.Name
		}
	}

	if id, ok := expr.(*ast.Ident); ok {
		return id.Name
	}

	return ""
}

func extractInt(expr ast.Expr) int {
	if bl, ok := expr.(*ast.BasicLit); ok && bl.Kind == token.INT {
		v, _ := strconv.ParseInt(bl.Value, 0, 64)

		return int(v)
	}

	return 0
}

func extractBytes(expr ast.Expr) []byte {
	if call, ok := expr.(*ast.CallExpr); ok {
		if len(call.Args) == 1 {
			if bl, ok := call.Args[0].(*ast.BasicLit); ok && bl.Kind == token.STRING {
				s, _ := strconv.Unquote(bl.Value)

				return []byte(s)
			}
		}
	}

	if comp, ok := expr.(*ast.CompositeLit); ok {
		var b []byte

		for _, elt := range comp.Elts {
			if bl, ok := elt.(*ast.BasicLit); ok {
				switch bl.Kind {
				case token.INT:
					v, _ := strconv.ParseInt(bl.Value, 0, 64)

					b = append(b, byte(v))
				case token.CHAR:
					s, _ := strconv.Unquote(bl.Value)

					b = append(b, s[0])
				}
			}
		}

		return b
	}

	return nil
}
