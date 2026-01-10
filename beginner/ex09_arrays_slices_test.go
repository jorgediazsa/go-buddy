package beginner

import "testing"

func TestEX09_NewIdentity(t *testing.T) {
	cases := []struct {
		n  int
		ok bool
	}{
		{1, true}, {2, true}, {3, true}, {0, false}, {-1, false}, {65, false},
	}
	for i, tc := range cases {
		t.Run("EX09_NewIdentity_"+string(rune('A'+i)), func(t *testing.T) {
			m, err := NewIdentity(tc.n)
			if tc.ok && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error")
			}
			if !tc.ok {
				return
			}
			if len(m) != tc.n {
				t.Fatalf("rows %d want %d", len(m), tc.n)
			}
			for r := 0; r < tc.n; r++ {
				if len(m[r]) != tc.n {
					t.Fatalf("row %d len %d", r, len(m[r]))
				}
				for c := 0; c < tc.n; c++ {
					want := 0
					if r == c {
						want = 1
					}
					if m[r][c] != want {
						t.Fatalf("m[%d][%d]=%d want %d", r, c, m[r][c], want)
					}
				}
			}
			// Row independence test
			if tc.n >= 2 {
				m[0][1] = 7
				if m[1][1] != 1 {
					t.Fatalf("rows share backing array; got %d want 1", m[1][1])
				}
			}
		})
	}
}

func TestEX09_DeepCopy2D(t *testing.T) {
	var nil2D [][]int
	ragged := [][]int{{1}, {2, 3}, {}}
	cases := []struct {
		in [][]int
	}{
		{nil2D},
		{[][]int{}},
		{[][]int{{}}},
		{[][]int{{1, 2}, {3, 4}}},
		{ragged},
	}
	for i, tc := range cases {
		t.Run("EX09_DeepCopy2D_"+string(rune('A'+i)), func(t *testing.T) {
			out := DeepCopy2D(tc.in)
			if (tc.in == nil) != (out == nil) {
				t.Fatalf("nil mismatch")
			}
			if tc.in == nil {
				return
			}
			if len(out) != len(tc.in) {
				t.Fatalf("len mismatch")
			}
			for r := range tc.in {
				if (tc.in[r] == nil) != (out[r] == nil) {
					t.Fatalf("row %d nil mismatch", r)
				}
				if len(out[r]) != len(tc.in[r]) {
					t.Fatalf("row len mismatch %d", r)
				}
				for c := range tc.in[r] {
					if out[r][c] != tc.in[r][c] {
						t.Fatalf("val mismatch")
					}
				}
			}
			// mutate out, ensure in unchanged
			if len(out) > 0 && len(out[0]) > 0 {
				out[0][0]++
				if tc.in[0] != nil && len(tc.in[0]) > 0 && out[0][0] == tc.in[0][0] {
					t.Fatalf("not a deep copy")
				}
			}
		})
	}
}

func TestEX09_SumArray3(t *testing.T) {
	cases := []struct {
		in   [3]int
		want int
	}{
		{[3]int{0, 0, 0}, 0},
		{[3]int{1, 2, 3}, 6},
		{[3]int{-1, 1, 0}, 0},
		{[3]int{100, 200, 300}, 600},
		{[3]int{-5, -5, -5}, -15},
		{[3]int{int(^uint(0) >> 1), 0, 0}, int(^uint(0) >> 1)},
		{[3]int{7, 7, 7}, 21},
		{[3]int{9, -3, 4}, 10},
	}
	for i, tc := range cases {
		t.Run("EX09_SumArray3_"+string(rune('A'+i)), func(t *testing.T) {
			got := SumArray3(tc.in)
			if got != tc.want {
				t.Fatalf("got %d want %d", got, tc.want)
			}
		})
	}
}
