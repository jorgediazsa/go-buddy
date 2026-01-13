package topic08_packages_modules

import "testing"

func TestEX10_ValidateGoFilename(t *testing.T) {
	cases := []struct {
		in string
		ok bool
	}{
		{"main.go", true},
		{"util_test.go", true},
		{"foo_bar.go", true},
		{"FooBar.go", true},
		{"a..go", true},
		{".hidden.go", false},
		{" has space.go", false},
		{"noext", false},
		{"endswithdot.", false},
		{"_start.go", false},
		{"-bad.go", false},
		{"nonascii_µ.go", false},
		{"unix_amd64.go", true},
		{"UNICODE_测试.go", false},
		{"good.name.go", true},
	}
	for i, tc := range cases {
		t.Run("EX10_GoFilename_"+string(rune('A'+i)), func(t *testing.T) {
			err := ValidateGoFilename(tc.in)
			if tc.ok && err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error for %q", tc.in)
			}
		})
	}
}
