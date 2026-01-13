package topic08_packages_modules

/*
Title: EX07 — List exported identifiers from a Go source file

Why this matters
- Tools and linters often need to inspect the public surface of packages. Parsing is more reliable than regex.

Requirements
- ListExportedIdents(src string) ([]string, error):
  - Parse the provided single-file Go source and collect exported identifiers:
    * top-level const, var, func, and type names
  - Return them sorted and unique.
  - On parse error, return it.

Constraints and pitfalls
- Only consider top-level declarations in the file; skip methods’ receiver types and unexported names.
- Do not panic.

Tricky edge case
- Mixed grouped declarations and multiple names in const/var specs.
*/

import (
	"go/parser"
	"go/token"
)

// ListExportedIdents returns sorted exported identifiers declared at top-level.
func ListExportedIdents(src string) ([]string, error) { // TODO implement
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "ex.go", src, 0)
	if err != nil {
		return nil, err
	}
	_ = file
	return nil, nil
}
