package topic08_packages_modules

import "testing"

func TestEX09_ValidateImportASCII(t *testing.T) {
	cases := []struct {
		in string
		ok bool
	}{
		{"example.com/mod/pkg", true},
		{"a/b/c", true},
		{"x", true},
		{"/leading", false},
		{"trailing/", false},
		{"double//slash", false},
		{"has space", false},
		{"dot/./pkg", false},
		{"dotdot/../pkg", false},
		{"..", false},
		{".", false},
		{"example.com/ユニコード", false},
		{"Upper/Case", true},
	}
	for i, tc := range cases {
		t.Run("EX09_ImportASCII_"+string(rune('A'+i)), func(t *testing.T) {
			err := ValidateImportASCII(tc.in)
			if tc.ok && err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error for %q", tc.in)
			}
		})
	}
}
