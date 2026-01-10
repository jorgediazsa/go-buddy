package beginner

/*
Exercise EX05 — Struct design, validation, and pointer updates

Why this matters
- Real systems pass structs around widely. Getting validation, zero-values, and pointer mutation right avoids entire classes of bugs.

Requirements
- Define a Person struct with fields: Name string, Email string, BirthYear int (Gregorian, four digits).
- ValidatePerson: ensure Name non-empty (trimmed), Email normalized via EX02.NormalizeEmail rules (domain lower), BirthYear within [1900..2100]. Return a sanitized copy.
- UpdateEmail: mutate the provided *Person email after validating with NormalizeEmail.
- AgeAtYear: compute age in a given year; errors if year < BirthYear or out of [1900..2100].

Constraints and pitfalls
- Do not panic; return errors.
- Do not silently coerce invalid years.
- Keep Person small and focused; no JSON tags or methods.

Tricky edge cases
- Names with leading/trailing spaces.
- Emails invalid per EX02 rules.
- Boundary years (1900, 2100) and future year equals BirthYear.
*/

import "errors"

type Person struct {
	Name      string
	Email     string
	BirthYear int
}

// ValidatePerson returns a sanitized copy if valid.
func ValidatePerson(p Person) (Person, error) { // TODO: implement
	return Person{}, errors.New("not implemented")
}

// UpdateEmail validates and updates the person's email in-place.
func UpdateEmail(p *Person, newEmail string) error { // TODO: implement
	if p == nil {
		return errors.New("nil person")
	}
	return errors.New("not implemented")
}

// AgeAtYear returns the age in targetYear.
func AgeAtYear(p Person, targetYear int) (int, error) { // TODO: implement
	return 0, errors.New("not implemented")
}
