package topic09_goroutines

/*
Title: EX02 — SafeCounter: atomic vs mutex under concurrent increments

Why this matters
- Data races corrupt state and produce heisenbugs. A correct counter under heavy concurrency is a basic building block.

Requirements
- Implement SafeCounter with:
  - Inc(n int): increments by n (n may be negative)
  - Load() int: returns current value
  - Reset(): sets value to 0
- The implementation must be race-free and safe under concurrent calls.
- Prefer minimal allocations; do not use channels.

Constraints and pitfalls
- Ensure memory visibility across goroutines.
- Avoid deadlocks; keep critical sections small.

Tricky edge case
- Negative increments should not underflow; store in int and allow negative totals.
*/

import "sync"

type SafeCounter struct {
	mu sync.Mutex
	v  int
}

func (c *SafeCounter) Inc(n int) { // TODO: implement (mutex or atomic)
}

func (c *SafeCounter) Load() int { // TODO
	return 0
}

func (c *SafeCounter) Reset() { // TODO
}
