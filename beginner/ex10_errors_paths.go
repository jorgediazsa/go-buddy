package beginner

/*
Exercise EX10 — Errors and simple path utilities

Why this matters
- Robust programs validate inputs and return clear errors. Path-like normalization is a frequent need even without filesystem access.

Requirements
- NormalizePathSimple(p string) (string, error)
  - Treat path segments separated by '/'
  - Collapse '.' segments; resolve '..' by removing the previous segment
  - Preserve leading '/' (absolute vs relative)
  - Reject inputs containing "//" (double slash), or attempts to go above root for absolute paths
  - Return "" error on empty input
- JoinClean(base, rel string) (string, error)
  - rel must be relative (no leading '/')
  - Join as base + "/" + rel and normalize with NormalizePathSimple
  - If base is empty, error
- SplitExt(name string) (base string, ext string)
  - Split extension at the last '.'; return ext including the dot
  - Filenames like ".bashrc" have no extension (base is ".bashrc", ext is "")

Constraints and pitfalls
- Do not use filepath.Clean; implement explicitly to practice control flow and errors.
- Keep behavior deterministic; do not interact with the OS.

Tricky edge cases
- ".." on an absolute root ("/..") → error; on relative path may produce leading ".." segments and should be allowed.
- Trailing slash semantics: "a/b/" → treat like "a/b" (no trailing slash in output).
- Names with multiple dots for SplitExt: "archive.tar.gz" → base "archive.tar", ext ".gz".
*/

import "errors"

func NormalizePathSimple(p string) (string, error) { // TODO: implement
	if p == "" {
		return "", errors.New("empty path")
	}
	return "", errors.New("not implemented")
}

func JoinClean(base, rel string) (string, error) { // TODO: implement
	if base == "" {
		return "", errors.New("empty base")
	}
	return "", errors.New("not implemented")
}

func SplitExt(name string) (base string, ext string) { // TODO: implement
	return name, ""
}
