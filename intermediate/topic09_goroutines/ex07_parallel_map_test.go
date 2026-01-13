package topic09_goroutines

import (
	"testing"
	"time"
)

func TestEX09_Ex07_ParallelMapInts_OrderAndCorrectness(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fn := func(x int) int { return x * x }
	out := ParallelMapInts(inputs, fn, 3)
	if len(out) != len(inputs) {
		t.Fatalf("len mismatch")
	}
	for i, v := range inputs {
		if out[i] != v*v {
			t.Fatalf("at %d got %d want %d", i, out[i], v*v)
		}
	}
}

func TestEX09_Ex07_ParallelMapInts_Empty(t *testing.T) {
	out := ParallelMapInts(nil, func(x int) int { return x }, 4)
	if len(out) != 0 {
		t.Fatalf("want empty out")
	}
}

func TestEX09_Ex07_ParallelMapInts_SpeedupHint(t *testing.T) {
	// Not relying on CPU cores. We simulate work via small sleeps per item.
	inputs := make([]int, 12)
	for i := range inputs {
		inputs[i] = i
	}
	work := func(x int) int { time.Sleep(2 * time.Millisecond); return x + 1 }
	start := time.Now()
	_ = ParallelMapInts(inputs, work, 1)
	seqDur := time.Since(start)

	start = time.Now()
	_ = ParallelMapInts(inputs, work, 4)
	parDur := time.Since(start)

	// Should be noticeably faster, give a generous factor.
	if !(parDur < seqDur/2) {
		t.Fatalf("expected parallel duration %v < %v/2", parDur, seqDur)
	}
}
