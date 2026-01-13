package topic09_goroutines

/*
Title: EX07 — ParallelMapInts: deterministic fan-out without channels

Why this matters
- Many workloads benefit from concurrency while preserving input order. This exercise practices WaitGroup + mutex/indexed writes.

Requirements
- ParallelMapInts(inputs []int, fn func(int) int, workers int) []int
  - Run up to `workers` goroutines to apply fn over inputs concurrently.
  - Preserve order: out[i] must be fn(inputs[i]).
  - workers <= 0 should be treated as 1; cap workers at len(inputs).
  - No channel-based API. Internal channels allowed but avoid teaching them.

Constraints and pitfalls
- Avoid data races on the output slice.
- Do not leak goroutines; all must finish before returning.

Tricky edge case
- Empty input should return empty output without starting workers.
*/

func ParallelMapInts(inputs []int, fn func(int) int, workers int) []int { // TODO implement
	// placeholder: sequential implementation to compile
	out := make([]int, len(inputs))
	for i, v := range inputs {
		if fn != nil {
			out[i] = fn(v)
		}
	}
	return out
}
