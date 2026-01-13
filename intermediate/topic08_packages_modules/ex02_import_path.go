package topic08_packages_modules

/*
Title: EX02 — Import path construction and internal path detection

Why this matters
- Consistent import paths and internal package fences prevent accidental dependencies and cycles.

Requirements
- BuildImportPath(modulePath, pkg string) (string, error):
  - modulePath and pkg must be non-empty ASCII without spaces.
  - Join as modulePath + "/" + pkg, collapsing duplicate slashes.
  - Reject relative elements ("./", "../").
- IsInternalPath(p string) bool: returns true if any path segment equals "internal".

Constraints and pitfalls
- Do not allow leading or trailing slashes in the result.
- No Unicode; ASCII-only validation.

Tricky edge case
- Paths like "example.com//mod////pkg" must normalize to "example.com/mod/pkg".
*/

import "errors"

// BuildImportPath builds a valid import path from modulePath and pkg.
func BuildImportPath(modulePath, pkg string) (string, error) { // TODO implement
	if modulePath == "" || pkg == "" {
		return "", errors.New("empty")
	}
	return "", errors.New("TODO")
}

// IsInternalPath reports whether the path contains an "internal" segment.
func IsInternalPath(p string) bool { // TODO implement
	return false
}
