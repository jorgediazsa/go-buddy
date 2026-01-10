# Go (Golang) Learning Roadmap

## 🟢 BASIC — Language Fundamentals

### 1. Core Syntax
- Go program structure (`package main`, `func main`)
- Variable declaration (`var`, `:=`)
- Basic types: `int`, `float64`, `string`, `bool`
- Constants (`const`)
- Type conversions

### 2. Control Flow
- `if`, `else if`, `else`
- `switch` (expression-less switch)
- `for` (the only loop in Go)
- `for range` (iterate over slices, maps, strings)

### 3. Functions
- Definition and multiple return values
- Named return parameters
- Variadic functions
- Anonymous functions
- Closures

### 4. Standard Collections
- Arrays
- Slices (length / capacity, `make`, `append`, slicing)
- Maps (declaration, assignment, safe lookup)
- Structs (definition and basic usage)

### 5. Pointers
- `&`, `*`
- Pointer to struct
- Mutability vs value copy

### 6. Errors
- Error as a value (`error`)
- `errors.New`, `fmt.Errorf`
- Common pattern: `value, err := ...`


## 🟡 INTERMEDIATE — Idiomatic Go

### 7. Methods and Composition
- Methods with value vs pointer receivers
- Embedding (composition → Go “inheritance”)
- Implicit interfaces
- Small interfaces (single-method)
- `io.Reader` / `io.Writer` as fundamental patterns

### 8. Packages and Modules
- Code organization
- `go mod init`, `go mod tidy`
- Workspaces (`go work`)
- Import cycles and how to avoid them

### 9. Goroutines
- `go func() {}`
- Concurrency vs parallelism
- Goroutine leaks
- Conceptual data races

### 10. Channels
- Creation (`make(chan T)`)
- Unbuffered vs buffered
- Send and receive (`ch <- v`, `<-ch`)
- Directional channels (`chan<-`, `<-chan`)
- `select` statement
- Channel closing and `range ch`

### 11. Context
- `context.Background`, `context.WithCancel`, `context.WithTimeout`
- Cooperative cancellation
- Context propagation
- Proper use in HTTP handlers

### 12. Idiomatic Error Handling
- `errors.Is`, `errors.As`, wrapping
- Custom error types
- Sentinel errors
- “Early return” pattern

### 13. Testing
- `testing` package
- Subtests (`t.Run`)
- Table-driven tests
- Benchmarks (`go test -bench`)
- Mocking with interfaces

### 14. HTTP and JSON
- `net/http`: handler, server, basic middleware
- JSON encoding/decoding (`json.Marshal`, `json.Unmarshal`)
- Tags (`json:"name,omitempty"`)


## 🔴 ADVANCED — Strong Interview Level

### 15. Advanced Concurrency
- Worker pool pattern
- Fan-in / fan-out
- Pipelines
- Semaphore pattern using channels
- Cascading cancellation with `context`

### 16. Memory Model
- Escape analysis (`go build -gcflags="-m"`)
- Stack vs heap
- Inlining
- Garbage collector overview

### 17. Synchronization Primitives
- `sync.WaitGroup`
- `sync.Mutex`
- `sync.RWMutex`
- `sync.Once`
- `sync.Map`
- `sync.Cond`

### 18. Generics
- Type parameters
- Constraints
- Generic slice/map helpers
- Interfaces with generics

### 19. Advanced Interfaces
- Deep duck typing
- Empty interface (`any`)
- Reflection (`reflect`)
- Proper use vs abuse

### 20. Profiling & Observability
- `pprof` (CPU, memory)
- `trace`
- `go tool pprof`
- Exposing `/debug/pprof` in a server

### 21. Advanced Error Handling
- Retried operations
- Backoff strategies
- Error aggregation
- Custom error stacks

### 22. Go in Production
- Project layout (Clean Architecture / Standard Layout)
- Graceful shutdown with `http.Server`
- Configuration (files, env vars)
- Structured logging (`zap`, `zerolog`)
- Dependency injection (manual, `uber/fx`, `wire`)

### 23. Go + Infrastructure
- Cross-platform compilation
- Docker multi-stage builds
- Health checks and readiness probes
- Workers and background jobs
