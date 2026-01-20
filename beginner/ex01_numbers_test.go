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
		{"-1", false, 0},
		{"+1", false, 0},
		{" 1", false, 0},
		{"1 ", false, 0},
		{"12a", false, 0},
		{"", false, 0},
	}
	for _, tc := range cases {
		t.Run("EX01_ParseUint8Strict_"+tc.in, func(t *testing.T) {
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
		a, b int32
		ok   bool
		want int32
	}{
		{0, 0, true, 0},
		{1, 2, true, 3},
		{-1, 1, true, 0},
		{123456, -123456, true, 0},
		{M, 0, true, M},
		{m, 0, true, m},
		{M, 1, false, 0},
		{m, -1, false, 0},
		{M, -1, true, M - 1},
		{m, 1, true, m + 1},
	}
	for i, tc := range cases {
		t.Run("EX01_AddInt32_"+string(rune('A'+i)), func(t *testing.T) {
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
		part, total int64
		ok          bool
		want        int64
	}{
		{0, 100, true, 0},
		{1, 100, true, 1},
		{33, 100, true, 33},
		{50, 200, true, 25},
		{100, 100, true, 100},
		{3, 7, true, 42}, // floor(42.857)
		{1, 3, true, 33},
		{10, 0, false, 0},
		{-1, 10, false, 0},
		{11, 10, false, 0},
	}
	for _, tc := range cases {
		name := "EX01_PercentOf_" + fmtInt64(tc.part) + "_" + fmtInt64(tc.total)
		t.Run(name, func(t *testing.T) {
			got, err := PercentOf(tc.part, tc.total)
			fmt.Println("part: ", tc.part, "total: ", tc.total)
			if tc.ok && err != nil {
				t.Fatalf("unexpected error: %v, got: %v", err, got)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if tc.ok && got != tc.want {
				t.Fatalf("got %d want %d", got, tc.want)
			}
		})
	}
}

func fmtInt64(v int64) string {
	if v < 0 {
		return "m"
	}
	return string(rune('a' + (v % 26)))
}
