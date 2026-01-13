package topic09_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestEX09_Ex13_RWCache(t *testing.T) {
	c := NewRWCache()
	var wg sync.WaitGroup

	// Writers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			c.Set(fmt.Sprintf("key-%d", id), "val")
		}(i)
	}

	// Readers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			c.Get(fmt.Sprintf("key-%d", id%10))
		}(i)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		val, ok := c.Get(fmt.Sprintf("key-%d", i))
		if !ok || val != "val" {
			t.Errorf("Expected key-%d to be 'val', got %s (ok=%v)", i, val, ok)
		}
	}
}
