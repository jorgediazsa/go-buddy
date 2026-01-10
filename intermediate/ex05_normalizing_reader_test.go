package intermediate

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func readAll(r io.Reader) []byte {
	b, _ := io.ReadAll(r)
	return b
}

func TestEX05_NormalizingReader_BOM_CRLF_Trim(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"\xEF\xBB\xBF", ""},
		{"\xEF\xBB\xBFabc", "abc"},
		{"a\r\nb\r\nc", "a\nb\nc"},
		{"a  \r\n\t\t\r\n", "a\n\n"},
		{"line with spaces   \nnext\t\t\n", "line with spaces\nnext\n"},
		{"no newline", "no newline"},
		{"trail space \r\ntrail\t\nend", "trail\ntrail\nend"},
		{"\rsolo\r\n", "\nsolo\n"},
		{"mix\r\nlines\nend\r", "mix\nlines\nend\n"},
		{"tabs\t\t\n\r\nX", "tabs\n\nX"},
	}
	for i, tc := range cases {
		t.Run("EX05_Normalize_"+string(rune('A'+i)), func(t *testing.T) {
			r := NewNormalizingReader(strings.NewReader(tc.in))
			got := readAll(r)
			if !bytes.Equal(got, []byte(tc.want)) {
				t.Fatalf("got %q want %q", got, tc.want)
			}
		})
	}
}
