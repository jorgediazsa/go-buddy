package topic09_goroutines

import (
	"testing"
	"time"
)

func TestEX09_Ex01_WorkerStartStopIdempotent(t *testing.T) {
	w := &Worker{}
	// Stop before start: should not panic
	w.Stop()
	w.Start()
	w.Start() // idempotent
	time.Sleep(5 * time.Millisecond)
	c1 := w.Count()
	if c1 == 0 {
		t.Fatalf("expected count to increase after start")
	}
	w.Stop()
	// After Stop, count should remain stable (allow small delay)
	time.Sleep(2 * time.Millisecond)
	c2 := w.Count()
	if c2 != c1 {
		t.Fatalf("count changed after stop: %d -> %d", c1, c2)
	}
	// Stop again is safe
	w.Stop()
}

func TestEX09_Ex01_WorkerNoLeakOnStop(t *testing.T) {
	w := &Worker{}
	w.Start()
	time.Sleep(2 * time.Millisecond)
	w.Stop()
	// If Stop returns, goroutine must be gone. We can only check that another Stop returns quickly.
	done := make(chan struct{})
	go func() { w.Stop(); close(done) }()
	select {
	case <-done:
		// ok
	case <-time.After(50 * time.Millisecond):
		t.Fatalf("stop not idempotent or leaked goroutine")
	}
}
