package beginner

import "testing"

func TestEX06_Deposit(t *testing.T) {
	cases := []struct {
		start  int64
		amount int64
		ok     bool
		want   int64
	}{
		{0, 1, true, 1},
		{100, 50, true, 150},
		{100, 0, false, 100},
		{100, -1, false, 100},
		{0, 9_223_372_036_854_775_807, true, 9_223_372_036_854_775_807}, // MaxInt64
		{1, 9_223_372_036_854_775_807, false, 1},                        // overflow
	}
	for i, tc := range cases {
		t.Run("EX06_Deposit_"+string(rune('A'+i)), func(t *testing.T) {
			a := &Account{Owner: "A", Balance: tc.start}
			err := Deposit(a, tc.amount)
			if tc.ok && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if a.Balance != tc.want {
				t.Fatalf("balance %d want %d", a.Balance, tc.want)
			}
		})
	}
}

func TestEX06_Withdraw(t *testing.T) {
	cases := []struct {
		start  int64
		amount int64
		ok     bool
		want   int64
	}{
		{100, 40, true, 60},
		{100, 100, true, 0},
		{100, 101, false, 100},
		{0, 1, false, 0},
		{50, 0, false, 50},
		{50, -1, false, 50},
	}
	for i, tc := range cases {
		t.Run("EX06_Withdraw_"+string(rune('A'+i)), func(t *testing.T) {
			a := &Account{Owner: "A", Balance: tc.start}
			err := Withdraw(a, tc.amount)
			if tc.ok && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error")
			}
			if a.Balance != tc.want {
				t.Fatalf("balance %d want %d", a.Balance, tc.want)
			}
		})
	}
}

func TestEX06_Transfer(t *testing.T) {
	cases := []struct {
		s, d   int64
		amount int64
		ok     bool
		ws, wd int64
	}{
		{100, 0, 50, true, 50, 50},
		{100, 0, 100, true, 0, 100},
		{10, 0, 11, false, 10, 0},
		{0, 9_223_372_036_854_775_700, 200, false, 0, 9_223_372_036_854_775_700}, // dst overflow if allowed; must remain unchanged
		{0, 0, 0, false, 0, 0},
		{0, 0, -1, false, 0, 0},
	}
	for i, tc := range cases {
		t.Run("EX06_Transfer_"+string(rune('A'+i)), func(t *testing.T) {
			src := &Account{Owner: "S", Balance: tc.s}
			dst := &Account{Owner: "D", Balance: tc.d}
			err := Transfer(src, dst, tc.amount)
			if tc.ok && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error")
			}
			if src.Balance != tc.ws || dst.Balance != tc.wd {
				t.Fatalf("got src=%d dst=%d want src=%d dst=%d", src.Balance, dst.Balance, tc.ws, tc.wd)
			}
		})
	}
}
