package topic09_goroutines

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestEX09_Ex05_BoundedRunner_LimitAndClose(t *testing.T) {
	br := NewBoundedRunner(3)
	var maxSeen int32
	var inFlight int32
	total := 20
	done := make(chan struct{})
	for i := 0; i < total; i++ {
		ok := br.Run(func() {
			n := atomic.AddInt32(&inFlight, 1)
			for {
				m := atomic.LoadInt32(&maxSeen)
				if n > m && atomic.CompareAndSwapInt32(&maxSeen, m, n) {
					break
				}
				if n <= m {
					break
				}
			}
			time.Sleep(2 * time.Millisecond)
			atomic.AddInt32(&inFlight, -1)
		})
		if !ok {
			t.Fatalf("Run returned false before Close")
		}
	}
	br.Close()
	close(done)
	if got := atomic.LoadInt32(&maxSeen); got > 3 {
		t.Fatalf("max concurrency %d exceeds limit", got)
	}
	if br.InFlight() != 0 {
		t.Fatalf("inflight after close: %d", br.InFlight())
	}
	// further Run must be rejected
	if br.Run(func() {}) {
		t.Fatalf("Run should return false after Close")
	}
}

func TestEX09_Ex05_BoundedRunner_LimitOne_NoDeadlock(t *testing.T) {
	br := NewBoundedRunner(1)
	started := int32(0)
	ok := br.Run(func() {
		atomic.AddInt32(&started, 1)
		time.Sleep(1 * time.Millisecond)
	})
	if !ok {
		t.Fatalf("Run failed")
	}
	br.Close()
	if s := atomic.LoadInt32(&started); s != 1 {
		t.Fatalf("started=%d want 1", s)
	}
}
