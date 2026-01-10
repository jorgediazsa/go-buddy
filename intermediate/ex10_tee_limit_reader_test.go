package intermediate

import (
	"bytes"
	"io"
	"testing"
)

type writeRecorder struct{ bytes.Buffer }

func TestEX10_TeeLimitReader_BasicAndLimits(t *testing.T) {
	cases := []struct {
		src     string
		limit   int64
		chunk   int
		wantOut string
		wantTee string
	}{
		{"", 0, 4, "", ""},
		{"hello", 0, 2, "hello", ""},
		{"hello", 2, 2, "hello", "he"},
		{"abcdef", 3, 4, "abcdef", "abc"},
		{"abcdef", 10, 2, "abcdef", "abcdef"},
		{"a", 1, 1, "a", "a"},
		{"abcd", 1, 3, "abcd", "a"},
		{"abcd", 2, 1, "abcd", "ab"},
		{"abcd", 3, 1, "abcd", "abc"},
		{"abcd", 4, 10, "abcd", "abcd"},
		{"xyz", 5, 1, "xyz", "xyz"},
	}
	for i, tc := range cases {
		t.Run("EX10_TeeLimit_"+string(rune('A'+i)), func(t *testing.T) {
			var dst writeRecorder
			tr := NewTeeLimitReader(bytes.NewBufferString(tc.src), &dst, tc.limit)
			buf := make([]byte, tc.chunk)
			var out bytes.Buffer
			for {
				n, err := tr.Read(buf)
				if n > 0 {
					out.Write(buf[:n])
				}
				if err == io.EOF {
					break
				}
				if err != nil {
					break
				}
				if n == 0 {
					break
				}
			}
			if out.String() != tc.wantOut {
				t.Fatalf("out=%q want %q", out.String(), tc.wantOut)
			}
			if dst.String() != tc.wantTee {
				t.Fatalf("tee=%q want %q", dst.String(), tc.wantTee)
			}
		})
	}
}
