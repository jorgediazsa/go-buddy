package topic09_goroutines

import (
	"sync"
	"testing"
)

func TestEX09_Ex02_SafeCounterConcurrent(t *testing.T) {
	var c SafeCounter
	const gor = 32
	const iters = 1000
	var wg sync.WaitGroup
	wg.Add(gor)
	for i := 0; i < gor; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < iters; j++ {
				c.Inc(1)
			}
		}()
	}
	wg.Wait()
	if got := c.Load(); got != gor*iters {
		t.Fatalf("got %d want %d", got, gor*iters)
	}
	c.Reset()
	if got := c.Load(); got != 0 {
		t.Fatalf("reset failed, got %d", got)
	}
}

func TestEX09_Ex02_NegativeIncrements(t *testing.T) {
	var c SafeCounter
	c.Inc(10)
	c.Inc(-3)
	if got := c.Load(); got != 7 {
		t.Fatalf("got %d want 7", got)
	}
}
