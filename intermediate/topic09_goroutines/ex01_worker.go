package topic09_goroutines

/*
Title: EX01 — Start/Stop worker without leaks (no channels)

Why this matters
- Goroutine lifecycle management is foundational. Stop must be idempotent and leave no leaks.

Requirements
- Implement type Worker with methods:
  - Start(): starts a background goroutine that increments an internal counter periodically.
  - Stop(): stops the background goroutine; safe to call multiple times.
  - Count() int64: returns the current counter (atomic-safe).
- No channels; use sync/atomic and sync primitives.

Constraints and pitfalls
- Stop must be blocking: after Stop returns, the goroutine is definitely gone.
- Start called twice should not start multiple goroutines; subsequent Start is a no-op.
- Use a low-cost sleep (e.g., time.Sleep(1ms)) to keep CPU low.

Tricky edge case
- Stop called before Start should be a no-op.
*/

import (
	"sync"
	"sync/atomic"
)

type Worker struct {
	started int32
	stopped int32
	count   int64
	mu      sync.Mutex
	wg      sync.WaitGroup
}

func (w *Worker) Start() { // TODO: implement
}

func (w *Worker) Stop() { // TODO: implement
}

func (w *Worker) Count() int64 { // TODO: implement
	return atomic.LoadInt64(&w.count)
}
