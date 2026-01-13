package topic09_goroutines

/*
Title: EX10 — WorkQueue with mutex/cond and shutdown

Why this matters
- Waiting on never-signaled conditions leaks goroutines. A correct queue must wake waiting consumers on shutdown.

Requirements
- Implement WorkQueue of ints with:
  - Enqueue(v int) bool        // false if closed
  - Dequeue() (int, bool)      // blocks until item or closed and empty; ok=false on closed+empty
  - Close()                    // idempotent; wakes all waiters
- Use sync.Mutex + sync.Cond; no channel-based API.

Constraints and pitfalls
- Dequeue must not busy-wait.
- After Close, Enqueue returns false and Dequeue returns ok=false once drained.

Tricky edge case
- Multiple Dequeue waiters must all be woken on Close.
*/

import "sync"

type WorkQueue struct {
	mu     sync.Mutex
	cond   *sync.Cond
	q      []int
	closed bool
}

func NewWorkQueue() *WorkQueue {
	w := &WorkQueue{}
	w.cond = sync.NewCond(&w.mu)
	return w
}

func (w *WorkQueue) Enqueue(v int) bool { // TODO implement
	return false
}

func (w *WorkQueue) Dequeue() (int, bool) { // TODO implement
	return 0, false
}

func (w *WorkQueue) Close() { // TODO implement
}
