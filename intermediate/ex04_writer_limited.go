package intermediate

/*
Exercise EX04 — io.Writer: bounded buffer with partial writes

Why this matters
- Writers often have capacity limits (e.g., network MTU, quotas). Handling partial writes and state correctly is crucial.

Requirements
- Implement LimitedBuffer with:
  - NewLimitedBuffer(cap int) *LimitedBuffer
  - func (b *LimitedBuffer) Write(p []byte) (n int, err error)
    * Writes as much as fits until capacity; when full, return n<len(p) and an error
  - func (b *LimitedBuffer) Bytes() []byte   // returns a copy of the current contents
  - func (b *LimitedBuffer) Cap() int        // total capacity
  - func (b *LimitedBuffer) Len() int        // current length
  - func (b *LimitedBuffer) Reset()          // empties contents but preserves capacity

Constraints and pitfalls
- Do not panic on nil receiver; methods should be safe no-ops where reasonable.
- Write must not exceed capacity; return io.ErrShortWrite when not all data fits.
- Bytes must return a copy to avoid external mutation of internal buffer.

Tricky edge cases
- cap==0: all writes should return 0, io.ErrShortWrite when p is non-empty.
- Multiple writes exactly filling capacity.
- Subsequent writes after full should continue to return 0, io.ErrShortWrite.
*/

import "io"

// LimitedBuffer is a bounded in-memory buffer implementing io.Writer.
type LimitedBuffer struct {
	cap int
	buf []byte
}

// NewLimitedBuffer creates a new LimitedBuffer with given capacity.
func NewLimitedBuffer(cap int) *LimitedBuffer { // TODO
	return &LimitedBuffer{cap: cap}
}

// Write appends up to capacity and returns io.ErrShortWrite if not all data fit.
func (b *LimitedBuffer) Write(p []byte) (int, error) { // TODO
	if b == nil {
		return 0, io.ErrShortWrite
	}
	return 0, io.ErrShortWrite
}

// Bytes returns a copy of the internal buffer.
func (b *LimitedBuffer) Bytes() []byte { // TODO
	return nil
}

// Cap returns the capacity.
func (b *LimitedBuffer) Cap() int { // TODO
	if b == nil {
		return 0
	}
	return b.cap
}

// Len returns the current length.
func (b *LimitedBuffer) Len() int { // TODO
	if b == nil {
		return 0
	}
	return len(b.buf)
}

// Reset empties the buffer.
func (b *LimitedBuffer) Reset() { // TODO
	if b == nil {
		return
	}
}
