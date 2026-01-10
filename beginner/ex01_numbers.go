package beginner

/*
Exercise EX01 — Numeric correctness and conversions

Why this matters
- Numeric boundary handling is a common source of production bugs (overflows, truncation, invalid input).
- Understanding Go’s explicit conversions and integer sizes prevents subtle portability issues.

Requirements
- Implement strict decimal parsing to uint8 without accepting signs/spaces.
- Implement safe int32 addition with overflow/underflow detection.
- Implement percentage calculation using int64 with input validation.

Constraints and pitfalls
- Do not use big.Int or generics. Only BASIC topics.
- You may use the standard library (strconv, errors, fmt) but keep logic explicit.
- Be careful with intermediate widening when checking for overflows.
- Do not panic; return errors for invalid inputs.

Tricky edge cases
- "" (empty), leading zeros, non-digits, and values > 255 for ParseUint8Strict.
- AddInt32 near math.MinInt32/MaxInt32 boundaries.
- PercentOf for total==0 and ranges outside [0..total].
*/

import (
	"errors"
)

// ParseUint8Strict parses a base-10 string into uint8.
// Only digits 0-9 are allowed (no spaces, signs, or underscores).
func ParseUint8Strict(s string) (uint8, error) { // TODO: implement
	// TODO: implement strict parsing with range checking (0..255)
	if s == "" {
		return 0, errors.New("empty input")
	}
	return 0, errors.New("not implemented")
}

// AddInt32 returns a+b with overflow/underflow detection.
func AddInt32(a, b int32) (int32, error) { // TODO: implement
	// TODO: detect overflow/underflow and return an error
	return 0, errors.New("not implemented")
}

// PercentOf returns floor(part*100/total) using int64 arithmetic.
// Validates 0<=part<=total and total>0.
func PercentOf(part, total int64) (int64, error) { // TODO: implement
	// TODO: implement with validation and no division by zero
	return 0, errors.New("not implemented")
}
