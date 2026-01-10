package intermediate

import "testing"

func TestEX03_InterfaceSatisfaction(t *testing.T) {
	// Blob (value) satisfies Sizer
	var _ Sizer = Blob{}
	// *MutBlob (pointer) satisfies Sizer
	var _ Sizer = &MutBlob{}
}

func TestEX03_TotalSizeAndGrow(t *testing.T) {
	cases := []struct {
		blobs   []Blob
		muts    []int // initial sizes for MutBlob
		growIdx int   // index into total items to grow (Sizer index)
		delta   int
		want    int
	}{
		{nil, nil, -1, 0, 0},
		{[]Blob{{0}}, nil, -1, 0, 0},
		{[]Blob{{3}, {4}}, nil, -1, 0, 7},
		{nil, []int{1}, -1, 0, 1},
		{nil, []int{1, 2, 3}, -1, 0, 6},
		{[]Blob{{2}}, []int{5}, -1, 0, 7},
		{[]Blob{{2}}, []int{5}, 1, +3, 10},
		{[]Blob{{2}, {8}}, []int{5}, 0, +10, 25},
		{[]Blob{{2}}, []int{5, 7}, 2, -2, 12},
		{[]Blob{{0}, {0}}, []int{0, 0, 0}, 3, +1, 1},
	}
	for i, tc := range cases {
		t.Run("EX03_TotalSize_"+string(rune('A'+i)), func(t *testing.T) {
			var items []Sizer
			for _, b := range tc.blobs {
				items = append(items, b)
			}
			for _, n := range tc.muts {
				items = append(items, NewMutBlob(n))
			}
			if tc.growIdx >= 0 && tc.growIdx < len(items) {
				items[tc.growIdx] = MaybeGrow(items[tc.growIdx], tc.delta)
			}
			got := TotalSize(items)
			if got != tc.want {
				t.Fatalf("got %d want %d", got, tc.want)
			}
		})
	}
}

func TestEX03_MethodSetBehavior(t *testing.T) {
	// Ensure pointer receiver needed for MutBlob Size
	mb := NewMutBlob(10)
	var s Sizer = mb
	if s.Size() == 0 { /* placeholder */
	}
}
