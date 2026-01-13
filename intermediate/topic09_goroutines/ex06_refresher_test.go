package topic09_goroutines

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestEX09_Ex06_BackgroundRefresher_StartStopRestart(t *testing.T) {
	var b BackgroundRefresher
	var calls int32
	refresh := func() error { atomic.AddInt32(&calls, 1); return nil }
	if err := b.Start(2*time.Millisecond, refresh); err != nil {
		t.Fatalf("start err: %v", err)
	}
	time.Sleep(6 * time.Millisecond)
	b.Stop()
	c1 := atomic.LoadInt32(&calls)
	if c1 == 0 {
		t.Fatalf("expected some refresh calls before stop")
	}
	// ensure stopped
	time.Sleep(4 * time.Millisecond)
	c2 := atomic.LoadInt32(&calls)
	if c2 != c1 {
		t.Fatalf("ticks increased after stop: %d -> %d", c1, c2)
	}
	// restart
	if err := b.Start(2*time.Millisecond, refresh); err != nil {
		t.Fatalf("restart err: %v", err)
	}
	time.Sleep(5 * time.Millisecond)
	b.Stop()
	if atomic.LoadInt64(&b.ticks) == 0 {
		t.Fatalf("Ticks should record successful refreshes")
	}
}

func TestEX09_Ex06_BackgroundRefresher_InvalidParams(t *testing.T) {
	var b BackgroundRefresher
	if err := b.Start(0, func() error { return nil }); err == nil {
		t.Fatalf("expected error on zero period")
	}
	if err := b.Start(1*time.Millisecond, nil); err == nil {
		t.Fatalf("expected error on nil refresh")
	}
}
