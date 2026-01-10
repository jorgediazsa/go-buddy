package intermediate

/*
Exercise EX09 — Embedding chain and overrides: Clock family

Why this matters
- Multi-level embedding demonstrates method promotion and how overrides at each level change behavior.

Requirements
- BaseClock holds a Stamp string.
  - func (b BaseClock) NowString() string           // returns b.Stamp
  - func (b *BaseClock) SetStamp(s string)          // pointer receiver
- ServiceClock embeds BaseClock and adds Prefix string.
  - func (s ServiceClock) NowString() string        // returns Prefix + b.NowString()
- AuditClock embeds ServiceClock and adds Suffix string.
  - func (a AuditClock) NowString() string          // returns s.NowString()+Suffix
- Helpers:
  - PromoteNow(b BaseClock) string
  - TouchStamp(s *ServiceClock, stamp string)
  - Identify(a *AuditClock) string // use most specific NowString

Constraints and pitfalls
- Ensure SetStamp is a pointer receiver to mutate BaseClock inside embedded structs.
- Zero-values must not panic and should produce deterministic strings.

Tricky edge cases
- Empty prefix/suffix; touching stamp via embedded method promotion.
*/

type BaseClock struct{ Stamp string }

func (b BaseClock) NowString() string { // TODO
	return ""
}

func (b *BaseClock) SetStamp(s string) { // TODO
}

type ServiceClock struct {
	BaseClock
	Prefix string
}

func (s ServiceClock) NowString() string { // TODO
	return ""
}

type AuditClock struct {
	ServiceClock
	Suffix string
}

func (a AuditClock) NowString() string { // TODO
	return ""
}

func PromoteNow(b BaseClock) string { // TODO
	return ""
}

func TouchStamp(s *ServiceClock, stamp string) { // TODO
}

func Identify(a *AuditClock) string { // TODO
	if a == nil {
		return ""
	}
	return ""
}
