package topic12_error_handling

/*
Title: EX01 — Wrapping with %w and sentinel errors

Why this matters
- Correct error wrapping preserves root causes while adding context, enabling callers to inspect with errors.Is.

Requirements
- Define a sentinel error: ErrInvalid.
- Implement ParseUint32Strict(s string) (uint32, error):
  - Reject empty, signed, spaced, or non-digit inputs.
  - Parse base-10 into uint32; if overflow, return wrapped ErrInvalid.
  - On any invalid input, return fmt.Errorf("parse: %q: %w", s, ErrInvalid).
  - On success, return the value and nil.

Constraints and pitfalls
- Do not panic. Use only standard library.
- Avoid strconv.Parse* shortcuts that accept signs/underscores; implement validation explicitly.

Tricky edge case
- "00000000000" (many zeros) is valid; very long numeric strings over 10 chars may overflow uint32.
*/

import (
	"errors"
)

// ErrInvalid indicates the input is invalid.
var ErrInvalid = errors.New("invalid input")

// ParseUint32Strict parses a strictly decimal unsigned 32-bit integer.
func ParseUint32Strict(s string) (uint32, error) { // TODO: implement
	if s == "" {
		return 0, wrapInvalid(s)
	}
	// TODO implement
	return 0, wrapInvalid(s)
}

func wrapInvalid(s string) error {
	// keep fmt out to avoid import until implemented by user
	return ErrInvalid
}
