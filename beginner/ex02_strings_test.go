package beginner

import (
	"strings"
	"testing"
)

func TestEX02_NormalizeEmail(t *testing.T) {
	cases := []struct {
		name string
		in   string
		ok   bool
		want string
	}{
		// --- Happy paths ---
		{"basic_domain_lower", "Alice@Example.COM", true, "Alice@example.com"},
		{"trim_outer_spaces", " alice@Example.com ", true, "alice@example.com"},
		{"plus_in_local", "a+b@sub.Example.com", true, "a+b@sub.example.com"},
		{"unicode_domain_ok", "alice@例え.テスト", true, "alice@例え.テスト"},
		{"local_preserved_case", "ALICE@Example.com", true, "ALICE@example.com"},

		// --- Structure invalid ---
		{"missing_at", "aliceexample.com", false, ""},
		{"two_ats", "alice@@example.com", false, ""},
		{"at_at_start", "@example.com", false, ""},
		{"at_at_end", "alice@", false, ""},
		{"empty", "", false, ""},
		{"spaces_only", "   ", false, ""},

		// --- Spaces inside (ASCII space) ---
		{"space_in_local", "al ice@example.com", false, ""},
		{"space_in_domain", "alice@exa mple.com", false, ""},
		{"space_around_at_left", "alice @example.com", false, ""},
		{"space_around_at_right", "alice@ example.com", false, ""},
		{"space_around_at_both", "alice @ example.com", false, ""},

		// --- Unicode whitespace inside (should be rejected) ---
		{"tab_in_local", "ali\tce@example.com", false, ""},
		{"tab_in_domain", "alice@exa\tmple.com", false, ""},
		{"newline_in_local", "ali\nce@example.com", false, ""},
		{"newline_in_domain", "alice@exa\nmple.com", false, ""},
		{"carriage_return_in_domain", "alice@exa\rmple.com", false, ""},
		{"non_breaking_space_in_domain", "alice@exa\u00A0mple.com", false, ""}, // NBSP
		{"ideographic_space_in_domain", "alice@exa\u3000mple.com", false, ""},  // U+3000

		// --- Control characters (should be rejected) ---
		{"nul_in_local", "ali\x00ce@example.com", false, ""},
		{"nul_in_domain", "alice@exa\x00mple.com", false, ""},
		{"bell_in_domain", "alice@exa\x07mple.com", false, ""},

		// --- Leading/trailing whitespace that TrimSpace should remove (still ok) ---
		{"outer_tabs_ok", "\talice@Example.com\t", true, "alice@example.com"},
		{"outer_newlines_ok", "\nalice@Example.com\n", true, "alice@example.com"},

		// --- Multiple @ with surrounding whitespace ---
		{"two_ats_with_spaces", " alice@@Example.com ", false, ""},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := NormalizeEmail(tc.in)
			if tc.ok {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if got != tc.want {
					t.Fatalf("got %q want %q", got, tc.want)
				}
				// sanity: normalized result should not have outer whitespace
				if got != strings.TrimSpace(got) {
					t.Fatalf("normalized email has outer whitespace: %q", got)
				}
				// sanity: must contain exactly one '@' on success
				if strings.Count(got, "@") != 1 {
					t.Fatalf("normalized email must contain exactly one '@': %q", got)
				}
			} else {
				if err == nil {
					t.Fatalf("expected error, got nil (got=%q)", got)
				}
			}
		})
	}
}

func TestEX02_CountRunesCategories(t *testing.T) {
	cases := []struct {
		name                    string
		in                      string
		letters, digits, spaces int
	}{
		{"empty", "", 0, 0, 0},
		{"ascii_letters", "abc", 3, 0, 0},
		{"ascii_mixed", "a1b2c3", 3, 3, 0},
		{"ascii_spaces", "a b c", 3, 0, 2},

		{"cjk_and_digits", "你好 123", 2, 3, 1},
		{"greek_and_digits", "ΑΒΓ123", 3, 3, 0},
		{"whitespace_tab_nl_space", "\t\n ", 0, 0, 3},

		{"emoji_zwj_sequence", "👩‍💻", 0, 0, 0},
		{"latin_accented", "é", 1, 0, 0},

		// Combining mark: "e" + U+0301 (accent). Only 'e' is a letter.
		{"combining_mark", "e\u0301", 1, 0, 0},

		// Unicode spaces
		{"nbsp_space", "a\u00A0b", 2, 0, 1},
		{"ideographic_space", "a\u3000b", 2, 0, 1},

		// Unicode digits (Arabic-Indic)
		{"arabic_indic_digits", "١٢٣", 0, 3, 0},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l, d, sp := CountRunesCategories(tc.in)
			if l != tc.letters || d != tc.digits || sp != tc.spaces {
				t.Fatalf("got (letters=%d,digits=%d,spaces=%d) want (letters=%d,digits=%d,spaces=%d)",
					l, d, sp, tc.letters, tc.digits, tc.spaces)
			}
		})
	}
}

func TestEX02_SubstringRunes(t *testing.T) {
	cases := []struct {
		name          string
		s             string
		start, length int
		ok            bool
		want          string
	}{
		// --- Basic ---
		{"basic_hé", "héllo世界", 0, 2, true, "hé"},
		{"middle_él", "héllo世界", 1, 2, true, "él"},
		{"ascii_only", "héllo世界", 2, 3, true, "llo"},
		{"cjk_tail", "héllo世界", 5, 2, true, "世界"},
		{"empty_slice_at_end", "héllo世界", 7, 0, true, ""},
		{"full_string", "héllo世界", 0, 7, true, "héllo世界"},
		{"zero_length_mid", "héllo世界", 3, 0, true, ""},

		// --- Invalid bounds ---
		{"neg_start", "héllo世界", -1, 1, false, ""},
		{"neg_length", "héllo世界", 0, -1, false, ""},
		{"start_gt_len", "héllo世界", 8, 0, false, ""},      // start > n
		{"start_eq_len_len1", "héllo世界", 7, 1, false, ""}, // start==n but length>0
		{"start_plus_len_oob", "héllo世界", 6, 2, false, ""},

		// --- Empty input ---
		{"empty_string_ok", "", 0, 0, true, ""},
		{"empty_string_oob", "", 0, 1, false, ""},

		// --- Combining mark: "e\u0301" are 2 runes; substring by runes is precise ---
		{"combining_first_rune", "e\u0301x", 0, 1, true, "e"},
		{"combining_second_rune", "e\u0301x", 1, 1, true, "\u0301"},

		// --- Emoji sequence: rune slicing doesn't equal glyph slicing (still correct rune-wise) ---
		{"emoji_zwj_first_rune", "👩‍💻x", 0, 1, true, "👩"},
		{"emoji_zwj_full_three_runes", "👩‍💻x", 0, 3, true, "👩‍💻"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := SubstringRunes(tc.s, tc.start, tc.length)
			if tc.ok {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if got != tc.want {
					t.Fatalf("got %q want %q", got, tc.want)
				}
			} else {
				if err == nil {
					t.Fatalf("expected error, got nil (got=%q)", got)
				}
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
