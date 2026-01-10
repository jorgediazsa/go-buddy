package intermediate

/*
Exercise EX05 — io.Reader: normalization (BOM strip, CRLF→LF, trim trailing spaces at line end)

Why this matters
- Readers frequently need to sanitize inputs. Implementing transformation readers is a core Go skill.

Requirements
- Implement a NormalizingReader that wraps an underlying io.Reader and applies:
  1) Strips a leading UTF-8 BOM (\xEF\xBB\xBF) if present at the very start of the stream.
  2) Converts CRLF ("\r\n") sequences to LF ("\n"). Lone "\r" should become "\n" as well.
  3) Trims trailing spaces and tabs (" ", "\t") at the end of each line (before the newline).
- API:
  - func NewNormalizingReader(r io.Reader) io.Reader
  - type NormalizingReader struct { R io.Reader (unexported state allowed) }
  - func (nr *NormalizingReader) Read(p []byte) (int, error)

Constraints and pitfalls
- Do not introduce concurrency; keep state minimal and deterministic.
- Preserve exact number of newlines; do not add or remove extra line breaks.
- Handle small Read buffers (Read may be called with p of any size).

Tricky edge cases
- Empty input; input containing only BOM; inputs ending without a trailing newline.
- Windows newlines across buffer boundaries ("\r" at end of one Read, "\n" at start of next).
- Lines with only spaces/tabs should trim to empty lines.
*/

import "io"

// NormalizingReader wraps an io.Reader applying normalization.
type NormalizingReader struct {
	R io.Reader
	// internal TODO state: carryover bytes, seenBOM, pending bytes, etc.
}

// NewNormalizingReader returns a new reader applying normalization.
func NewNormalizingReader(r io.Reader) io.Reader { // TODO
	return &NormalizingReader{R: r}
}

// Read implements io.Reader.
func (nr *NormalizingReader) Read(p []byte) (int, error) { // TODO
	if nr == nil || nr.R == nil {
		return 0, io.EOF
	}
	return 0, io.EOF
}
