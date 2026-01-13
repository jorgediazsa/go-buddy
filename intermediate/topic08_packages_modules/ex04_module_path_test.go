package topic08_packages_modules

import "testing"

func TestEX04_ValidateModulePath(t *testing.T) {
	cases := []struct {
		in string
		ok bool
	}{
		{"example.com/mod", true},
		{"example.com/mod/v2", true},
		{"example.com/mod/v10", true},
		{"example.com/mod/v1", true}, // not semantic major, but still a valid path
		{"example.com/mod/v0", true},
		{"example.com/mod/", false},
		{"/example.com/mod", false},
		{"example.com/ mod", false},
		{"", false},
		{"example.com/モジュール", false},
		{"EXAMPLE.COM/MOD/V2", true},
	}
	for i, tc := range cases {
		t.Run("EX04_Validate_"+string(rune('A'+i)), func(t *testing.T) {
			err := ValidateModulePath(tc.in)
			if tc.ok && err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error")
			}
		})
	}
}

func TestEX04_ExtractMajorVersion(t *testing.T) {
	cases := []struct {
		in     string
		wantN  int
		wantOK bool
	}{
		{"example.com/mod", 0, false},
		{"example.com/mod/v2", 2, true},
		{"example.com/mod/v10", 10, true},
		{"example.com/mod/v1", 0, false},
		{"example.com/mod/v0", 0, false},
		{"example.com/mod/v2/extra", 0, false},
		{"example.com/v3", 3, true},
		{"EXAMPLE.COM/MOD/V2", 0, false},
		{"example.com/mod/V2", 0, false},
		{"example.com/mod/v02", 2, true},
	}
	for i, tc := range cases {
		t.Run("EX04_Extract_"+string(rune('A'+i)), func(t *testing.T) {
			gotN, ok := ExtractMajorVersion(tc.in)
			if ok != tc.wantOK || gotN != tc.wantN {
				t.Fatalf("got (%d,%v) want (%d,%v)", gotN, ok, tc.wantN, tc.wantOK)
			}
		})
	}
}
