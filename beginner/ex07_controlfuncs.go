package beginner

/*
Exercise EX07 — Control flow and higher-order functions

Why this matters
- Business logic often combines branching and iteration with configurable predicates.
- Writing tight, allocation-aware loops is critical for performance.

Requirements
- FilterInPlace(nums []int, keep func(int) bool) int: keep elements for which keep(x)==true, in-place; return new length.
- RangeSums(n int) (evens int, odds int): sum of 0..n inclusive; validate n>=0.
- CategorizeTemps(ts []int) map[string]int: counts of "cold" (<10), "mild" (10..24), "hot" (>=25).

Constraints and pitfalls
- FilterInPlace must not allocate a new slice; mutate input slice content and return new logical length.
- RangeSums must validate input and avoid overflow for reasonable n (use int64 internally then cast?). Keep it int.
- CategorizeTemps must not modify input and should return non-nil map.

Tricky edge cases
- FilterInPlace with all removed or all kept; empty slice.
- RangeSums with n==0 and large n; negative n should error via negative marker return? Use named returns? Keep it simple: return zeros on negative and signal via bool? Instead, return (int,int) and negative n should yield both zeros.
- Temperatures on category boundaries.
*/

// FilterInPlace keeps elements satisfying keep and returns the new length.
func FilterInPlace(nums []int, keep func(int) bool) int { // TODO: implement
	// Two-pointer compaction
	return 0
}

// RangeSums returns sums of even and odd numbers from 0..n inclusive. If n<0, both sums must be 0.
func RangeSums(n int) (evens int, odds int) { // TODO: implement
	return 0, 0
}

// CategorizeTemps returns counts for cold(<10), mild(10..24), hot(>=25).
func CategorizeTemps(ts []int) map[string]int { // TODO: implement
	return map[string]int{"cold": 0, "mild": 0, "hot": 0}
}
