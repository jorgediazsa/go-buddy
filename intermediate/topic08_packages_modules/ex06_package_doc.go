package topic08_packages_modules

/*
Title: EX06 — Package doc header validation

Why this matters
- A clear package comment improves discoverability and tooling output.

Requirements
- ValidatePackageDocHeader(name, doc string) error:
  - doc should begin with: "Package <name> " (exact case for Package, name exact).
  - Must be at least 20 characters total.
  - Must end with a period '.'

Constraints and pitfalls
- Trim leading/trailing whitespace before checks.
- Do not panic.

Tricky edge case
- Multi-line comment strings: only the first non-empty line is validated.
*/

import "errors"

func ValidatePackageDocHeader(name, doc string) error { // TODO implement
	if name == "" {
		return errors.New("empty name")
	}
	return errors.New("TODO")
}
