package beginner

import "testing"

func TestEX02_NormalizeEmail(t *testing.T) {
	cases := []struct {
		in   string
		ok   bool
		want string
	}{
		{"Alice@Example.COM", true, "Alice@example.com"},
		{" alice@Example.com ", true, "alice@example.com"},
		{"a+b@sub.Example.com", true, "a+b@sub.example.com"},
		{"@example.com", false, ""},
		{"alice@", false, ""},
		{"aliceexample.com", false, ""},
		{"al ice@example.com", false, ""},
		{"alice@@example.com", false, ""},
		{"", false, ""},
		{"alice@例え.テスト", true, "alice@例え.テスト"},
	}
	for _, tc := range cases {
		t.Run("EX02_NormalizeEmail_"+safeName(tc.in), func(t *testing.T) {
			got, err := NormalizeEmail(tc.in)
			if tc.ok && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error, got nil for %q", tc.in)
			}
			if tc.ok && got != tc.want {
				t.Fatalf("got %q want %q", got, tc.want)
			}
		})
	}
}

func TestEX02_CountRunesCategories(t *testing.T) {
	cases := []struct {
		in                      string
		letters, digits, spaces int
	}{
		{"", 0, 0, 0},
		{"abc", 3, 0, 0},
		{"a1b2c3", 3, 3, 0},
		{"a b c", 3, 0, 2},
		{"你好 123", 2, 3, 1},
		{"\t\n ", 0, 0, 3},
		{"ΑΒΓ123", 3, 3, 0},
		{"👩‍💻", 0, 0, 0}, // zero letters/digits/spaces per simple categories
		{"é", 1, 0, 0},
	}
	for _, tc := range cases {
		t.Run("EX02_CountRunesCategories_"+safeName(tc.in), func(t *testing.T) {
			l, d, s := CountRunesCategories(tc.in)
			if l != tc.letters || d != tc.digits || s != tc.spaces {
				t.Fatalf("got (%d,%d,%d) want (%d,%d,%d)", l, d, s, tc.letters, tc.digits, tc.spaces)
			}
		})
	}
}

func TestEX02_SubstringRunes(t *testing.T) {
	s := "héllo世界"
	cases := []struct {
		start, length int
		ok            bool
		want          string
	}{
		{0, 2, true, "hé"},
		{1, 2, true, "él"},
		{2, 3, true, "llo"},
		{5, 2, true, "世界"},
		{7, 0, true, ""},
		{-1, 1, false, ""},
		{0, -1, false, ""},
		{8, 1, false, ""},
		{6, 3, false, ""},
	}
	for i, tc := range cases {
		t.Run("EX02_SubstringRunes_"+string(rune('A'+i)), func(t *testing.T) {
			got, err := SubstringRunes(s, tc.start, tc.length)
			if tc.ok && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if tc.ok && got != tc.want {
				t.Fatalf("got %q want %q", got, tc.want)
			}
		})
	}
}

func safeName(s string) string {
	if s == "" {
		return "empty"
	}
	out := make([]rune, 0, len(s))
	for _, r := range s {
		if r == '@' || r == ' ' || r == '\\' || r == '/' || r == '\t' || r == '\n' {
			out = append(out, '_')
		} else if r < 128 {
			out = append(out, r)
		} else {
			out = append(out, 'u')
		}
		if len(out) >= 12 {
			break
		}
	}
	return string(out)
}
