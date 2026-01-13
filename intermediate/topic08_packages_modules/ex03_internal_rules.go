package topic08_packages_modules

/*
Title: EX03 — Internal package import rules

Why this matters
- The internal package pattern prevents external packages from importing implementation details. Enforcing this avoids brittle couplings.

Requirements
- AllowedImport(from, to string) bool:
  - Returns false if to has an "internal" path segment and from is outside the parent of that internal tree.
  - Otherwise returns true.
  - Treat paths as slash-separated import paths (no filesystem probing).

Constraints and pitfalls
- Normalize duplicate slashes.
- An import "x/y/internal/z" is allowed only to importers whose path shares prefix "x/y" (the parent of internal).

Tricky edge case
- Nested internal segments (e.g., x/y/internal/a/internal/b) — the nearest internal’s parent controls access.
*/

// AllowedImport reports whether import from -> to is allowed under internal rules.
func AllowedImport(from, to string) bool { // TODO implement
	return false
}
