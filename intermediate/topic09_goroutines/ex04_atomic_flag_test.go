package topic09_goroutines

import "testing"

func TestEX09_Ex04_AtomicFlagBasic(t *testing.T) {
	var f AtomicFlag
	if f.Locked() {
		t.Fatalf("new flag should be unlocked")
	}
	if !f.TryLock() {
		t.Fatalf("first TryLock should succeed")
	}
	if !f.Locked() {
		t.Fatalf("flag should be locked")
	}
	if f.TryLock() {
		t.Fatalf("second TryLock should fail while held")
	}
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("unexpected panic on first unlock: %v", r)
		}
	}()
	f.Unlock()
	if f.Locked() {
		t.Fatalf("flag should be unlocked after Unlock")
	}
}

func TestEX09_Ex04_AtomicFlagDoubleUnlockPanics(t *testing.T) {
	var f AtomicFlag
	// Unlock without lock should panic
	didPanic := false
	func() {
		defer func() {
			if recover() != nil {
				didPanic = true
			}
		}()
		f.Unlock()
	}()
	if !didPanic {
		t.Fatalf("expected panic on Unlock without lock")
	}

	// Lock, unlock twice
	if !f.TryLock() {
		t.Fatalf("TryLock failed")
	}
	didPanic = false
	func() {
		defer func() {
			if recover() != nil {
				didPanic = true
			}
		}()
		f.Unlock()
		f.Unlock()
	}()
	if !didPanic {
		t.Fatalf("expected panic on double Unlock")
	}
}
