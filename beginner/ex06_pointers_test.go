package beginner

import "testing"

func TestEX06_Deposit(t *testing.T) {
	const maxInt64 = int64(^uint64(0) >> 1)

	t.Run("nil_account", func(t *testing.T) {
		if err := Deposit(nil, 1); err == nil {
			t.Fatalf("expected error for nil account")
		}
	})

	cases := []struct {
		name   string
		start  int64
		amount int64
		ok     bool
		want   int64
	}{
		{"basic", 0, 1, true, 1},
		{"add_positive", 100, 50, true, 150},
		{"reject_zero_amount", 100, 0, false, 100},
		{"reject_negative_amount", 100, -1, false, 100},

		{"max_exact", 0, maxInt64, true, maxInt64},
		{"overflow_by_1", maxInt64, 1, false, maxInt64},
		{"overflow_large_amount", 1, maxInt64, false, 1},
		{"edge_no_overflow", maxInt64 - 1, 1, true, maxInt64},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			a := &Account{Owner: "A", Balance: tc.start}
			before := a.Balance

			err := Deposit(a, tc.amount)

			if tc.ok {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if a.Balance != tc.want {
					t.Fatalf("balance %d want %d", a.Balance, tc.want)
				}
			} else {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				// must not change on failure
				if a.Balance != before {
					t.Fatalf("balance changed on error: got %d want %d", a.Balance, before)
				}
			}
		})
	}
}

func TestEX06_Withdraw(t *testing.T) {
	t.Run("nil_account", func(t *testing.T) {
		if err := Withdraw(nil, 1); err == nil {
			t.Fatalf("expected error for nil account")
		}
	})

	cases := []struct {
		name   string
		start  int64
		amount int64
		ok     bool
		want   int64
	}{
		{"basic", 100, 40, true, 60},
		{"withdraw_all_to_zero", 100, 100, true, 0},
		{"insufficient_by_1", 100, 101, false, 100},
		{"insufficient_from_zero", 0, 1, false, 0},
		{"reject_zero_amount", 50, 0, false, 50},
		{"reject_negative_amount", 50, -1, false, 50},
		{"negative_balance_still_insufficient", -10, 1, false, -10},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			a := &Account{Owner: "A", Balance: tc.start}
			before := a.Balance

			err := Withdraw(a, tc.amount)

			if tc.ok {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if a.Balance != tc.want {
					t.Fatalf("balance %d want %d", a.Balance, tc.want)
				}
			} else {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if a.Balance != before {
					t.Fatalf("balance changed on error: got %d want %d", a.Balance, before)
				}
			}
		})
	}
}

func TestEX06_Transfer(t *testing.T) {
	const maxInt64 = int64(^uint64(0) >> 1)

	t.Run("nil_accounts", func(t *testing.T) {
		a := &Account{Owner: "A", Balance: 10}
		if err := Transfer(nil, a, 1); err == nil {
			t.Fatalf("expected error for nil src")
		}
		if err := Transfer(a, nil, 1); err == nil {
			t.Fatalf("expected error for nil dst")
		}
	})

	cases := []struct {
		name   string
		s, d   int64
		amount int64
		ok     bool
		ws, wd int64
	}{
		{"basic", 100, 0, 50, true, 50, 50},
		{"transfer_all", 100, 0, 100, true, 0, 100},

		{"insufficient_funds", 10, 0, 11, false, 10, 0},
		{"reject_zero_amount", 10, 0, 0, false, 10, 0},
		{"reject_negative_amount", 10, 0, -1, false, 10, 0},

		// critical: withdraw would succeed, but deposit would overflow -> NO partial changes
		{"atomic_on_deposit_overflow", 200, maxInt64 - 100, 150, false, 200, maxInt64 - 100},

		// another overflow edge: dst at max, any positive transfer must fail, balances unchanged
		{"dst_at_max_overflow", 1, maxInt64, 1, false, 1, maxInt64},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			src := &Account{Owner: "S", Balance: tc.s}
			dst := &Account{Owner: "D", Balance: tc.d}

			srcBefore := src.Balance
			dstBefore := dst.Balance

			err := Transfer(src, dst, tc.amount)

			if tc.ok {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if src.Balance != tc.ws || dst.Balance != tc.wd {
					t.Fatalf("got src=%d dst=%d want src=%d dst=%d", src.Balance, dst.Balance, tc.ws, tc.wd)
				}
			} else {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				// atomicity: balances unchanged on failure
				if src.Balance != srcBefore || dst.Balance != dstBefore {
					t.Fatalf("balances changed on failure: src %d->%d dst %d->%d",
						srcBefore, src.Balance, dstBefore, dst.Balance)
				}
			}
		})
	}
}
