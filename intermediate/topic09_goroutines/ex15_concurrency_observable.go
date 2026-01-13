package topic09_goroutines

/*
Title: EX15 — Concurrency Observable: observing interleaving

Why this matters
- Concurrency is about structure, parallelism is about execution. Even with GOMAXPROCS=1, Go can interleave multiple goroutines. Understanding how to yield or perform I/O to allow interleaving is key to responsive concurrent programs.

Requirements
- Implement RunInterleaved(w io.Writer, iterations int) that:
  - Starts two goroutines: one printing "A" and another printing "B", each `iterations` times.
  - Uses `runtime.Gosched()` after each print to yield the processor.
  - Ensures the output in `w` is correctly written (use a mutex for the writer).
  - Waits for both goroutines to finish.

Constraints and pitfalls
- Without `runtime.Gosched()` or blocking operations, one goroutine might starve the other on a single-core machine (or GOMAXPROCS=1).
- io.Writer is not thread-safe by default.

Tricky edge case
- iterations = 0: should do nothing and return.
*/

import (
	"io"
)

func RunInterleaved(w io.Writer, iterations int) { // TODO: implement
}
