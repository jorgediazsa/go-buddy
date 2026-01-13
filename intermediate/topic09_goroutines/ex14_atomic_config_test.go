package topic09_goroutines

import (
	"reflect"
	"sync"
	"testing"
)

func TestEX09_Ex14_ConfigManager(t *testing.T) {
	initial := Config{Endpoints: []string{"a"}, Timeout: 10}
	cm := NewConfigManager(initial)

	var wg sync.WaitGroup
	const iterations = 1000

	// Concurrent readers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				_ = cm.Get()
			}
		}()
	}

	// Concurrent writers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				cm.Update(Config{Endpoints: []string{"updated"}, Timeout: j})
			}
		}(i)
	}

	wg.Wait()

	final := cm.Get()
	if !reflect.DeepEqual(final.Endpoints, []string{"updated"}) {
		t.Errorf("Config was not updated correctly, got %v", final.Endpoints)
	}
}
