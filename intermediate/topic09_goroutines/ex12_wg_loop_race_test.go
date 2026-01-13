package topic09_goroutines

import (
	"sync"
	"testing"
)

func TestEX09_Ex12_RunTasks(t *testing.T) {
	tasks := []string{"a", "b", "c", "d", "e"}
	processed := make(map[string]int)
	var mu sync.Mutex

	RunTasks(tasks, func(s string) {
		mu.Lock()
		processed[s]++
		mu.Unlock()
	})

	if len(processed) != len(tasks) {
		t.Errorf("Expected %d unique tasks, got %d. Possible loop variable capture race.", len(tasks), len(processed))
	}

	for _, k := range tasks {
		if processed[k] != 1 {
			t.Errorf("Task %s was processed %d times, expected 1", k, processed[k])
		}
	}
}

func TestEX09_Ex12_RunTasks_Empty(t *testing.T) {
	// Should not panic or hang
	RunTasks(nil, func(s string) {})
}
