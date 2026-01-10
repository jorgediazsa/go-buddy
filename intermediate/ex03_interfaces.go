package intermediate

/*
Exercise EX03 — Implicit interfaces and method sets

Why this matters
- In Go, types satisfy interfaces implicitly. Pointer vs value receiver methods affect interface satisfaction.

Requirements
- Define small interface:
    type Sizer interface { Size() int }
- Types:
  - Blob { Len int } with func (b Blob) Size() int
  - MutBlob { n int } with func (b *MutBlob) Size() int and func (b *MutBlob) Add(d int)
- Functions:
  - TotalSize(items []Sizer) int
  - NewMutBlob(n int) *MutBlob
  - MaybeGrow(s Sizer, d int) Sizer // if s is *MutBlob, call Add(d) and return s; otherwise return s unchanged

Constraints and pitfalls
- Do not use reflection or generics; rely on small interfaces and type assertions.
- Ensure that Blob (value) satisfies Sizer, while MutBlob requires a pointer to satisfy Sizer.

Tricky edge cases
- MaybeGrow with negative d should still call Add.
- Passing MutBlob by value to MaybeGrow should not compile in user code; tests will use interface dynamics instead.
*/

// Sizer reports size.
type Sizer interface{ Size() int }

// Blob implements Sizer via value receiver.
type Blob struct{ Len int }

func (b Blob) Size() int { // TODO
	return 0
}

// MutBlob implements Sizer via pointer receiver.
type MutBlob struct{ n int }

func (b *MutBlob) Size() int { // TODO
	return 0
}

func (b *MutBlob) Add(d int) { // TODO
}

// NewMutBlob constructs a *MutBlob.
func NewMutBlob(n int) *MutBlob { // TODO
	return nil
}

// TotalSize sums sizes.
func TotalSize(items []Sizer) int { // TODO
	return 0
}

// MaybeGrow adds d to *MutBlob, else returns input unchanged.
func MaybeGrow(s Sizer, d int) Sizer { // TODO
	return s
}
