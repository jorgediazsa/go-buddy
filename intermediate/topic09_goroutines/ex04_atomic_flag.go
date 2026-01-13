package topic09_goroutines

/*
Title: EX04 — AtomicFlag for mutual exclusion (no channels)

Why this matters
- Mutual exclusion is a core primitive. Implementing a simple CAS-based flag clarifies atomic vs. blocking semantics.

Requirements
- Implement AtomicFlag with methods:
  - TryLock() bool  // atomically acquires the flag if free
  - Unlock()        // releases, panics if not held
  - Locked() bool   // reports current state
- Provide linearizable semantics under concurrency.

Constraints and pitfalls
- Use sync/atomic; do not spin forever in TryLock (it must attempt once).
- Unlock must make the flag available to other goroutines.

Tricky edge case
- Double Unlock must panic; Unlock on zero value (never locked) must panic.
*/

import "sync/atomic"

type AtomicFlag struct{ v int32 }

func (f *AtomicFlag) TryLock() bool { // TODO: implement using CAS from 0->1
	return false
}

func (f *AtomicFlag) Unlock() { // TODO: implement, panic if not held
}

func (f *AtomicFlag) Locked() bool { // TODO
	return atomic.LoadInt32(&f.v) == 1
}
