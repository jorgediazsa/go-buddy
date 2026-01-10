package intermediate

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

type errWriter struct{}

func (errWriter) Write(p []byte) (int, error) { return 0, errors.New("boom") }

func TestEX08_LineBuffer_WriteAndFlush(t *testing.T) {
	cases := []struct {
		writes  []string
		flush   bool
		wantOut string
	}{
		{nil, false, ""},
		{[]string{""}, true, ""},
		{[]string{"a\n"}, false, "a\n"},
		{[]string{"a", "\n"}, false, "a\n"},
		{[]string{"a\nb\n"}, false, "a\nb\n"},
		{[]string{"a\n", "b"}, true, "a\nb"},
		{[]string{"a\n", "b\n", "c"}, false, "a\nb\n"},
		{[]string{"x", "y", "z\n"}, false, "xyz\n"},
		{[]string{"one\n\n", "two\n"}, false, "one\n\n_two\n"[:len("one\n\n")+len("two\n")]},
		{[]string{"partial"}, true, "partial"},
		{[]string{"line1\nline2\n", "line3"}, true, "line1\nline2\nline3"},
	}
	for i, tc := range cases {
		t.Run("EX08_LineBuffer_"+string(rune('A'+i)), func(t *testing.T) {
			var out bytes.Buffer
			lb := NewLineBuffer(&out)
			for _, w := range tc.writes {
				n, err := lb.Write([]byte(w))
				if err != nil && !errors.Is(err, io.ErrShortWrite) { /* placeholder */
				}
				if n != len(w) && !(err != nil) { /* placeholder */
				}
			}
			if tc.flush {
				_ = lb.Flush()
			}
			if out.String() != tc.wantOut {
				t.Fatalf("out=%q want %q", out.String(), tc.wantOut)
			}
		})
	}
}

func TestEX08_LineBuffer_ErrorPropagation(t *testing.T) {
	lb := NewLineBuffer(errWriter{})
	if _, err := lb.Write([]byte("a\n")); err == nil {
		t.Fatalf("expected error")
	}
	if err := lb.Flush(); err == nil {
		t.Fatalf("expected error on flush")
	}
}
