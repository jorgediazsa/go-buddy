# Go Buddy — Go Mastery Exercises

This repository provides challenging, test-driven Go exercises for senior engineers, covering basic and intermediate topics from the roadmap.

Structure:

```text
.
├─ go.mod
├─ README.md
├─ cmd/
│  └─ gobuddy/
│     └─ main.go
├─ beginner/
│  ├─ ex01_*.go and ex01_*_test.go
│  └─ ...
├─ intermediate/
│  ├─ topic09_goroutines/
│  │  ├─ ex01_*.go and ex01_*_test.go
│  │  └─ ...
│  └─ ...
├─ advanced/ (reserved)
```

How to use:
- Read the comment block at the top of each exercise file in beginner/ to understand goals, constraints, and edge cases.
- Implement the TODOs in the exercise .go files.
- Run tests via the CLI or directly with `go test`.

CLI (cross‑platform):
- List:
  - `go run ./cmd/gobuddy list`
- Run all tests:
  - `go run ./cmd/gobuddy test`
- Run a level:
  - `go run ./cmd/gobuddy test beginner`
  - or alias: `go run ./cmd/gobuddy beginner`
  - `go run ./cmd/gobuddy test intermediate`
  - or alias: `go run ./cmd/gobuddy intermediate`
  - Run a topic (subfolder) under intermediate (topics 08–14):
    - `go run ./cmd/gobuddy test intermediate/topic10_channels`
    - or alias: `go run ./cmd/gobuddy intermediate/topic10_channels`
  - Topic 09 (Goroutines) examples:
    - Run all topic 09 tests: `go run ./cmd/gobuddy intermediate/topic09_goroutines`
    - Run a single exercise: `go run ./cmd/gobuddy intermediate/topic09_goroutines ex07`
- Run a single exercise (by prefix like ex01):
  - `go run ./cmd/gobuddy test beginner ex01`
  - or alias: `go run ./cmd/gobuddy beginner ex01`
  - `go run ./cmd/gobuddy test intermediate ex03`
  - or alias: `go run ./cmd/gobuddy intermediate ex03`
  - Run a single exercise inside a topic folder:
    - `go run ./cmd/gobuddy test intermediate/topic10_channels ex07`
    - or alias: `go run ./cmd/gobuddy intermediate/topic10_channels ex07`

Notes:
- The repository compiles out of the box. Tests are designed to fail until you implement the TODOs.
- Exercises use standard library only.
- For intermediate/advanced topics, we focus on senior-level challenges, including concurrency, memory safety, and idiomatic patterns.
