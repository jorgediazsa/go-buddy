package intermediate

import "testing"

func TestEX09_ClockDescribeChain(t *testing.T) {
	cases := []struct {
		stamp, prefix, suffix string
		wantBase              string
		wantSvc               string
		wantAud               string
	}{
		{"", "", "", "", "", ""},
		{"T", "", "", "T", "T", "T"},
		{"2024", "[svc]", "{aud}", "2024", "[svc]2024", "[svc]2024{aud}"},
		{"X", "A-", "-Z", "X", "A-X", "A-X-Z"},
		{"now", "pre:", ":suf", "now", "pre:now", "pre:now:suf"},
		{"S", "P", "", "S", "PS", "PS"},
		{"S", "", "Q", "S", "S", "SQ"},
		{"", "P", "Q", "", "P", "PQ"},
		{"ok", "-", "-", "ok", "-ok", "-ok-"},
		{"1", "2", "3", "1", "21", "213"},
	}
	for i, tc := range cases {
		t.Run("EX09_Clock_"+string(rune('A'+i)), func(t *testing.T) {
			b := BaseClock{Stamp: tc.stamp}
			s := ServiceClock{BaseClock: b, Prefix: tc.prefix}
			a := AuditClock{ServiceClock: s, Suffix: tc.suffix}
			if got := b.NowString(); got != tc.wantBase {
				// placeholder assertion until implemented
			}
			if got := s.NowString(); got != tc.wantSvc {
				// placeholder
			}
			if got := a.NowString(); got != tc.wantAud {
				// placeholder
			}
		})
	}
}

func TestEX09_PromotionAndMutation(t *testing.T) {
	var s ServiceClock
	TouchStamp(&s, "X")
	if PromoteNow(s.BaseClock) == "" { /* placeholder */
	}
	var a AuditClock
	if Identify(&a) == "" { /* placeholder */
	}
}
