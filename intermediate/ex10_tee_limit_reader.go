package intermediate

/*
Exercise EX10 — io.Reader composition: TeeLimitReader

Why this matters
- Composing readers/writers with limits is a common pattern (logging, metering, hashing). Getting edge cases right is crucial.

Requirements
- Implement a reader that tees all bytes read from src into dst up to a maximum of limit bytes. After the limit is reached, continue reading from src but do not write further to dst.
- API:
  - func NewTeeLimitReader(src io.Reader, dst io.Writer, limit int64) io.Reader
  - type TeeLimitReader struct { Src io.Reader; Dst io.Writer; Limit int64 }
  - func (tr *TeeLimitReader) Read(p []byte) (int, error)

Constraints and pitfalls
- Do not introduce concurrency; keep it stateful and deterministic.
- Propagate errors from src reads. If writing to dst fails before reaching the limit, return the write error.
- Handle small reads; maintain an internal counter of bytes teed so far.

Tricky edge cases
- limit == 0 means never tee (pass-through).
- Partial writes to dst; ensure correct accounting and error handling.
- dst may be nil; treat as no-op teeing.
*/

import "io"

// TeeLimitReader tees up to Limit bytes from Src into Dst during Read.
type TeeLimitReader struct {
	Src   io.Reader
	Dst   io.Writer
	Limit int64
	done  int64 // bytes already written to Dst
}

// NewTeeLimitReader constructs a new TeeLimitReader.
func NewTeeLimitReader(src io.Reader, dst io.Writer, limit int64) io.Reader { // TODO
	return &TeeLimitReader{Src: src, Dst: dst, Limit: limit}
}

// Read implements io.Reader.
func (tr *TeeLimitReader) Read(p []byte) (int, error) { // TODO
	if tr == nil || tr.Src == nil {
		return 0, io.EOF
	}
	return 0, io.EOF
}
