package main

import (
	"fmt"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"os"
	"strconv"
	"strings"
)

const customElementStruct = "github.com/myrjola/templ-htmx-tailwind-custom-elements/components.CustomElement"
const componentPackage = "github.com/myrjola/templ-htmx-tailwind-custom-elements/components"

type bundler struct {
	scripts strings.Builder
	styles  strings.Builder
}

func newBundler() *bundler {
	return &bundler{}
}

// ProcessPackage extracts Script and Style fields from CustomElement declarations for later bundling
func (b *bundler) ProcessPackage(pkg *packages.Package) {
	types := pkg.TypesInfo.Types
	for key := range types {
		var (
			decl  *ast.CompositeLit
			ident *ast.Ident
			ok    bool
		)

		// Skip everything except CustomElement declarations

		if decl, ok = key.(*ast.CompositeLit); !ok {
			continue
		}
		if ident, ok = decl.Type.(*ast.Ident); !ok {
			continue
		}
		if pkg.TypesInfo.TypeOf(ident).String() != customElementStruct {
			continue
		}

		// Extract Script and Style fields from CustomElement declarations

		for _, elt := range decl.Elts {
			if kv, ok := elt.(*ast.KeyValueExpr); ok {
				if ident, ok := kv.Key.(*ast.Ident); ok {
					var lit *ast.BasicLit

					// Only extract string literals
					if lit, ok = kv.Value.(*ast.BasicLit); !ok {
						continue
					}

					// Unquote the string literal
					val, err := strconv.Unquote(lit.Value)
					if err != nil {
						_, _ = fmt.Fprintf(os.Stderr, "unquote: %v\n", err)
						os.Exit(1)
					}

					if ident.Name == "Script" {
						b.scripts.WriteString(val)
					} else if ident.Name == "Style" {
						b.styles.WriteString(val)
					}
				}
			}
		}
	}
}

func (b *bundler) Bundle() error {
	// Minify scripts
	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)
	minifiedScripts, err := m.String("text/javascript", b.scripts.String())

	// Write script bundle
	scriptF, err := os.Create("./static/bundle.js")
	if err != nil {
		return fmt.Errorf("create bundle.js: %w", err)
	}
	defer scriptF.Close()
	_, err = scriptF.WriteString(minifiedScripts)
	if err != nil {
		return fmt.Errorf("write bundle.js: %w", err)
	}

	// Write style bundle
	styleF, err := os.Create("./bundle.css")
	if err != nil {
		return fmt.Errorf("create bundle.css: %w", err)
	}
	defer styleF.Close()
	_, err = styleF.WriteString(b.styles.String())
	if err != nil {
		return fmt.Errorf("write bundle.css: %w", err)
	}

	return nil
}

func main() {
	pkgs, err := packages.Load(&packages.Config{
		Mode:  packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
		Tests: false,
	}, componentPackage)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "load: %v\n", err)
		os.Exit(1)
	}
	if packages.PrintErrors(pkgs) > 0 {
		_, _ = fmt.Fprintf(os.Stderr, "print errors detected")
		os.Exit(1)
	}

	b := newBundler()

	for _, pkg := range pkgs {
		b.ProcessPackage(pkg)
	}

	err = b.Bundle()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "load: %v\n", err)
		os.Exit(1)
	}
}
