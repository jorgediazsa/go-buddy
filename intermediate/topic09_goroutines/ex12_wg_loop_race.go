package topic09_goroutines

/*
Title: EX12 — WaitGroup Loop Pitfalls

Why this matters
- Senior engineers must recognize common patterns that lead to data races and deadlocks. Capturing loop variables and misplacing Add() are classic mistakes.

Requirements
- Implement RunTasks(tasks []string, process func(string)) that:
  - Starts a goroutine for each task.
  - Waits for all tasks to finish.
  - Passes the CORRECT string to each process call (avoid the loop variable capture trap).
  - Correctly manages the WaitGroup (avoid Add() inside the goroutine race).

Constraints and pitfalls
- No data races according to `go test -race`.
- Must not return until all tasks are done.

Tricky edge case
- Empty tasks slice: should return immediately without panic.
*/

import "sync"

func RunTasks(tasks []string, process func(string)) { // TODO: implement correctly
	var wg sync.WaitGroup
	for _, t := range tasks {
		_ = t // use t or shadow it
		// ... implement goroutine start and wait
	}
	wg.Wait()
}
