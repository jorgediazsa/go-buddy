package intermediate

/*
Exercise EX01 — Methods: value vs pointer receivers

Why this matters
- Choosing pointer vs value receivers impacts mutability, method sets, and interface satisfaction.

Requirements
- Implement a Counter type with:
  - func (c Counter) Value() int
  - func (c *Counter) Inc()
  - func (c Counter) IncValue() Counter // returns an incremented copy; original unchanged
  - func (c Counter) Clone() Counter
- Implement helpers:
  - MaybeReset(c *Counter, threshold int) bool // if c.Value() >= threshold and threshold>=0, set to 0 and return true
  - SumValues(cs []Counter) int // sum of values; do not mutate inputs

Constraints and pitfalls
- Inc must be safe to call on a nil *Counter (no panic; treat as no-op).
- MaybeReset must be safe on nil *Counter; return false and do nothing.
- Keep API small and explicit.

Tricky edge cases
- threshold == 0 should reset when value >= 0 (i.e., any non-negative value), but negative thresholds never reset.
*/

// Counter holds an integer count.
type Counter struct{ n int }

// Value returns the current value.
func (c Counter) Value() int { // TODO: implement
	return 0
}

// Inc increments the counter in-place. Safe on nil receiver (no-op).
func (c *Counter) Inc() { // TODO: implement
}

// IncValue returns a copy incremented by 1; original is unchanged.
func (c Counter) IncValue() Counter { // TODO: implement
	return Counter{}
}

// Clone returns a shallow copy of the counter value.
func (c Counter) Clone() Counter { // TODO: implement
	return Counter{}
}

// MaybeReset sets the counter to 0 if value >= threshold and threshold>=0.
func MaybeReset(c *Counter, threshold int) bool { // TODO: implement
	return false
}

// SumValues returns the sum of values across counters.
func SumValues(cs []Counter) int { // TODO: implement
	return 0
}
