package topic08_packages_modules

/*
Title: EX05 — Detect forbidden imports by parsing Go source

Why this matters
- Enforcing dependency boundaries prevents architectural erosion. Static checks via the Go parser are robust and fast.

Requirements
- NoForbiddenImports(src string, forbidden []string) ([]string, error):
  - Parse the provided Go source (single file content) and collect import paths.
  - Return a sorted unique slice of forbidden imports found (intersection) without duplicates.
  - If parsing fails, return an error.

Constraints and pitfalls
- Use only go/parser and go/ast from the standard library; do not regex imports.
- Treat import paths as raw string literals and unquote them when needed.

Tricky edge case
- Mixed import forms: single and grouped import blocks, alias imports, and blank identifier imports.
*/

import (
	"go/parser"
	"go/token"
	"sort"
	"strconv"
)

// NoForbiddenImports returns the list of forbidden import paths present in src.
func NoForbiddenImports(src string, forbidden []string) ([]string, error) { // TODO implement
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "ex.go", src, parser.ImportsOnly)
	if err != nil {
		return nil, err
	}
	forbid := map[string]struct{}{}
	for _, p := range forbidden {
		forbid[p] = struct{}{}
	}
	seen := map[string]struct{}{}
	var out []string
	for _, imp := range file.Imports {
		if imp.Path == nil {
			continue
		}
		val := imp.Path.Value
		// unquote
		s, _ := strconv.Unquote(val)
		if _, ok := forbid[s]; ok {
			if _, dup := seen[s]; !dup {
				seen[s] = struct{}{}
				out = append(out, s)
			}
		}
		_ = imp // avoid unused if build tags prune
	}
	sort.Strings(out)
	return out, nil
}
