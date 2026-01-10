# Go Buddy — BASIC Language Fundamentals Exercises

This repository provides challenging, test-driven Go exercises for senior engineers, focused strictly on BASIC topics (1–6) from the roadmap: core syntax, control flow, functions, collections, pointers, and errors.

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
│  ├─ ex02_*.go and ex02_*_test.go
│  └─ ... (10+ exercises total)
├─ intermediate/ (reserved)
└─ advanced/ (reserved)
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
- Run a single exercise (by prefix like ex01):
  - `go run ./cmd/gobuddy test beginner ex01`
  - or alias: `go run ./cmd/gobuddy beginner ex01`
  - `go run ./cmd/gobuddy test intermediate ex03`
  - or alias: `go run ./cmd/gobuddy intermediate ex03`

Notes:
- The repository compiles out of the box. Tests are designed to fail until you implement the TODOs.
- Only BASIC topics are required for solving the exercises. No goroutines, channels, generics, or reflection.
