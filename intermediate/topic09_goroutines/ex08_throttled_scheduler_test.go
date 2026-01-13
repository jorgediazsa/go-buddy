package topic09_goroutines

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestEX09_Ex08_ThrottledScheduler_Basics(t *testing.T) {
	ts := NewThrottledScheduler(2)
	var running int32
	var maxSeen int32
	tasks := 12
	for i := 0; i < tasks; i++ {
		ok := ts.Submit(func() {
			n := atomic.AddInt32(&running, 1)
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
			atomic.AddInt32(&running, -1)
		})
		if !ok {
			t.Fatalf("submit failed before close")
		}
	}
	ts.Close()
	if atomic.LoadInt32(&running) != 0 {
		t.Fatalf("running not zero after close")
	}
	if atomic.LoadInt32(&maxSeen) > 2 {
		t.Fatalf("maxSeen exceeded: %d", maxSeen)
	}
	if ts.Submit(func() {}) {
		t.Fatalf("Submit should be false after Close")
	}
}
