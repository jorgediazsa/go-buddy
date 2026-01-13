package topic09_goroutines

/*
Title: EX08 — ThrottledScheduler: schedule tasks with latency bounds

Why this matters
- Concurrency improves responsiveness even without parallel CPUs. Throttling avoids overload while maintaining interleaving.

Requirements
- Implement ThrottledScheduler with:
  - NewThrottledScheduler(maxConcurrent int) *ThrottledScheduler
  - Submit(fn func()) bool // returns false if closed
  - Close() // drains and blocks until all submitted tasks finish
- Guarantee no more than maxConcurrent run at once.
- Do not expose channels in the API.

Constraints and pitfalls
- maxConcurrent <= 0 => treat as 1.
- Close is idempotent and must be blocking.
- Avoid deadlocks when maxConcurrent==1.

Tricky edge case
- Submit after Close must immediately return false and not leak goroutines.
*/

import "sync"

type ThrottledScheduler struct {
	mu      sync.Mutex
	cond    *sync.Cond
	closed  bool
	running int
	max     int
	wg      sync.WaitGroup
}

func NewThrottledScheduler(maxConcurrent int) *ThrottledScheduler { // TODO
	if maxConcurrent <= 0 {
		maxConcurrent = 1
	}
	ts := &ThrottledScheduler{max: maxConcurrent}
	ts.cond = sync.NewCond(&ts.mu)
	return ts
}

func (t *ThrottledScheduler) Submit(fn func()) bool { // TODO
	if fn == nil {
		return false
	}
	return false
}

func (t *ThrottledScheduler) Close() { // TODO
}
