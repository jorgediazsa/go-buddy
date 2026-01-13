package topic08_packages_modules

/*
Title: EX10 — Go source filename validation

Why this matters
- Consistent filenames enable platform/arch-specific builds and tooling.

Requirements
- ValidateGoFilename(name string) error rules:
  - Non-empty; must end with ".go".
  - Only ASCII letters, digits, underscore, and dot.
  - No spaces.
  - Optional build tags suffixes like "_unix", "_amd64" are allowed but not required.
  - Name must start with a letter.

Constraints and pitfalls
- Do not parse build tags; just validate characters and extension.
- Hidden files (starting with dot) are invalid for package source (e.g., ".foo.go").

Tricky edge case
- Names like "a..go" (double dots) should be allowed, but "a .go" (space) is not.
*/

import "errors"

func ValidateGoFilename(name string) error { // TODO implement
	if name == "" {
		return errors.New("empty")
	}
	return errors.New("TODO")
}
