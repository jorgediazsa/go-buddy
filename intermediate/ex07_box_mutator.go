package intermediate

/*
Exercise EX07 — Pointer vs value receivers with small interfaces

Why this matters
- Method sets differ between values and pointers. Interfaces often require pointer receivers for mutating behavior.

Requirements
- Define Box with private field v int.
- Methods:
  - func (b Box) Value() int
  - func (b *Box) Set(n int)
  - func (b Box) WithAdded(d int) Box // return a new Box with value+ d
  - func (b *Box) Apply(fn func(int) int) // mutate in-place using fn
- Small interface:
  - type Applier interface { Apply(func(int) int) }
- Helpers:
  - NewBox(n int) Box
  - TryApply(x any, fn func(int) int) bool // if x implements Applier, call Apply(fn) and return true

Constraints and pitfalls
- Ensure pointer receiver for Apply to emphasize method set behavior.
- Avoid panics on nil receivers (Apply should be safe on nil *Box: no-op, return).

Tricky edge cases
- TryApply on Box value (not pointer) should return false due to method set mismatch.
*/

// Applier is a small interface for in-place mutation.
type Applier interface{ Apply(func(int) int) }

// Box holds an integer value.
type Box struct{ v int }

func NewBox(n int) Box { // TODO
	return Box{}
}

func (b Box) Value() int { // TODO
	return 0
}

func (b *Box) Set(n int) { // TODO
}

func (b Box) WithAdded(d int) Box { // TODO
	return Box{}
}

func (b *Box) Apply(fn func(int) int) { // TODO
}

// TryApply attempts to apply fn if x implements Applier.
func TryApply(x any, fn func(int) int) bool { // TODO
	return false
}
