package intermediate

/*
Exercise EX02 — Embedding, method promotion, and shadowing

Why this matters
- Embedding composes behavior and promotes methods, but overrides (shadowing) can change semantics.

Requirements
- Define type Base with methods:
  - ID() string                     // value receiver, returns base id
  - SetID(id string)                // pointer receiver, sets base id
  - Describe() string               // value receiver, returns "Base:<id>"
- Define type User embedding Base with fields Name string, and method:
  - Describe() string               // overrides to include Name (e.g., "User:<Name>:<Base.ID>")
- Define type Manager embedding User and adding field Level int and method:
  - Describe() string               // overrides to include Level
- Helper functions:
  - PromoteDescribe(b Base) string  // returns b.Describe()
  - TouchID(u *User, id string)     // sets base id via promoted method and returns nothing
  - WhoAmI(m *Manager) string       // returns Describe() using the most specific override

Constraints and pitfalls
- Ensure method promotion works for embedded fields.
- Value vs pointer receiver methods: SetID must be on pointer receiver to mutate.
- Zero-values should be handled gracefully (empty ids, names).

Tricky edge cases
- Calling Describe on zero-value User/Manager should not panic and produce deterministic strings.
- Setting ID via TouchID should update the embedded Base inside User.
*/

// Base is a foundational type holding an id.
type Base struct{ id string }

// ID returns the id.
func (b Base) ID() string { // TODO
	return ""
}

// SetID sets the id (pointer receiver).
func (b *Base) SetID(id string) { // TODO
}

// Describe returns a string description for Base.
func (b Base) Describe() string { // TODO
	return ""
}

// User embeds Base and adds a Name.
type User struct {
	Base
	Name string
}

// Describe overrides Base.Describe.
func (u User) Describe() string { // TODO
	return ""
}

// Manager embeds User and adds Level.
type Manager struct {
	User
	Level int
}

// Describe overrides User.Describe.
func (m Manager) Describe() string { // TODO
	return ""
}

// PromoteDescribe returns Base.Describe(b).
func PromoteDescribe(b Base) string { // TODO
	return ""
}

// TouchID sets the id on the embedded Base via promoted methods.
func TouchID(u *User, id string) { // TODO
}

// WhoAmI returns the most specific Describe.
func WhoAmI(m *Manager) string { // TODO
	if m == nil {
		return ""
	}
	return ""
}
