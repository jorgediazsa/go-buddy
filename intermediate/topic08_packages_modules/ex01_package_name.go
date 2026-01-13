package topic08_packages_modules

/*
Title: EX01 — Package naming and public surface sanity

Why this matters
- Consistent, valid package names improve discoverability and avoid subtle tooling issues.

Requirements
- Implement SanitizePackageName(s string) (string, error):
  - Trim spaces.
  - Allow only ASCII letters, digits, and underscores.
  - Must start with a letter.
  - Convert to lower-case.
  - Enforce length 1..63.
- Implement IsExportedName(name string) bool for symbol names (first rune uppercase A–Z).

Constraints and pitfalls
- No panics; return clear errors for invalid inputs.
- ASCII-only check; do not use Unicode categories here.

Tricky edge case
- Names like "_abc" (invalid start), "" (empty), and extremely long names should be rejected.
*/

import (
	"errors"
	"unicode"
)

// SanitizePackageName normalizes and validates a Go package name.
func SanitizePackageName(s string) (string, error) { // TODO: implement
	if s == "" {
		return "", errors.New("empty")
	}
	return "", errors.New("TODO: implement SanitizePackageName")
}

// IsExportedName reports whether name is exported (starts with uppercase A–Z).
func IsExportedName(name string) bool { // TODO: implement
	if name == "" {
		return false
	}
	r, _ := utf8DecodeRuneInString(name)
	return unicode.IsUpper(r)
}

// Minimal helper to avoid importing utf8; we only need first rune.
func utf8DecodeRuneInString(s string) (r rune, size int) {
	for i := range s {
		return rune(s[i]), 1
	}
	return 0, 0
}
