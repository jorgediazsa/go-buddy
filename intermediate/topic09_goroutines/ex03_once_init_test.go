package topic09_goroutines

import (
	"errors"
	"sync"
	"testing"
)

func TestEX09_Ex03_OnceInit_SingleExecutionSuccess(t *testing.T) {
	var o OnceInit
	var initCount int
	fn := func() error { initCount++; return nil }
	const gor = 64
	var wg sync.WaitGroup
	wg.Add(gor)
	for i := 0; i < gor; i++ {
		go func() { defer wg.Done(); _ = o.Do(fn) }()
	}
	wg.Wait()
	if initCount != 1 {
		t.Fatalf("initializer called %d times, want 1", initCount)
	}
	if !o.Done() {
		t.Fatalf("Done should be true after first Do")
	}
}

func TestEX09_Ex03_OnceInit_ErrorSticks(t *testing.T) {
	var o OnceInit
	var initCount int
	errSentinel := errors.New("boom")
	fnErr := func() error { initCount++; return errSentinel }
	// First attempt fails
	if err := o.Do(fnErr); err == nil || !errors.Is(err, errSentinel) {
		t.Fatalf("expected sentinel error, got %v", err)
	}
	if initCount != 1 {
		t.Fatalf("init count=%d want 1", initCount)
	}
	// Later attempt with success must not run and must return same error
	fnOK := func() error { initCount++; return nil }
	if err := o.Do(fnOK); !errors.Is(err, errSentinel) {
		t.Fatalf("expected sticky error, got %v", err)
	}
	if initCount != 1 {
		t.Fatalf("initializer ran again: count=%d", initCount)
	}
}
