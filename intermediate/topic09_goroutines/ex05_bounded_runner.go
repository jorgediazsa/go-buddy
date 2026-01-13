package topic09_goroutines

/*
Title: EX05 — BoundedRunner: limit concurrent tasks, guarantee shutdown

Why this matters
- Unbounded goroutine creation causes resource exhaustion. A bounded runner caps concurrency and shuts down cleanly.

Requirements
- Implement BoundedRunner with:
  - NewBoundedRunner(limit int) *BoundedRunner
  - Run(fn func()) bool // schedules fn if runner not closed; returns false if closed
  - Close()             // waits for all started tasks to finish; idempotent
  - InFlight() int      // number of currently running tasks
- Enforce at most `limit` goroutines run tasks simultaneously. Extra Run calls should block or reject until capacity frees.
- Do not use channels for API; you may use internal channels or conds if needed.

Constraints and pitfalls
- Close must be blocking; after Close, Run must return false and not start new work.
- Avoid deadlocks when limit==1 and tasks call nested Run.

Tricky edge case
- limit <= 0 should be treated as 1.
*/

import (
	"sync"
)

type BoundedRunner struct {
	mu      sync.Mutex
	cond    *sync.Cond
	limit   int
	running int
	closed  bool
	wg      sync.WaitGroup
}

func NewBoundedRunner(limit int) *BoundedRunner { // TODO implement
	if limit <= 0 {
		limit = 1
	}
	br := &BoundedRunner{limit: limit}
	br.cond = sync.NewCond(&br.mu)
	return br
}

func (b *BoundedRunner) Run(fn func()) bool { // TODO
	if fn == nil {
		return false
	}
	return false
}

func (b *BoundedRunner) Close() { // TODO
}

func (b *BoundedRunner) InFlight() int { // TODO
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.running
}
