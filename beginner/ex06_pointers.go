package beginner

/*
Exercise EX06 — Pointers, mutation, and defensive arithmetic

Why this matters
- Many production bugs stem from misunderstanding when mutations affect callers. Pointers make intent explicit.
- Arithmetic on money-like integers must guard against overflow/underflow.

Requirements
- Define Account struct { Owner string; Balance int64 }.
- Deposit(a *Account, amount int64) error: amount>0; check overflow; mutate Balance.
- Withdraw(a *Account, amount int64) error: amount>0; sufficient funds; mutate Balance.
- Transfer(src, dst *Account, amount int64) error: perform withdraw then deposit atomically (no partial change on error).

Constraints and pitfalls
- Never panic; validate inputs (nil accounts, non-positive amounts).
- Avoid using big.Int; stick to int64 and explicit checks.
- Transfer must leave balances unchanged on failure.

Tricky edge cases
- Amount equals 0 or negative.
- Withdraw exact balance to 0.
- Overflow when adding to a large balance.
*/

import "errors"

type Account struct {
	Owner   string
	Balance int64
}

func Deposit(a *Account, amount int64) error { // TODO: implement
	if a == nil {
		return errors.New("nil account")
	}
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	// TODO: overflow check and mutation
	return errors.New("not implemented")
}

func Withdraw(a *Account, amount int64) error { // TODO: implement
	if a == nil {
		return errors.New("nil account")
	}
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	// TODO: sufficient balance and mutation
	return errors.New("not implemented")
}

func Transfer(src, dst *Account, amount int64) error { // TODO: implement
	if src == nil || dst == nil {
		return errors.New("nil account")
	}
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	// TODO: ensure atomicity on failure
	return errors.New("not implemented")
}
