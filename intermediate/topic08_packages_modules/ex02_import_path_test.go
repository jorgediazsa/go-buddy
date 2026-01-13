package topic08_packages_modules

import "testing"

func TestEX02_BuildImportPath(t *testing.T) {
	cases := []struct {
		mod, pkg string
		ok       bool
		want     string
	}{
		{"example.com/mod", "pkg", true, "example.com/mod/pkg"},
		{"example.com/mod/", "/pkg", true, "example.com/mod/pkg"},
		{"example.com//mod", "//pkg//sub", true, "example.com/mod/pkg/sub"},
		{"example.com/mod", "./pkg", false, ""},
		{"example.com/mod", "../pkg", false, ""},
		{"", "pkg", false, ""},
		{"mod", "", false, ""},
		{" spaced ", "pkg", false, ""},
		{"mod", "space d", false, ""},
		{"example.com/m", "sub/dir", true, "example.com/m/sub/dir"},
		{"EXAMPLE.COM/MOD", "PKG", true, "EXAMPLE.COM/MOD/PKG"},
	}
	for i, tc := range cases {
		t.Run("EX02_Build_"+string(rune('A'+i)), func(t *testing.T) {
			got, err := BuildImportPath(tc.mod, tc.pkg)
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

func TestEX02_IsInternalPath(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"example.com/mod/internal/x", true},
		{"internal", true},
		{"a/b/c", false},
		{"a/internalb/c", false},
		{"a/internal/c", true},
		{"/internal/", true},
		{"////internal////x", true},
		{"", false},
		{"internal_/x", false},
		{"a/b/Internal/c", false},
		{"internalx", false},
	}
	for i, tc := range cases {
		t.Run("EX02_Internal_"+string(rune('A'+i)), func(t *testing.T) {
			if got := IsInternalPath(tc.in); got != tc.want {
				t.Fatalf("got %v want %v for %q", got, tc.want, tc.in)
			}
		})
	}
}
