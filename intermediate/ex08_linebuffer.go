package intermediate

/*
Exercise EX08 — io.Writer + small interface: LineBuffer with Flush

Why this matters
- Writers that buffer until a delimiter (e.g., newline) are common. Exposing a tiny Flush method improves control and testability.

Requirements
- Implement LineBuffer that buffers writes until a '\n' is seen, then forwards complete lines to an underlying io.Writer.
- API:
  - func NewLineBuffer(dst io.Writer) *LineBuffer
  - type LineBuffer struct { W io.Writer }
  - func (lb *LineBuffer) Write(p []byte) (int, error)
  - func (lb *LineBuffer) Flush() error // writes any pending bytes as a final line (without adding extra newline)
- Small interface:
  - type Flusher interface { Flush() error }

Constraints and pitfalls
- Preserve newlines: do not strip or add extra newlines except when flushing pending bytes (no newline appended).
- Handle multiple lines and partial segments across writes.
- Be safe on nil receivers; return io.ErrShortWrite for operations when underlying is nil.

Tricky edge cases
- Empty writes; multiple consecutive newlines; Flush on empty buffer should be a no-op.
*/

import "io"

// Flusher is a small interface for buffered writers.
type Flusher interface{ Flush() error }

// LineBuffer buffers until newline and forwards complete lines.
type LineBuffer struct {
	W   io.Writer
	buf []byte
}

func NewLineBuffer(dst io.Writer) *LineBuffer { // TODO
	return &LineBuffer{W: dst}
}

func (lb *LineBuffer) Write(p []byte) (int, error) { // TODO
	if lb == nil || lb.W == nil {
		return 0, io.ErrShortWrite
	}
	return 0, io.ErrShortWrite
}

func (lb *LineBuffer) Flush() error { // TODO
	if lb == nil || lb.W == nil {
		return io.ErrShortWrite
	}
	return nil
}
