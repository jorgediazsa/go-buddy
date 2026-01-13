package topic09_goroutines

/*
Title: EX06 — BackgroundRefresher with Start/Stop and restart support

Why this matters
- Long-lived services often keep data fresh in the background. Leaking refreshers creates subtle memory/timer leaks.

Requirements
- Implement BackgroundRefresher that:
  - Start(period time.Duration, refresh func() error) error
    * Starts a background goroutine that calls refresh() periodically (best-effort; ignore refresh error for scheduling but return first error from Start if period <= 0 or nil refresh).
    * Calling Start while already running should be a no-op (returns nil) and not start another goroutine.
  - Stop() // blocks until goroutine exits; idempotent
  - Ticks() int64 // number of successful calls to refresh() (atomic)
- No channel-based API. You may use timers or time.Sleep internally.

Constraints and pitfalls
- Stop must guarantee that the goroutine and any associated timers are stopped before returning.
- After Stop, Start can be called again to restart with a new period.

Tricky edge case
- Stop called immediately after Start should not leak the goroutine or timer.
*/

import (
	"sync"
	"sync/atomic"
	"time"
)

type BackgroundRefresher struct {
	mu      sync.Mutex
	running bool
	ticks   int64
	stopCh  chan struct{}
	wg      sync.WaitGroup
}

func (b *BackgroundRefresher) Start(period time.Duration, refresh func() error) error { // TODO implement
	return nil
}

func (b *BackgroundRefresher) Stop() { // TODO implement
}

func (b *BackgroundRefresher) Ticks() int64 { return atomic.LoadInt64(&b.ticks) }

// Sentinel for invalid parameters; used to compile before user changes.
var ErrBadParam = &badParamError{"bad param"}

type badParamError struct{ s string }

func (e *badParamError) Error() string { return e.s }
