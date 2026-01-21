package beginner

import "testing"

func TestEX03_CloneInts(t *testing.T) {
	t.Run("nil_to_nil", func(t *testing.T) {
		var in []int = nil
		out := CloneInts(in)
		if out != nil {
			t.Fatalf("want nil, got %#v", out)
		}
	})

	t.Run("empty_non_nil_stays_non_nil", func(t *testing.T) {
		in := []int{}
		out := CloneInts(in)

		if out == nil {
			t.Fatalf("want non-nil empty slice, got nil")
		}
		if len(out) != 0 {
			t.Fatalf("want len 0, got %d", len(out))
		}
	})

	type tc struct {
		name string
		in   []int
	}
	cases := []tc{
		{"one", []int{1}},
		{"many", []int{1, 2, 3, 4}},
		{"with_zero_values", []int{0, 0, 0}},
		{"with_negatives", []int{-1, -2, 3}},
	}

	for _, c := range cases {
		c := c
		t.Run("deep_copy_"+c.name, func(t *testing.T) {
			in := append([]int(nil), c.in...) // defensive: keep original stable
			out := CloneInts(in)

			// same contents
			if len(out) != len(in) {
				t.Fatalf("length mismatch: in=%d out=%d", len(in), len(out))
			}
			for i := range in {
				if in[i] != out[i] {
					t.Fatalf("value mismatch at %d: in=%d out=%d", i, in[i], out[i])
				}
			}

			// mutate out, in must not change
			if len(out) > 0 {
				orig := in[0]
				out[0] = orig + 999
				if in[0] != orig {
					t.Fatalf("input mutated when output changed: in[0]=%d want %d", in[0], orig)
				}
			}

			// mutate in, out must not change
			if len(in) > 0 {
				orig := out[0]
				in[0] = out[0] + 555
				if out[0] != orig {
					t.Fatalf("output mutated when input changed: out[0]=%d want %d", out[0], orig)
				}
			}
		})
	}
}

func TestEX03_RotateLeftInPlace(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		k    int
		ok   bool
		want []int
	}{
		{"empty_k0", []int{}, 0, true, []int{}},
		{"empty_k5", []int{}, 5, true, []int{}}, // must not panic
		{"one_k0", []int{1}, 0, true, []int{1}},
		{"one_k1", []int{1}, 1, true, []int{1}},
		{"two_k0", []int{1, 2}, 0, true, []int{1, 2}},
		{"two_k1", []int{1, 2}, 1, true, []int{2, 1}},
		{"three_k0", []int{1, 2, 3}, 0, true, []int{1, 2, 3}},
		{"three_k1", []int{1, 2, 3}, 1, true, []int{2, 3, 1}},
		{"three_k2", []int{1, 2, 3}, 2, true, []int{3, 1, 2}},
		{"three_k3", []int{1, 2, 3}, 3, true, []int{1, 2, 3}},
		{"three_k4", []int{1, 2, 3}, 4, true, []int{2, 3, 1}}, // k>len
		{"five_k7", []int{1, 2, 3, 4, 5}, 7, true, []int{3, 4, 5, 1, 2}},
		{"k_negative", []int{1, 2}, -1, false, nil},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			in := append([]int(nil), tc.in...) // clone input so we can mutate safely

			err := RotateLeftInPlace(in, tc.k)
			if tc.ok {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if len(in) != len(tc.want) {
					t.Fatalf("len mismatch: got %d want %d", len(in), len(tc.want))
				}
				for i := range in {
					if in[i] != tc.want[i] {
						t.Fatalf("mismatch at %d: got %v want %v", i, in, tc.want)
					}
				}
			} else {
				if err == nil {
					t.Fatalf("expected error, got nil (got=%v)", in)
				}
			}
		})
	}
}

func TestEX03_SortedInsert(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		v    int
		want []int
	}{
		{"nil", nil, 5, []int{5}},
		{"empty", []int{}, 5, []int{5}},

		{"insert_head", []int{1, 3, 5}, 0, []int{0, 1, 3, 5}},
		{"insert_tail", []int{1, 3, 5}, 6, []int{1, 3, 5, 6}},
		{"insert_mid", []int{1, 3, 5}, 4, []int{1, 3, 4, 5}},

		// stability: insert AFTER equals (upper bound)
		{"insert_equal_single", []int{1, 3, 5}, 3, []int{1, 3, 3, 5}},
		{"insert_equal_all", []int{2, 2, 2}, 2, []int{2, 2, 2, 2}},
		{"insert_equal_many", []int{1, 2, 2, 2, 3}, 2, []int{1, 2, 2, 2, 2, 3}},

		{"insert_negative", []int{-5, 0, 10}, -5, []int{-5, -5, 0, 10}},
		{"insert_less_than_all", []int{10, 20}, 5, []int{5, 10, 20}},
		{"insert_greater_than_all", []int{10, 20}, 30, []int{10, 20, 30}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			orig := append([]int(nil), tc.in...) // exact snapshot of input

			out := SortedInsert(tc.in, tc.v)

			// input not mutated
			if len(tc.in) != len(orig) {
				t.Fatalf("input len mutated: got %d want %d", len(tc.in), len(orig))
			}
			for i := range orig {
				if tc.in[i] != orig[i] {
					t.Fatalf("input mutated at %d: got %v want %v", i, tc.in, orig)
				}
			}

			// output correct
			if len(out) != len(tc.want) {
				t.Fatalf("len mismatch: got %d want %d", len(out), len(tc.want))
			}
			for i := range out {
				if out[i] != tc.want[i] {
					t.Fatalf("mismatch at %d: got %v want %v", i, out, tc.want)
				}
			}

			// sanity: output is sorted (ascending)
			for i := 1; i < len(out); i++ {
				if out[i-1] > out[i] {
					t.Fatalf("output not sorted: %v", out)
				}
			}
		})
	}
}
