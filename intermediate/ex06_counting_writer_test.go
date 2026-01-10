package intermediate

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

// partialWriter writes at most limit bytes per Write, returning nil error.
type partialWriter struct {
	buf   *bytes.Buffer
	limit int
}

func (p *partialWriter) Write(b []byte) (int, error) {
	if p.limit <= 0 {
		return 0, nil
	}
	n := p.limit
	if n > len(b) {
		n = len(b)
	}
	return p.buf.Write(b[:n])
}

func TestEX06_CountingWriter_Basics(t *testing.T) {
	// nil receiver
	var nilCW *CountingWriter
	if _, err := nilCW.Write([]byte("x")); !errors.Is(err, io.ErrShortWrite) {
		t.Fatalf("nil writer: expected ErrShortWrite, got %v", err)
	}

	cases := []struct {
		writes    []string
		wantCount int64
		wantOut   string
	}{
		{nil, 0, ""},
		{[]string{""}, 0, ""},
		{[]string{"a"}, 1, "a"},
		{[]string{"hello"}, 5, "hello"},
		{[]string{"hello", "!"}, 6, "hello!"},
		{[]string{"x", "y", "z"}, 3, "xyz"},
		{[]string{"ab", "", "cd"}, 4, "abcd"},
		{[]string{"12345", "67890"}, 10, "1234567890"},
		{[]string{"multi", "-", "part", "-", "write"}, 16, "multi-part-write"},
		{[]string{"A", "B", "C", "D", "E"}, 5, "ABCDE"},
	}
	for i, tc := range cases {
		t.Run("EX06_Basics_"+string(rune('A'+i)), func(t *testing.T) {
			var buf bytes.Buffer
			cw := NewCountingWriter(&buf)
			for _, s := range tc.writes {
				n, err := cw.Write([]byte(s))
				if err != nil {
					t.Fatalf("unexpected err: %v", err)
				}
				if n != len(s) {
					t.Fatalf("n=%d want %d", n, len(s))
				}
			}
			if cw.Count() != tc.wantCount {
				t.Fatalf("count=%d want %d", cw.Count(), tc.wantCount)
			}
			if buf.String() != tc.wantOut {
				t.Fatalf("out=%q want %q", buf.String(), tc.wantOut)
			}
		})
	}
}

func TestEX06_CountingWriter_PartialWrites(t *testing.T) {
	var b bytes.Buffer
	pw := &partialWriter{buf: &b, limit: 3}
	cw := NewCountingWriter(pw)
	n, err := cw.Write([]byte("abcdef"))
	if n != 3 || !errors.Is(err, io.ErrShortWrite) {
		t.Fatalf("partial: n=%d err=%v", n, err)
	}
	if cw.Count() != 3 {
		t.Fatalf("count=%d want 3", cw.Count())
	}
	// subsequent write with limit reached still partial
	pw.limit = 2
	n, err = cw.Write([]byte("XYZ"))
	if n != 2 || !errors.Is(err, io.ErrShortWrite) {
		t.Fatalf("partial2: n=%d err=%v", n, err)
	}
	if cw.Count() != 5 {
		t.Fatalf("count=%d want 5", cw.Count())
	}
}

func TestEX06_ResetCountAndTryReset(t *testing.T) {
	var buf bytes.Buffer
	cw := NewCountingWriter(&buf)
	_, _ = cw.Write([]byte("abc"))
	cw.ResetCount()
	if cw.Count() != 0 {
		t.Fatalf("count=%d want 0", cw.Count())
	}
	buf.WriteString("123")
	if !TryReset(&buf) {
		t.Fatalf("TryReset should be true for bytes.Buffer")
	}
	if buf.Len() != 0 {
		t.Fatalf("buffer not reset")
	}
	if TryReset(cw) {
		t.Fatalf("TryReset should be false for CountingWriter without Reset()")
	}
}
