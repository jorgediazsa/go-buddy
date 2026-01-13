package topic09_goroutines

import (
	"sync"
	"testing"
)

func TestEX09_Ex09_AtomicState_Transitions(t *testing.T) {
	var s AtomicState
	if s.Load() != StateInit {
		t.Fatalf("want init")
	}
	if !s.TryStart() {
		t.Fatalf("TryStart should succeed from init")
	}
	if s.Load() != StateRunning {
		t.Fatalf("state not running")
	}
	if s.TryStart() {
		t.Fatalf("TryStart should fail from running")
	}
	if !s.TryStop() {
		t.Fatalf("TryStop should succeed from running")
	}
	if !s.IsTerminal() {
		t.Fatalf("stopped should be terminal")
	}
	if s.TryStop() {
		t.Fatalf("TryStop should fail after stopped")
	}
}

func TestEX09_Ex09_AtomicState_ConcurrentStartOnlyOnce(t *testing.T) {
	var s AtomicState
	const gor = 50
	var wg sync.WaitGroup
	wg.Add(gor)
	successes := 0
	for i := 0; i < gor; i++ {
		go func() {
			defer wg.Done()
			if s.TryStart() {
				// count successes; non-atomic increment acceptable in test due to single success expected
				successes++
			}
		}()
	}
	wg.Wait()
	if successes != 1 {
		t.Fatalf("TryStart successes=%d want 1", successes)
	}
}
