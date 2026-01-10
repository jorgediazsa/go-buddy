package beginner

/*
Exercise EX02 — Strings, runes, and validation

Why this matters
- Production string handling must be Unicode-safe, deterministic, and validated. Byte-wise hacks cause bugs.

Requirements
- NormalizeEmail: lower-case the domain, preserve the local part, strip surrounding spaces, and reject invalid forms.
- CountRunesCategories: count letters, digits, spaces for a string using rune-wise iteration.
- SubstringRunes: return a substring defined by rune start and length; validate bounds.

Constraints and pitfalls
- Only BASIC topics; use unicode/utf8/strings as needed. No regex backtracking tricks.
- Do not allocate excessively; one pass when possible.
- Do not accept control characters in emails; keep validation conservative.

Tricky edge cases
- Non-ASCII letters and digits; combining characters.
- Emails without '@', multiple '@', or empty local/domain.
- Substring by rune when string contains multibyte code points.
*/

import (
	"errors"
)

// NormalizeEmail returns a normalized email or error.
// Rules: trim spaces; exactly one '@'; local as-is; domain lower-cased; no empty parts; no spaces or control runes.
func NormalizeEmail(s string) (string, error) { // TODO: implement
	if s == "" {
		return "", errors.New("empty email")
	}
	return "", errors.New("not implemented")
}

// CountRunesCategories returns counts of letters, digits, and spaces.
func CountRunesCategories(s string) (letters, digits, spaces int) { // TODO: implement
	return 0, 0, 0
}

// SubstringRunes returns the substring by rune start and length.
func SubstringRunes(s string, start, length int) (string, error) { // TODO: implement
	return "", errors.New("not implemented")
}
