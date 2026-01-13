package topic09_goroutines

/*
Title: EX11 — CondBroadcast: coordinating multiple waiters

Why this matters
- Signaling multiple goroutines efficiently requires sync.Cond. Incorrect usage leads to lost signals or goroutine leaks (waiting forever).

Requirements
- Implement BroadcastNode with:
  - Wait(): blocks until the next signal is received.
  - SignalAll(): wakes up ALL current waiters.
- Use sync.Cond.
- Ensure that calls to Wait() started AFTER SignalAll() do not get woken by that past signal.

Constraints and pitfalls
- Wait() must always be called while holding the lock.
- SignalAll() (Broadcast) should be called to wake everyone.

Tricky edge case
- If SignalAll is called while no one is waiting, the next Wait should still block (signals are not buffered like channels).
*/

import "sync"

type BroadcastNode struct {
	mu   sync.Mutex
	cond *sync.Cond
}

func NewBroadcastNode() *BroadcastNode {
	bn := &BroadcastNode{}
	bn.cond = sync.NewCond(&bn.mu)
	return bn
}

func (bn *BroadcastNode) Wait() { // TODO: implement
}

func (bn *BroadcastNode) SignalAll() { // TODO: implement
}
