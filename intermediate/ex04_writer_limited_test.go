package intermediate

import (
	"bytes"
	"io"
	"testing"
)

func TestEX04_LimitedBuffer_WriteAndState(t *testing.T) {
	cases := []struct {
		cap       int
		writes    [][]byte
		wantN     []int
		wantErr   []error
		wantLen   int
		wantBytes []byte
	}{
		{0, [][]byte{}, []int{}, []error{}, 0, nil},
		{0, [][]byte{[]byte("a")}, []int{0}, []error{io.ErrShortWrite}, 0, nil},
		{3, [][]byte{[]byte("ab")}, []int{2}, []error{nil}, 2, []byte("ab")},
		{3, [][]byte{[]byte("abc")}, []int{3}, []error{nil}, 3, []byte("abc")},
		{3, [][]byte{[]byte("abcd")}, []int{3}, []error{io.ErrShortWrite}, 3, []byte("abc")},
		{5, [][]byte{[]byte("ab"), []byte("cde")}, []int{2, 3}, []error{nil, nil}, 5, []byte("abcde")},
		{5, [][]byte{[]byte("ab"), []byte("cdef")}, []int{2, 3}, []error{nil, io.ErrShortWrite}, 5, []byte("abcde")},
		{2, [][]byte{[]byte("a"), []byte("b"), []byte("c")}, []int{1, 1, 0}, []error{nil, nil, io.ErrShortWrite}, 2, []byte("ab")},
		{1, [][]byte{[]byte(""), []byte("x")}, []int{0, 1}, []error{nil, nil}, 1, []byte("x")},
		{4, [][]byte{[]byte("xy"), []byte(""), []byte("z")}, []int{2, 0, 1}, []error{nil, nil, nil}, 3, []byte("xyz")},
		{6, [][]byte{[]byte("ab"), []byte("cd"), []byte("efg")}, []int{2, 2, 2}, []error{nil, nil, io.ErrShortWrite}, 6, []byte("abcdef")},
	}
	for i, tc := range cases {
		t.Run("EX04_LimitedBuffer_"+string(rune('A'+i)), func(t *testing.T) {
			b := NewLimitedBuffer(tc.cap)
			for j, w := range tc.writes {
				n, err := b.Write(w)
				wantN := tc.wantN[j]
				wantErr := tc.wantErr[j]
				if n != wantN {
					t.Fatalf("write %d: n=%d want %d", j, n, wantN)
				}
				if (err == nil) != (wantErr == nil) {
					t.Fatalf("write %d: err=%v want %v", j, err, wantErr)
				}
			}
			if b.Len() != tc.wantLen {
				t.Fatalf("len=%d want %d", b.Len(), tc.wantLen)
			}
			if b.Cap() != tc.cap {
				t.Fatalf("cap=%d want %d", b.Cap(), tc.cap)
			}
			if got := b.Bytes(); !bytes.Equal(got, tc.wantBytes) {
				t.Fatalf("bytes=%q want %q", got, tc.wantBytes)
			}
		})
	}
}

func TestEX04_LimitedBuffer_Reset(t *testing.T) {
	b := NewLimitedBuffer(3)
	_, _ = b.Write([]byte("ab"))
	b.Reset()
	if b.Len() != 0 || b.Cap() != 3 {
		t.Fatalf("after reset len=%d cap=%d", b.Len(), b.Cap())
	}
	_, err := b.Write([]byte("abcd"))
	if err == nil {
		t.Fatalf("expected short write")
	}
}
