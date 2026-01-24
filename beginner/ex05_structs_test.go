package beginner

import "testing"

func TestEX05_ValidatePerson(t *testing.T) {
	cases := []struct {
		name      string
		in        Person
		ok        bool
		wantName  string
		wantEmail string
	}{
		{
			name:      "valid_trims_name_and_lowercases_domain",
			in:        Person{Name: " Alice ", Email: "alice@Example.com", BirthYear: 1990},
			ok:        true,
			wantName:  "Alice",
			wantEmail: "alice@example.com",
		},
		{
			name:      "valid_boundary_1900",
			in:        Person{Name: "Bob", Email: "bob@example.com", BirthYear: 1900},
			ok:        true,
			wantName:  "Bob",
			wantEmail: "bob@example.com",
		},
		{
			name:      "valid_boundary_2100",
			in:        Person{Name: "Eve", Email: "eve@example.com", BirthYear: 2100},
			ok:        true,
			wantName:  "Eve",
			wantEmail: "eve@example.com",
		},
		{
			name:      "valid_unicode_domain",
			in:        Person{Name: "Ana", Email: "ana@例え.テスト", BirthYear: 2001},
			ok:        true,
			wantName:  "Ana",
			wantEmail: "ana@例え.テスト",
		},
		{
			name: "invalid_empty_name",
			in:   Person{Name: "", Email: "nobody@example.com", BirthYear: 1990},
			ok:   false,
		},
		{
			name: "invalid_whitespace_name",
			in:   Person{Name: " \t\n ", Email: "nobody@example.com", BirthYear: 1990},
			ok:   false,
		},
		{
			name: "invalid_empty_email",
			in:   Person{Name: "Zoe", Email: "", BirthYear: 1990},
			ok:   false,
		},
		{
			name: "invalid_email_missing_at",
			in:   Person{Name: "Zoe", Email: "zoeexample.com", BirthYear: 1990},
			ok:   false,
		},
		{
			name: "invalid_email_missing_domain",
			in:   Person{Name: "Zoe", Email: "zoe@", BirthYear: 1990},
			ok:   false,
		},
		{
			name: "invalid_year_too_low",
			in:   Person{Name: "Zoe", Email: "zoe@example.com", BirthYear: 1800},
			ok:   false,
		},
		{
			name: "invalid_year_too_high",
			in:   Person{Name: "Zoe", Email: "zoe@example.com", BirthYear: 2200},
			ok:   false,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := ValidatePerson(tc.in)

			if tc.ok {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if got.Name != tc.wantName {
					t.Fatalf("name mismatch: got %q want %q", got.Name, tc.wantName)
				}
				if got.Email != tc.wantEmail {
					t.Fatalf("email mismatch: got %q want %q", got.Email, tc.wantEmail)
				}
				if got.BirthYear != tc.in.BirthYear {
					t.Fatalf("birthYear mismatch: got %d want %d", got.BirthYear, tc.in.BirthYear)
				}
			} else {
				if err == nil {
					t.Fatalf("expected error, got nil (got=%+v)", got)
				}
			}
		})
	}
}

func TestEX05_UpdateEmail(t *testing.T) {
	t.Run("success_normalizes_and_updates", func(t *testing.T) {
		p := Person{Name: "Bob", Email: "bob@EXAMPLE.com", BirthYear: 1980}

		if err := UpdateEmail(&p, "Bob@Example.COM"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if p.Email != "Bob@example.com" { // local preserved, domain lowercased (según tus reglas)
			t.Fatalf("email not normalized/updated: got %q", p.Email)
		}
	})

	t.Run("invalid_email_does_not_mutate", func(t *testing.T) {
		p := Person{Name: "Bob", Email: "bob@example.com", BirthYear: 1980}
		before := p.Email

		if err := UpdateEmail(&p, "bobexample.com"); err == nil {
			t.Fatalf("expected error for invalid email")
		}
		if p.Email != before {
			t.Fatalf("email mutated on failure: got %q want %q", p.Email, before)
		}
	})

	t.Run("nil_person_returns_error", func(t *testing.T) {
		if err := UpdateEmail(nil, "x@y"); err == nil {
			t.Fatalf("expected error for nil person")
		}
	})

	t.Run("unicode_domain_ok", func(t *testing.T) {
		p := Person{Name: "Ana", Email: "ana@example.com", BirthYear: 2001}

		if err := UpdateEmail(&p, "ana@例え.テスト"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if p.Email != "ana@例え.テスト" {
			t.Fatalf("unexpected email: got %q", p.Email)
		}
	})

	t.Run("rejects_spaces_and_controls", func(t *testing.T) {
		p := Person{Name: "Bob", Email: "bob@example.com", BirthYear: 1980}
		before := p.Email

		if err := UpdateEmail(&p, " bob@example.com "); err != nil {
			t.Fatalf("unexpected error")
		}
		if p.Email != before {
			t.Fatalf("email mutated on failure: got %q want %q", p.Email, before)
		}
	})
}

func TestEX05_AgeAtYear(t *testing.T) {
	t.Run("basic_cases", func(t *testing.T) {
		p := Person{Name: "A", Email: "a@b.c", BirthYear: 2000}

		cases := []struct {
			name string
			year int
			ok   bool
			want int
		}{
			{"same_year", 2000, true, 0},
			{"next_year", 2001, true, 1},
			{"upper_bound", 2100, true, 100},
			{"before_birth", 1999, false, 0},
			{"below_min_range", 1800, false, 0},
			{"above_max_range", 2200, false, 0},
			{"min_range_but_before_birth", 1900, false, 0}, // still before birth year
		}

		for _, tc := range cases {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				got, err := AgeAtYear(p, tc.year)
				if tc.ok {
					if err != nil {
						t.Fatalf("unexpected error: %v", err)
					}
					if got != tc.want {
						t.Fatalf("got %d want %d", got, tc.want)
					}
				} else {
					if err == nil {
						t.Fatalf("expected error, got nil (got=%d)", got)
					}
				}
			})
		}
	})

	t.Run("range_check_independent_of_birthyear", func(t *testing.T) {
		// This catches implementations that only compare against BirthYear
		p := Person{Name: "Old", Email: "o@b.c", BirthYear: 1900}

		if _, err := AgeAtYear(p, 1899); err == nil {
			t.Fatalf("expected error for targetYear < 1900")
		}
		if _, err := AgeAtYear(p, 2101); err == nil {
			t.Fatalf("expected error for targetYear > 2100")
		}
		if got, err := AgeAtYear(p, 1900); err != nil || got != 0 {
			t.Fatalf("expected 0 at birth year, got=%d err=%v", got, err)
		}
	})
}
