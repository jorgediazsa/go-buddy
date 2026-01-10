package beginner

import "testing"

func TestEX07_FilterInPlace(t *testing.T) {
	even := func(x int) bool { return x%2 == 0 }
	gt2 := func(x int) bool { return x > 2 }
	cases := []struct {
		in   []int
		keep func(int) bool
		want []int
	}{
		{nil, even, nil},
		{[]int{}, even, []int{}},
		{[]int{1}, even, []int{}},
		{[]int{2}, even, []int{2}},
		{[]int{1, 2, 3, 4, 5, 6}, even, []int{2, 4, 6}},
		{[]int{1, 2, 3, 4, 5, 6}, gt2, []int{3, 4, 5, 6}},
		{[]int{3, 3, 3}, gt2, []int{3, 3, 3}},
		{[]int{1, 1, 1}, gt2, []int{}},
	}
	for i, tc := range cases {
		t.Run("EX07_FilterInPlace_"+string(rune('A'+i)), func(t *testing.T) {
			in := append([]int(nil), tc.in...)
			n := FilterInPlace(in, tc.keep)
			if n != len(tc.want) {
				t.Fatalf("new len %d want %d", n, len(tc.want))
			}
			for i := 0; i < n; i++ {
				if in[i] != tc.want[i] {
					t.Fatalf("mismatch at %d: %v vs %v", i, in[:n], tc.want)
				}
			}
		})
	}
}

func TestEX07_RangeSums(t *testing.T) {
	cases := []struct {
		n    int
		even int
		odd  int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{2, 2, 1}, // evens: 0+2
		{3, 2, 4}, // odds: 1+3
		{4, 6, 4}, // evens: 0+2+4
		{10, 30, 25},
		{-1, 0, 0},
		{11, 36, 36},
	}
	for i, tc := range cases {
		t.Run("EX07_RangeSums_"+string(rune('A'+i)), func(t *testing.T) {
			e, o := RangeSums(tc.n)
			if e != tc.even || o != tc.odd {
				t.Fatalf("got (%d,%d) want (%d,%d)", e, o, tc.even, tc.odd)
			}
		})
	}
}

func TestEX07_CategorizeTemps(t *testing.T) {
	cases := []struct {
		in              []int
		cold, mild, hot int
	}{
		{nil, 0, 0, 0},
		{[]int{}, 0, 0, 0},
		{[]int{9}, 1, 0, 0},
		{[]int{10}, 0, 1, 0},
		{[]int{24}, 0, 1, 0},
		{[]int{25}, 0, 0, 1},
		{[]int{5, 10, 15, 24, 25, 30}, 1, 3, 2},
		{[]int{-5, 0, 100}, 2, 0, 1},
	}
	for i, tc := range cases {
		t.Run("EX07_CategorizeTemps_"+string(rune('A'+i)), func(t *testing.T) {
			got := CategorizeTemps(append([]int(nil), tc.in...))
			if got == nil {
				t.Fatalf("nil map")
			}
			if got["cold"] != tc.cold || got["mild"] != tc.mild || got["hot"] != tc.hot {
				t.Fatalf("got %v want cold=%d mild=%d hot=%d", got, tc.cold, tc.mild, tc.hot)
			}
		})
	}
}
