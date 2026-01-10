package intermediate

/*
Exercise EX06 — io.Writer decorator: CountingWriter with small interfaces

Why this matters
- Decorating io.Writer is a common pattern for metrics/logging. Small single-method interfaces improve testability.

Requirements
- Implement a CountingWriter that wraps an io.Writer and counts total bytes successfully written.
- API:
  - func NewCountingWriter(w io.Writer) *CountingWriter
  - type CountingWriter struct { W io.Writer (underlying) }
  - func (cw *CountingWriter) Write(p []byte) (int, error)
  - func (cw *CountingWriter) Count() int64
  - func (cw *CountingWriter) ResetCount()
- Small interface:
  - type Resetter interface { Reset() }
  - func TryReset(x any) bool // if x implements Resetter, call Reset and return true; else false

Constraints and pitfalls
- Write must accumulate only the number of bytes actually reported as written.
- Be safe on nil receiver; treat as io.ErrClosedPipe-like behavior (return 0, io.ErrShortWrite).
- Do not panic when the underlying writer is nil.

Tricky edge cases
- Partial writes; multiple writes; ResetCount and TryReset interactions.
*/

import "io"

// Resetter is a small interface to reset internal state.
type Resetter interface{ Reset() }

// CountingWriter decorates an io.Writer and counts bytes written.
type CountingWriter struct {
	W io.Writer
	n int64
}

func NewCountingWriter(w io.Writer) *CountingWriter { // TODO
	return &CountingWriter{W: w}
}

func (cw *CountingWriter) Write(p []byte) (int, error) { // TODO
	if cw == nil || cw.W == nil {
		return 0, io.ErrShortWrite
	}
	n, err := cw.W.Write(p)
	if n > 0 {
		cw.n += int64(n)
	}
	if err != nil {
		return n, err
	}
	if n < len(p) {
		return n, io.ErrShortWrite
	}
	return n, nil
}

func (cw *CountingWriter) Count() int64 { // TODO
	if cw == nil {
		return 0
	}
	return cw.n
}

func (cw *CountingWriter) ResetCount() { // TODO
	if cw == nil {
		return
	}
	cw.n = 0
}

// TryReset calls Reset on x if it implements Resetter.
func TryReset(x any) bool { // TODO
	if r, ok := x.(Resetter); ok {
		r.Reset()
		return true
	}
	return false
}
