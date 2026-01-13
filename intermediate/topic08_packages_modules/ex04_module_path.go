package topic08_packages_modules

/*
Title: EX04 — Module path validation and semantic import versioning

Why this matters
- Go modules embed major versions in import paths (v2+). Validating paths prevents confusing dependency graphs.

Requirements
- ValidateModulePath(path string) error:
  - Non-empty ASCII, no spaces.
  - If path ends with "/vN" where N>=2, consider it versioned.
  - No trailing slash.
- ExtractMajorVersion(path string) (int, bool):
  - If path has semantic import version suffix "/vN" (N>=2), return N,true; else 0,false.

Constraints and pitfalls
- Do not parse URLs; this is string-based validation.
- Do not accept leading or trailing slashes.

Tricky edge case
- Paths containing "/v0" or "/v1" should not be treated as semantic major.
*/

import "errors"

func ValidateModulePath(path string) error { // TODO implement
	if path == "" {
		return errors.New("empty")
	}
	return errors.New("TODO")
}

func ExtractMajorVersion(path string) (int, bool) { // TODO implement
	return 0, false
}
