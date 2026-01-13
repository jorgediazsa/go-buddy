package topic08_packages_modules

import "testing"

func TestEX01_SanitizePackageName(t *testing.T) {
	cases := []struct {
		in   string
		ok   bool
		want string
	}{
		{"foo", true, "foo"},
		{" Foo ", true, "foo"},
		{"FOO_BAR", true, "foo_bar"},
		{"foo123", true, "foo123"},
		{"_hidden", false, ""},
		{"9start", false, ""},
		{"", false, ""},
		{"has-dash", false, ""},
		{"white space", false, ""},
		{"MiXeD_Case_123", true, "mixed_case_123"},
		{"a", true, "a"},
		{string(make([]byte, 64)), false, ""}, // length > 63
	}
	for i, tc := range cases {
		t.Run("EX01_Sanitize_"+string(rune('A'+i)), func(t *testing.T) {
			got, err := SanitizePackageName(tc.in)
			if tc.ok && err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error")
			}
			if tc.ok && got != tc.want {
				t.Fatalf("got %q want %q", got, tc.want)
			}
		})
	}
}

func TestEX01_IsExportedName(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"Foo", true},
		{"URL", true},
		{"foo", false},
		{"_X", false},
		{"", false},
		{"Äbc", false}, // ASCII check by simple rune upper test may fail; we expect false due to non-ASCII handling here
		{"X1", true},
		{"x1", false},
		{"JSON", true},
		{"λ", false},
		{"A_B", true},
	}
	for i, tc := range cases {
		t.Run("EX01_IsExported_"+string(rune('A'+i)), func(t *testing.T) {
			if got := IsExportedName(tc.in); got != tc.want {
				t.Fatalf("got %v want %v for %q", got, tc.want, tc.in)
			}
		})
	}
}
