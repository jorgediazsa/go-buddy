package beginner

/*
Exercise EX03 — Slices: copies, in-place transforms, and insertion

Why this matters
- Slice headers are small values, but they alias shared backing arrays — easy to mutate unintentionally.
- In-place algorithms reduce allocations and improve performance.

Requirements
- CloneInts: deep-copy an []int (nil → nil, empty → empty distinct backing array).
- RotateLeftInPlace: rotate elements left by k (k≥0) in-place. k may exceed len; use k%len.
- SortedInsert: given an ascending slice, return a new slice with v inserted while preserving order (stable among equals).

Constraints and pitfalls
- Do not use generics; only []int.
- For RotateLeftInPlace, do nothing for len<2 or k==0. Negative k is invalid.
- For SortedInsert, the input must not be mutated; return a new slice.

Tricky edge cases
- k==0, k==len, k>len, len==0/1.
- SortedInsert at head, middle (between equals), and tail.
*/

import "errors"

// CloneInts returns a deep copy of src (nil→nil, empty→empty with distinct backing).
func CloneInts(src []int) []int {
	// TODO: implement
	return nil
}

// RotateLeftInPlace rotates the slice left by k steps in-place.
func RotateLeftInPlace(nums []int, k int) error { // TODO: implement
	if k < 0 {
		return errors.New("k must be non-negative")
	}
	// TODO: implement real rotation
	return errors.New("not implemented")
}

// SortedInsert returns a new slice with v inserted into ascending src.
func SortedInsert(src []int, v int) []int { // TODO: implement
	// TODO: binary search and allocate new slice
	out := make([]int, len(src))
	copy(out, src)
	return out
}
