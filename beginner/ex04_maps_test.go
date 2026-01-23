package beginner

import "testing"

func TestEX04_MergeSafe(t *testing.T) {
	clone := func(m map[string]int) map[string]int {
		if m == nil {
			return nil
		}
		c := make(map[string]int, len(m))
		for k, v := range m {
			c[k] = v
		}
		return c
	}

	cases := []struct {
		name string
		a, b map[string]int
		want map[string]int
	}{
		{"nil_nil", nil, nil, map[string]int{}},
		{"empty_empty", map[string]int{}, map[string]int{}, map[string]int{}},
		{"a_only", map[string]int{"x": 1}, nil, map[string]int{"x": 1}},
		{"b_only", nil, map[string]int{"y": 2}, map[string]int{"y": 2}},
		{"override_conflict", map[string]int{"k": 1, "z": 9}, map[string]int{"k": 3}, map[string]int{"k": 3, "z": 9}},
		{"disjoint", map[string]int{"a": 1}, map[string]int{"b": 2}, map[string]int{"a": 1, "b": 2}},
		{"same_value_conflict", map[string]int{"a": 1, "b": 2}, map[string]int{"b": 2}, map[string]int{"a": 1, "b": 2}},
		{"single_key_same", map[string]int{"a": 1}, map[string]int{"a": 1}, map[string]int{"a": 1}},
		{"b_overrides_with_zero", map[string]int{"k": 9}, map[string]int{"k": 0}, map[string]int{"k": 0}},
		{"many_keys", map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"b": 20, "d": 4}, map[string]int{"a": 1, "b": 20, "c": 3, "d": 4}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			aBefore := clone(tc.a)
			bBefore := clone(tc.b)

			got := MergeSafe(tc.a, tc.b)

			// must return non-nil map
			if got == nil {
				t.Fatalf("got nil map")
			}

			// inputs not mutated
			if !mapsEqual(tc.a, aBefore) {
				t.Fatalf("input a mutated: got=%v want=%v", tc.a, aBefore)
			}
			if !mapsEqual(tc.b, bBefore) {
				t.Fatalf("input b mutated: got=%v want=%v", tc.b, bBefore)
			}

			// output correct (exact match)
			if !mapsEqual(got, tc.want) {
				t.Fatalf("got=%v want=%v", got, tc.want)
			}

			// sanity: output is independent (mutating output doesn't affect inputs)
			got["__probe__"] = 123
			if tc.a != nil {
				if _, ok := tc.a["__probe__"]; ok {
					t.Fatalf("output mutation leaked into a")
				}
			}
			if tc.b != nil {
				if _, ok := tc.b["__probe__"]; ok {
					t.Fatalf("output mutation leaked into b")
				}
			}
		})
	}
}

func mapsEqual(a, b map[string]int) bool {
	if a == nil && b == nil {
		return true
	}
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for k, va := range a {
		vb, ok := b[k]
		if !ok || vb != va {
			return false
		}
	}
	return true
}

func TestEX04_SortedKeysByValue(t *testing.T) {
	cases := []struct {
		name string
		in   map[string]int
		want []string
	}{
		{"nil", nil, []string{}},
		{"empty", map[string]int{}, []string{}},
		{"single", map[string]int{"a": 1}, []string{"a"}},
		{"strict_desc", map[string]int{"a": 1, "b": 2, "c": 3}, []string{"c", "b", "a"}},

		// ties: must break by key asc
		{"tie_top", map[string]int{"a": 2, "b": 2, "c": 1}, []string{"a", "b", "c"}},
		{"all_tie", map[string]int{"b": 2, "a": 2, "c": 2}, []string{"a", "b", "c"}},
		{"negatives", map[string]int{"x": -1, "y": -2}, []string{"x", "y"}},
		{"all_zero", map[string]int{"k": 0, "j": 0, "a": 0}, []string{"a", "j", "k"}},

		// edge: mixed ties and ordering
		{"mixed", map[string]int{"aa": 5, "b": 5, "a": 5, "z": 4}, []string{"a", "aa", "b", "z"}},
	}

	// Run tie-heavy tests multiple times to catch nondeterminism from map iteration
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			for rep := 0; rep < 20; rep++ {
				got := SortedKeysByValue(tc.in)

				// exact match if want is provided
				if len(got) != len(tc.want) {
					t.Fatalf("len mismatch: got=%v want=%v", got, tc.want)
				}
				for i := range got {
					if got[i] != tc.want[i] {
						t.Fatalf("rep=%d mismatch at %d: got=%v want=%v", rep, i, got, tc.want)
					}
				}

				// property checks: contains exactly keys in map
				if tc.in != nil {
					seen := make(map[string]bool, len(got))
					for _, k := range got {
						if seen[k] {
							t.Fatalf("rep=%d duplicate key %q in result: %v", rep, k, got)
						}
						seen[k] = true
						if _, ok := tc.in[k]; !ok {
							t.Fatalf("rep=%d key %q not in input map", rep, k)
						}
					}
					if len(seen) != len(tc.in) {
						t.Fatalf("rep=%d missing keys: got=%v in=%v", rep, got, tc.in)
					}

					// property checks: sorted by desc value, then asc key
					for i := 1; i < len(got); i++ {
						prev, cur := got[i-1], got[i]
						pv, cv := tc.in[prev], tc.in[cur]
						if pv < cv {
							t.Fatalf("rep=%d not sorted by value desc: %v (%d) before %v (%d)", rep, prev, pv, cur, cv)
						}
						if pv == cv && prev > cur {
							t.Fatalf("rep=%d not sorted by key asc for tie: %q before %q", rep, prev, cur)
						}
					}
				}
			}
		})
	}
}

func TestEX04_GetInt(t *testing.T) {
	cases := []struct {
		name string
		m    map[string]string
		key  string
		ok   bool
		want int
		err  bool
	}{
		{"valid_simple", map[string]string{"a": "10"}, "a", true, 10, false},
		{"valid_negaive", map[string]string{"a": "-5"}, "a", true, -5, false},
		{"valid_leading_zeros", map[string]string{"a": "001"}, "a", true, 1, false},

		// strict: reject plus sign, spaces, empty
		{"invalid_plus", map[string]string{"a": "+1"}, "a", true, 1, false},
		{"invalid_spaces_inside", map[string]string{"a": "1 2"}, "a", true, 0, true},
		{"invalid_empty", map[string]string{"a": ""}, "a", true, 0, true},
		{"invalid_alpha", map[string]string{"x": "2a"}, "x", true, 0, true},
		{"invalid_trailing_space", map[string]string{"x": "7 "}, "x", true, 0, true},
		{"invalid_leading_space", map[string]string{"x": " 7"}, "x", true, 0, true},

		// missing key is not an error
		{"missing_key", map[string]string{"b": "7"}, "a", false, 0, false},

		// nil map treated as empty (no panic, no error)
		{"nil_map_missing", nil, "a", false, 0, false},

		// big number (still fits in int on 64-bit; keep it as string parse check)
		{"valid_large", map[string]string{"x": "2147483647"}, "x", true, 2147483647, false},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, ok, err := GetInt(tc.m, tc.key)

			if tc.err {
				if err == nil {
					t.Fatalf("expected error")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}

			if ok != tc.ok {
				t.Fatalf("ok mismatch: got %v want %v", ok, tc.ok)
			}

			if !tc.err && ok && got != tc.want {
				t.Fatalf("got %d want %d", got, tc.want)
			}
		})
	}
}
