package beginner

import "testing"

func TestEX10_NormalizePathSimple(t *testing.T) {
	cases := []struct {
		in   string
		ok   bool
		want string
	}{
		{"/a/b/./c/../d", true, "/a/b/d"},
		{"/", true, "/"},
		{"/..", false, ""},
		{"a/b/..", true, "a"},
		{"../a", true, "../a"},
		{"", false, ""},
		{"a//b", false, ""},
		{"a/b/../../c", true, "c"},
		{"a/b/../../..", true, ".."},
		{"/a//b/c", false, ""},
		{"/a/../../b", false, ""},
		{"a/b/", true, "a/b"},
	}
	for _, tc := range cases {
		t.Run("EX10_NormalizePathSimple_"+safeName(tc.in), func(t *testing.T) {
			got, err := NormalizePathSimple(tc.in)
			if tc.ok && err != nil {
				t.Fatalf("unexpected error: %v", err)
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

func TestEX10_JoinClean(t *testing.T) {
	cases := []struct {
		base, rel string
		ok        bool
		want      string
	}{
		{"a/b", "c/./d", true, "a/b/c/d"},
		{"a/b", "../c", true, "a/c"},
		{"", "a", false, ""},
		{"a/b", "/c", false, ""},
		{"/a", "../..", false, ""},
		{"/a", "b/./c/..", true, "/a/b"},
		{"/", "a", true, "/a"},
		{"/a", "b/c", true, "/a/b/c"},
	}
	for i, tc := range cases {
		t.Run("EX10_JoinClean_"+string(rune('A'+i)), func(t *testing.T) {
			got, err := JoinClean(tc.base, tc.rel)
			if tc.ok && err != nil {
				t.Fatalf("unexpected error: %v", err)
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

func TestEX10_SplitExt(t *testing.T) {
	cases := []struct {
		in        string
		base, ext string
	}{
		{"file.txt", "file", ".txt"},
		{"archive.tar.gz", "archive.tar", ".gz"},
		{".bashrc", ".bashrc", ""},
		{"noext", "noext", ""},
		{"a.b.c.", "a.b.c.", ""},
		{"a.b.c", "a.b", ".c"},
		{"..hidden.txt", "..hidden", ".txt"},
		{"", "", ""},
	}
	for i, tc := range cases {
		t.Run("EX10_SplitExt_"+string(rune('A'+i)), func(t *testing.T) {
			b, e := SplitExt(tc.in)
			if b != tc.base || e != tc.ext {
				t.Fatalf("got (%q,%q) want (%q,%q)", b, e, tc.base, tc.ext)
			}
		})
	}
}
