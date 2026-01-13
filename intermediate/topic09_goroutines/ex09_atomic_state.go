package topic09_goroutines

/*
Title: EX09 — AtomicState: safe state transitions under concurrency

Why this matters
- Shared state machines are prone to races. Using atomic/CAS ensures valid transitions without locks.

Requirements
- Implement AtomicState with states: 0=Init, 1=Running, 2=Stopped.
- Methods:
  - Load() int
  - TryStart() bool     // Init -> Running only
  - TryStop() bool      // Running -> Stopped only
  - IsTerminal() bool   // Stopped
- All operations must be race-free.

Constraints and pitfalls
- Do not allow Starting from Running, or Stopping from Init.
- Multiple concurrent TryStart must result in at most one success.

Tricky edge case
- TryStop called multiple times should succeed only once.
*/

import "sync/atomic"

type AtomicState struct{ s int32 }

const (
	StateInit    = 0
	StateRunning = 1
	StateStopped = 2
)

func (a *AtomicState) Load() int { return int(atomic.LoadInt32(&a.s)) }

func (a *AtomicState) TryStart() bool { // TODO implement CAS Init->Running
	return false
}

func (a *AtomicState) TryStop() bool { // TODO implement CAS Running->Stopped
	return false
}

func (a *AtomicState) IsTerminal() bool { return a.Load() == StateStopped }
