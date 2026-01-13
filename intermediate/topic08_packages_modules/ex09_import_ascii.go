package topic08_packages_modules

/*
Title: EX09 — Import path ASCII and structure validation

Why this matters
- Enforcing well-formed import paths avoids surprises across platforms and tooling.

Requirements
- ValidateImportASCII(p string) error:
  - Non-empty ASCII; no spaces or control characters.
  - No leading or trailing slash.
  - No "//" segments; collapse is not allowed here.
  - No "." or ".." path elements.

Constraints and pitfalls
- Do not normalize; reject invalid input as-is.
- ASCII only; disallow non-ASCII bytes.

Tricky edge case
- A single segment "." or ".." by itself must be rejected.
*/

import "errors"

func ValidateImportASCII(p string) error { // TODO implement
	if p == "" {
		return errors.New("empty")
	}
	return errors.New("TODO")
}
