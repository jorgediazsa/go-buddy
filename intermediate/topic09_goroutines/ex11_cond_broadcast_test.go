package topic09_goroutines

import (
	"sync"
	"testing"
	"time"
)

func TestEX09_Ex11_BroadcastNode(t *testing.T) {
	bn := NewBroadcastNode()
	var wg sync.WaitGroup

	const numWaiters = 5
	started := make(chan struct{}, numWaiters)
	results := make(chan int, numWaiters)

	for i := 0; i < numWaiters; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			started <- struct{}{}
			bn.Wait()
			results <- id
		}(i)
	}

	// Wait for all to start
	for i := 0; i < numWaiters; i++ {
		<-started
	}

	// Give them a moment to actually call Wait()
	time.Sleep(50 * time.Millisecond)

	if len(results) != 0 {
		t.Fatal("Wait() did not block; it should wait for SignalAll()")
	}

	bn.SignalAll()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// success
	case <-time.After(2 * time.Second):
		t.Fatal("Timeout: waiters did not wake up (leak or deadlock)")
	}

	if len(results) != numWaiters {
		t.Errorf("Expected %d results, got %d", numWaiters, len(results))
	}
}

func TestEX09_Ex11_BroadcastNode_NoSignalBuffering(t *testing.T) {
	bn := NewBroadcastNode()

	// Signal when no one is waiting
	bn.SignalAll()

	// Now wait - it should block (use a timeout to verify)
	woke := make(chan struct{})
	go func() {
		bn.Wait()
		close(woke)
	}()

	select {
	case <-woke:
		t.Error("Wait() should have blocked because SignalAll happened before Wait()")
	case <-time.After(100 * time.Millisecond):
		// Correct: it blocked. Now signal to clean up.
		bn.SignalAll()
		<-woke
	}
}
