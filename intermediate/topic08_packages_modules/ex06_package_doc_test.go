package topic08_packages_modules

import "testing"

func TestEX06_ValidatePackageDocHeader(t *testing.T) {
	cases := []struct {
		name, doc string
		ok        bool
	}{
		{"foo", "Package foo provides utilities for X.", true},
		{"bar", "  Package bar has detailed docs spanning multiple lines.\nMore...", true},
		{"zap", "package zap lowercase start.", false},
		{"zip", "Package zip missing period", false},
		{"zoom", "Package zoom short.", false}, // too short (<20)
		{"alpha", " Package beta leading space wrong name.", false},
		{"alpha", "Package alpha correct first line but no period on first line\n.", true},
		{"m", "Package m exactly twenty chars.", true},
		{"m", "Package m nineteen chars", false},
		{"n", "\n\nPackage n valid header on third line.\ntrailing", true},
		{"p", "", false},
		{"", "Package q has name but empty param.", false},
	}
	for i, tc := range cases {
		t.Run("EX06_PkgDoc_"+string(rune('A'+i)), func(t *testing.T) {
			err := ValidatePackageDocHeader(tc.name, tc.doc)
			if tc.ok && err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error")
			}
		})
	}
}
