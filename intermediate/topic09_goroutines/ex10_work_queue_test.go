package topic09_goroutines

import (
	"sync"
	"testing"
	"time"
)

func TestEX09_Ex10_WorkQueue_EnqueueDequeueClose(t *testing.T) {
	wq := NewWorkQueue()
	var got []int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			v, ok := wq.Dequeue()
			if !ok {
				return
			}
			got = append(got, v)
		}
	}()
	for i := 0; i < 10; i++ {
		if !wq.Enqueue(i) {
			t.Fatalf("enqueue failed before close")
		}
	}
	wq.Close()
	wg.Wait()
	if len(got) != 10 {
		t.Fatalf("got %d items", len(got))
	}
	// Enqueue after close must fail
	if wq.Enqueue(99) {
		t.Fatalf("enqueue succeeded after close")
	}
}

func TestEX09_Ex10_WorkQueue_CloseWakesWaiters(t *testing.T) {
	wq := NewWorkQueue()
	var wg sync.WaitGroup
	waiters := 5
	wg.Add(waiters)
	for i := 0; i < waiters; i++ {
		go func() {
			defer wg.Done()
			_, ok := wq.Dequeue()
			if ok {
				t.Errorf("expected ok=false on close")
			}
		}()
	}
	time.Sleep(2 * time.Millisecond)
	wq.Close()
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(50 * time.Millisecond):
		t.Fatalf("waiters not woken by Close")
	}
}
