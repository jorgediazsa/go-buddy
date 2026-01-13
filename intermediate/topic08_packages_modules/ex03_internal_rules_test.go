package topic08_packages_modules

import "testing"

func TestEX03_AllowedImport_InternalRules(t *testing.T) {
	cases := []struct {
		from, to string
		want     bool
	}{
		{"example.com/a/b", "example.com/a/b/internal/x", true},
		{"example.com/a/b/c", "example.com/a/b/internal/x", true},
		{"example.com/a/c", "example.com/a/b/internal/x", false},
		{"example.com/a", "example.com/a/b/internal/x", false},
		{"example.com/a/b", "example.com/a/b/c", true},
		{"example.com/a/b", "example.com/a/b/internal", true},
		{"example.com/a/b", "example.com/a/b/internal/x/y", true},
		{"example.com/a/bb", "example.com/a/b/internal/x", false},
		{"example.com/a/b", "example.com/a/b/internalx/z", true},
		{"example.com/x/y", "example.com/x/y/internal/a/internal/b", true},
		{"example.com/x/z", "example.com/x/y/internal/a/internal/b", false},
		{"example.com/x/y/a", "example.com/x/y/internal", true},
		{"example.com/x", "example.com/x/internal/y", true},
		{"example.com/z", "example.com/x/internal/y", false},
	}
	for i, tc := range cases {
		t.Run("EX03_InternalRules_"+string(rune('A'+i)), func(t *testing.T) {
			if got := AllowedImport(tc.from, tc.to); got != tc.want {
				t.Fatalf("from %q to %q => got %v want %v", tc.from, tc.to, got, tc.want)
			}
		})
	}
}
