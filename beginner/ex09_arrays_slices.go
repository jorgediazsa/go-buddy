package beginner

/*
Exercise EX09 — Arrays vs Slices: aliasing and deep copies

Why this matters
- Arrays are values; slices are descriptors of a shared backing array. Confusing them causes aliasing bugs.

Requirements
- NewIdentity(n int) ([][]int, error): build an n×n identity matrix with no shared rows (mutating one row must not affect others).
- DeepCopy2D(src [][]int) [][]int: deep-copy a 2D slice; preserve nil vs empty semantics.
- SumArray3(a [3]int) int: return the sum to emphasize array value semantics.

Constraints and pitfalls
- Validate 1 <= n <= 64 for NewIdentity; return error otherwise.
- Preallocate efficiently; avoid quadratic appends.
- DeepCopy2D must allocate new inner slices and copy contents.

Tricky edge cases
- NewIdentity with n==1.
- Deep copy of nil slice, empty slice, and ragged matrices.
- SumArray3 on different arrays with same contents.
*/

import "errors"

// NewIdentity returns an identity matrix of size n with independent rows.
func NewIdentity(n int) ([][]int, error) { // TODO: implement
	if n < 1 || n > 64 {
		return nil, errors.New("n out of range")
	}
	return nil, errors.New("not implemented")
}

// DeepCopy2D returns a deep copy of a 2D slice.
func DeepCopy2D(src [][]int) [][]int { // TODO: implement
	if src == nil {
		return nil
	}
	out := make([][]int, len(src))
	// TODO: copy each row deeply
	return out
}

// SumArray3 returns the sum of a 3-element array.
func SumArray3(a [3]int) int { // TODO: implement
	return 0
}
