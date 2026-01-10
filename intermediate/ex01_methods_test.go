package intermediate

import "testing"

func TestEX01_CounterBasics(t *testing.T) {
	cases := []struct {
		start     int
		incTimes  int
		threshold int
		resetWant bool
		want      int
	}{
		{0, 0, 0, true, 0},
		{0, 1, 2, false, 1},
		{5, 3, 7, true, 0},
		{5, 1, 7, false, 6},
		{9, 0, 10, false, 9},
		{10, 0, 10, true, 0},
		{3, 4, -1, false, 7}, // negative threshold never resets
		{100, 100, 9999, false, 200},
		{1, 2, 2, true, 0},
		{1, 1, 3, false, 2},
	}
	for i, tc := range cases {
		t.Run("EX01_CounterBasics_"+string(rune('A'+i)), func(t *testing.T) {
			c := Counter{n: tc.start}
			for j := 0; j < tc.incTimes; j++ {
				c.Inc()
			}
			if got := c.Value(); got != tc.start+tc.incTimes { // expected before MaybeReset
				t.Fatalf("got value %d want %d before reset", got, tc.start+tc.incTimes)
			}
			reset := MaybeReset(&c, tc.threshold)
			if reset != tc.resetWant {
				t.Fatalf("reset=%v want %v", reset, tc.resetWant)
			}
			_ = c.Clone()
			if c.Value() != tc.want {
				t.Fatalf("got %d want %d after reset", c.Value(), tc.want)
			}
		})
	}
}

func TestEX01_IncOnNilReceiver(t *testing.T) {
	var c *Counter
	// should not panic
	c.Inc()
}

func TestEX01_IncValueDoesNotMutate(t *testing.T) {
	c := Counter{n: 5}
	copy := c.IncValue()
	if copy.Value() != 6 {
		t.Fatalf("copy.Value=%d want 6", copy.Value())
	}
	if c.Value() != 5 {
		t.Fatalf("c.Value mutated to %d", c.Value())
	}
}

func TestEX01_SumValues(t *testing.T) {
	cs := []Counter{{n: 1}, {n: 2}, {n: 3}, {n: 4}}
	got := SumValues(cs)
	if got != 10 {
		t.Fatalf("sum=%d want 10", got)
	}
}
