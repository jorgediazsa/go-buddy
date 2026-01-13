package topic09_goroutines

/*
Title: EX03 — Once initialization with error and visibility

Why this matters
- Double-initialization leads to races and inconsistent state. A robust once pattern must guarantee single execution and proper error propagation.

Requirements
- Implement type OnceInit providing:
  - Do(func() error) error: runs the initializer at most once across all goroutines; returns the same error to all callers.
  - Done() bool: reports whether initialization attempted (even if failed).
- Memory visibility: after Do returns nil to any caller, subsequent Load calls on provided access function must see fully initialized state.

Constraints and pitfalls
- Use sync.Once or your own CAS/mutex; ensure error is captured only from the first initializer call.
- Do must be safe under heavy concurrency.

Tricky edge case
- If the first initializer returns an error, a later Do must NOT re-run the initializer and must return the same error.
*/

import "sync"

type OnceInit struct {
	once sync.Once
	mu   sync.Mutex
	done bool
	err  error
}

func (o *OnceInit) Do(fn func() error) error { // TODO: implement
	return nil
}

func (o *OnceInit) Done() bool { // TODO
	return false
}
