package intermediate

import "testing"

func TestEX07_BoxValueSetWithAdded(t *testing.T) {
	cases := []struct {
		start int
		set   int
		add   int
		want  int
	}{
		{0, 0, 0, 0},
		{1, 0, 2, 3},
		{5, 10, -3, 7},
		{-2, -2, 5, 3},
		{100, 100, 0, 100},
		{7, 8, 9, 17},
		{9, 9, -9, 0},
		{-10, -5, -5, -10},
		{1, 2, 3, 5},
		{42, 42, 1, 43},
	}
	for i, tc := range cases {
		t.Run("EX07_BoxBasic_"+string(rune('A'+i)), func(t *testing.T) {
			b := NewBox(tc.start)
			if b.Value() != tc.start {
				t.Fatalf("start=%d got=%d", tc.start, b.Value())
			}
			b.Set(tc.set)
			nb := b.WithAdded(tc.add)
			if nb.Value() != tc.want {
				t.Fatalf("WithAdded=%d want %d", nb.Value(), tc.want)
			}
			if b.Value() != tc.set {
				t.Fatalf("original mutated to %d", b.Value())
			}
		})
	}
}

func TestEX07_ApplyAndTryApply(t *testing.T) {
	inc := func(x int) int { return x + 1 }
	dbl := func(x int) int { return x * 2 }
	b := NewBox(10)
	// Box value should NOT implement Applier (pointer receiver)
	if TryApply(b, inc) {
		t.Fatalf("TryApply on value should be false")
	}
	// Pointer should implement Applier
	if !TryApply(&b, inc) {
		t.Fatalf("TryApply on pointer should be true")
	}
	if b.Value() == 10 {
		t.Fatalf("Apply had no effect")
	}
	// Nil receiver safety
	var pb *Box
	if TryApply(pb, dbl) {
		t.Fatalf("TryApply on nil pointer should be false")
	}
}
