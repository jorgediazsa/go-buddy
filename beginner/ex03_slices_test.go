package beginner

import "testing"

func TestEX03_CloneInts(t *testing.T) {
	cases := []struct{ in []int }{
		{nil},
		{[]int{}},
		{[]int{1}},
		{[]int{1, 2, 3, 4}},
	}
	for i, tc := range cases {
		t.Run("EX03_CloneInts_"+string(rune('A'+i)), func(t *testing.T) {
			out := CloneInts(tc.in)
			if tc.in == nil {
				if out != nil {
					t.Fatalf("want nil, got %v", out)
				}
				return
			}
			if len(tc.in) != len(out) {
				t.Fatalf("length mismatch: %d vs %d", len(tc.in), len(out))
			}
			for i := range tc.in {
				if tc.in[i] != out[i] {
					t.Fatalf("value mismatch at %d", i)
				}
			}
			if len(tc.in) == 0 {
				// ensure distinct backing: append to out should not change in
				out = append(out, 9)
				if len(tc.in) != 0 {
					t.Fatalf("input mutated for empty slice")
				}
			} else {
				out[0] ^= 1
				if out[0] == tc.in[0] {
					t.Fatalf("expected deep copy, mutation leaked")
				}
			}
		})
	}
}

func TestEX03_RotateLeftInPlace(t *testing.T) {
	cases := []struct {
		in   []int
		k    int
		ok   bool
		want []int
	}{
		{[]int{}, 0, true, []int{}},
		{[]int{1}, 0, true, []int{1}},
		{[]int{1}, 1, true, []int{1}},
		{[]int{1, 2, 3}, 0, true, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 1, true, []int{2, 3, 1}},
		{[]int{1, 2, 3}, 2, true, []int{3, 1, 2}},
		{[]int{1, 2, 3}, 3, true, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 7, true, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2}, -1, false, []int{}},
	}
	for i, tc := range cases {
		t.Run("EX03_RotateLeftInPlace_"+string(rune('A'+i)), func(t *testing.T) {
			in := make([]int, len(tc.in))
			copy(in, tc.in)
			err := RotateLeftInPlace(in, tc.k)
			if tc.ok && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if tc.ok {
				if len(in) != len(tc.want) {
					t.Fatalf("len mismatch")
				}
				for i := range in {
					if in[i] != tc.want[i] {
						t.Fatalf("mismatch at %d: %v vs %v", i, in, tc.want)
					}
				}
			}
		})
	}
}

func TestEX03_SortedInsert(t *testing.T) {
	cases := []struct {
		in   []int
		v    int
		want []int
	}{
		{nil, 5, []int{5}},
		{[]int{}, 5, []int{5}},
		{[]int{1, 3, 5}, 0, []int{0, 1, 3, 5}},
		{[]int{1, 3, 5}, 6, []int{1, 3, 5, 6}},
		{[]int{1, 3, 5}, 4, []int{1, 3, 4, 5}},
		{[]int{1, 3, 5}, 3, []int{1, 3, 3, 5}},
		{[]int{2, 2, 2}, 2, []int{2, 2, 2, 2}},
		{[]int{-5, 0, 10}, -5, []int{-5, -5, 0, 10}},
	}
	for i, tc := range cases {
		t.Run("EX03_SortedInsert_"+string(rune('A'+i)), func(t *testing.T) {
			in := make([]int, len(tc.in))
			copy(in, tc.in)
			out := SortedInsert(in, tc.v)
			// ensure input not mutated
			for j := range in {
				if in[j] != tc.in[j] {
					t.Fatalf("input mutated")
				}
			}
			if len(out) != len(tc.want) {
				t.Fatalf("len mismatch: %d vs %d", len(out), len(tc.want))
			}
			for j := range out {
				if out[j] != tc.want[j] {
					t.Fatalf("mismatch at %d: %v vs %v", j, out, tc.want)
				}
			}
		})
	}
}
