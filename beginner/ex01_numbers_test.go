package beginner

import (
	"fmt"
	"math"
	"testing"
)

func TestEX01_ParseUint8Strict(t *testing.T) {
	cases := []struct {
		in   string
		ok   bool
		want uint8
	}{
		{"0", true, 0},
		{"255", true, 255},
		{"001", true, 1},
		{"000", true, 0},

		{"256", false, 0},
		{"000256", false, 0}, // overflow even with leading zeros

		{"-1", false, 0},
		{"+1", false, 0},   // strict: no signs
		{" 1", false, 0},   // strict: no spaces
		{"1 ", false, 0},   // strict: no spaces
		{"12a", false, 0},  // non-digit
		{"1_0", false, 0},  // underscore not allowed
		{"01\t", false, 0}, // hidden whitespace
		{"01\n", false, 0}, // hidden whitespace
		{"", false, 0},     // empty
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("in=%q", tc.in), func(t *testing.T) {
			t.Parallel()

			got, err := ParseUint8Strict(tc.in)
			if tc.ok && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error, got nil for %q", tc.in)
			}
			if tc.ok && got != tc.want {
				t.Fatalf("got %d want %d", got, tc.want)
			}
		})
	}
}

func TestEX01_AddInt32(t *testing.T) {
	M := int32(math.MaxInt32)
	m := int32(math.MinInt32)

	cases := []struct {
		name string
		a, b int32
		ok   bool
		want int32
	}{
		{"zero", 0, 0, true, 0},
		{"small_pos", 1, 2, true, 3},
		{"cancel", -1, 1, true, 0},
		{"cancel_big", 123456, -123456, true, 0},
		{"max_plus_zero", M, 0, true, M},
		{"min_plus_zero", m, 0, true, m},

		{"overflow_max_plus_1", M, 1, false, 0},
		{"underflow_min_minus_1", m, -1, false, 0},

		{"near_max_minus_1", M, -1, true, M - 1},
		{"near_min_plus_1", m, 1, true, m + 1},

		{"overflow_max_plus_max", M, M, false, 0},
		{"underflow_min_plus_min", m, m, false, 0},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := AddInt32(tc.a, tc.b)
			if tc.ok && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error, got nil for %+v", tc)
			}
			if tc.ok && got != tc.want {
				t.Fatalf("got %d want %d", got, tc.want)
			}
		})
	}
}

func TestEX01_PercentOf(t *testing.T) {
	cases := []struct {
		name        string
		part, total int64
		ok          bool
		want        int64
	}{
		{"zero", 0, 100, true, 0},
		{"one_percent", 1, 100, true, 1},
		{"thirty_three", 33, 100, true, 33},
		{"quarter", 50, 200, true, 25},
		{"all", 100, 100, true, 100},
		{"floor_3_7", 3, 7, true, 42}, // floor(42.857...)
		{"floor_1_3", 1, 3, true, 33},

		{"invalid_total_zero", 10, 0, false, 0},
		{"invalid_part_negative", -1, 10, false, 0},
		{"invalid_part_gt_total", 11, 10, false, 0},

		// Forces overflow guard for (part*100)
		{"overflow_part_times_100", math.MaxInt64/100 + 1, math.MaxInt64, false, 0},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := PercentOf(tc.part, tc.total)
			if tc.ok && err != nil {
				t.Fatalf("unexpected error: %v (got=%d)", err, got)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error, got nil (got=%d)", got)
			}
			if tc.ok && got != tc.want {
				t.Fatalf("got %d want %d", got, tc.want)
			}
		})
	}
}
