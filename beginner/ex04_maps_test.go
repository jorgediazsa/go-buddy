package beginner

import "testing"

func TestEX04_MergeSafe(t *testing.T) {
	cases := []struct {
		a, b map[string]int
		want map[string]int
	}{
		{nil, nil, map[string]int{}},
		{map[string]int{}, map[string]int{}, map[string]int{}},
		{map[string]int{"x": 1}, nil, map[string]int{"x": 1}},
		{nil, map[string]int{"y": 2}, map[string]int{"y": 2}},
		{map[string]int{"k": 1, "z": 9}, map[string]int{"k": 3}, map[string]int{"k": 3, "z": 9}},
		{map[string]int{"a": 1}, map[string]int{"b": 2}, map[string]int{"a": 1, "b": 2}},
		{map[string]int{"a": 1, "b": 2}, map[string]int{"b": 2}, map[string]int{"a": 1, "b": 2}},
		{map[string]int{"a": 1}, map[string]int{"a": 1}, map[string]int{"a": 1}},
	}
	for i, tc := range cases {
		t.Run("EX04_MergeSafe_"+string(rune('A'+i)), func(t *testing.T) {
			got := MergeSafe(tc.a, tc.b)
			if got == nil {
				t.Fatalf("got nil map")
			}
			// inputs unmodified
			if tc.a != nil && len(tc.a) > 0 {
				if tc.a["k"] == 3 && (tc.b == nil || tc.b["k"] == 0) {
					t.Fatalf("input map a mutated")
				}
			}
			if len(got) != len(tc.want) {
				t.Fatalf("len %d want %d: %v", len(got), len(tc.want), got)
			}
			for k, v := range tc.want {
				if got[k] != v {
					t.Fatalf("key %q: got %d want %d", k, got[k], v)
				}
			}
		})
	}
}

func TestEX04_SortedKeysByValue(t *testing.T) {
	cases := []struct {
		in   map[string]int
		want []string
	}{
		{nil, []string{}},
		{map[string]int{}, []string{}},
		{map[string]int{"a": 1}, []string{"a"}},
		{map[string]int{"a": 1, "b": 2, "c": 3}, []string{"c", "b", "a"}},
		{map[string]int{"a": 2, "b": 2, "c": 1}, []string{"a", "b", "c"}},
		{map[string]int{"b": 2, "a": 2, "c": 2}, []string{"a", "b", "c"}},
		{map[string]int{"x": -1, "y": -2}, []string{"x", "y"}},
		{map[string]int{"k": 0, "j": 0, "a": 0}, []string{"a", "j", "k"}},
	}
	for i, tc := range cases {
		t.Run("EX04_SortedKeysByValue_"+string(rune('A'+i)), func(t *testing.T) {
			got := SortedKeysByValue(tc.in)
			if len(got) != len(tc.want) {
				t.Fatalf("len mismatch: %v vs %v", got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("mismatch at %d: %v vs %v", i, got, tc.want)
				}
			}
		})
	}
}

func TestEX04_GetInt(t *testing.T) {
	cases := []struct {
		m    map[string]string
		key  string
		ok   bool
		want int
		err  bool
	}{
		{map[string]string{"a": "10"}, "a", true, 10, false},
		{map[string]string{"a": "-5"}, "a", true, -5, false},
		{map[string]string{"a": "001"}, "a", true, 1, false},
		{map[string]string{"a": "+1"}, "a", false, 0, true},
		{map[string]string{"a": "1 2"}, "a", false, 0, true},
		{map[string]string{"a": ""}, "a", false, 0, true},
		{map[string]string{"b": "7"}, "a", false, 0, false},
		{nil, "a", false, 0, true},
		{map[string]string{"x": "2147483647"}, "x", true, 2147483647, false},
		{map[string]string{"x": "2a"}, "x", false, 0, true},
	}
	for i, tc := range cases {
		t.Run("EX04_GetInt_"+string(rune('A'+i)), func(t *testing.T) {
			got, ok, err := GetInt(tc.m, tc.key)
			if tc.err && err == nil {
				t.Fatalf("expected error")
			}
			if !tc.err && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if ok != tc.ok {
				t.Fatalf("ok mismatch: %v vs %v", ok, tc.ok)
			}
			if !tc.err && ok && got != tc.want {
				t.Fatalf("got %d want %d", got, tc.want)
			}
		})
	}
}
