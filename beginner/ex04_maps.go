package beginner

/*
Exercise EX04 — Maps: safe operations and ordering

Why this matters
- Maps are reference types; careless mutation leaks across call sites. Deterministic ordering requires explicit sorting.

Requirements
- MergeSafe: return a new map containing keys from a and b; b overrides conflicts. Do not mutate inputs.
- SortedKeysByValue: return keys sorted by descending value, then ascending key for ties.
- GetInt: read an int from a map[string]string with strict base-10 parsing and optional defaulting via ok flag.

Constraints and pitfalls
- Return non-nil empty maps when appropriate.
- No panics on nil inputs.
- Avoid extra allocations beyond what’s necessary.

Tricky edge cases
- Nil vs empty maps.
- Conflicting keys; stable tie-breaking when values match.
- Non-numeric strings for GetInt.
*/

import (
	"errors"
)

// MergeSafe merges two maps into a new map (b overrides).
func MergeSafe(a, b map[string]int) map[string]int { // TODO: implement
	out := make(map[string]int)
	// TODO: copy a, then overlay b
	return out
}

// SortedKeysByValue returns keys sorted by desc value, then asc key.
func SortedKeysByValue(m map[string]int) []string { // TODO: implement
	// TODO: collect and sort
	return nil
}

// GetInt retrieves and parses key from m strictly as base-10 int.
func GetInt(m map[string]string, key string) (int, bool, error) { // TODO: implement
	if m == nil {
		return 0, false, errors.New("nil map")
	}
	return 0, false, errors.New("not implemented")
}
