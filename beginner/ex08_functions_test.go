package beginner

import "testing"

func TestEX08_SplitHostPortStrict(t *testing.T) {
	cases := []struct {
		in   string
		ok   bool
		host string
		port int
	}{
		{"localhost:80", true, "localhost", 80},
		{"example.com:443", true, "example.com", 443},
		{"a:0", true, "a", 0},
		{"a:65535", true, "a", 65535},
		{":80", false, "", 0},
		{"host:", false, "", 0},
		{"host:65536", false, "", 0},
		{"host:-1", false, "", 0},
		{"host:+1", false, "", 0},
		{"host:12 3", false, "", 0},
		{"a:b", false, "", 0},
		{"a:b:c", false, "", 0},
		{" a:1", false, "", 0},
	}
	for _, tc := range cases {
		t.Run("EX08_SplitHostPortStrict_"+safeName(tc.in), func(t *testing.T) {
			h, p, err := SplitHostPortStrict(tc.in)
			if tc.ok && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error")
			}
			if tc.ok && (h != tc.host || p != tc.port) {
				t.Fatalf("got %q,%d want %q,%d", h, p, tc.host, tc.port)
			}
		})
	}
}

func TestEX08_RepeatJoin(t *testing.T) {
	cases := []struct {
		parts []string
		sep   string
		count int
		ok    bool
		want  string
	}{
		{nil, ",", 0, true, ""},
		{[]string{}, ",", 0, true, ""},
		{[]string{"a"}, "-", 1, true, "a"},
		{[]string{"a"}, "-", 2, true, "aa"},
		{[]string{"a", "b"}, ":", 2, true, "aa:bb"},
		{[]string{"", "b"}, "/", 2, true, ":/bb"},
		{[]string{"x", "y", "z"}, "", 3, true, "xxx yyy zzz"[:9-1]}, // placeholder; will be validated precisely
		{[]string{"a"}, ",", -1, false, ""},
	}
	// fix the placeholder expected precisely
	cases[6].want = "xxxyyyzzz"
	for i, tc := range cases {
		t.Run("EX08_RepeatJoin_"+string(rune('A'+i)), func(t *testing.T) {
			got, err := RepeatJoin(tc.parts, tc.sep, tc.count)
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

func TestEX08_MakeSuffixer(t *testing.T) {
	addTxt := MakeSuffixer(".txt")
	noOp := MakeSuffixer("")
	cases := []struct {
		f    func(string) string
		in   string
		want string
	}{
		{addTxt, "readme", "readme.txt"},
		{addTxt, "file.txt", "file.txt"},
		{addTxt, "a.b", "a.b.txt"},
		{noOp, "x", "x"},
		{MakeSuffixer("/"), "path/", "path/"},
		{MakeSuffixer("/"), "path", "path/"},
		{MakeSuffixer("xyz"), "", "xyz"},
		{MakeSuffixer("-end"), "mid-end", "mid-end"},
	}
	for i, tc := range cases {
		t.Run("EX08_MakeSuffixer_"+string(rune('A'+i)), func(t *testing.T) {
			got := tc.f(tc.in)
			if got != tc.want {
				t.Fatalf("got %q want %q", got, tc.want)
			}
		})
	}
}
