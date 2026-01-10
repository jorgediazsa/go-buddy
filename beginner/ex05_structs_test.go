package beginner

import "testing"

func TestEX05_ValidatePerson(t *testing.T) {
	cases := []struct {
		in Person
		ok bool
	}{
		{Person{" Alice ", "alice@Example.com", 1990}, true},
		{Person{"Bob", "bob@example.com", 1900}, true},
		{Person{"Eve", "eve@example.com", 2100}, true},
		{Person{"", "nobody@example.com", 1990}, false},
		{Person{"Zoe", "", 1990}, false},
		{Person{"Zoe", "zoeexample.com", 1990}, false},
		{Person{"Zoe", "zoe@", 1990}, false},
		{Person{"Zoe", "zoe@example.com", 1800}, false},
		{Person{"Zoe", "zoe@example.com", 2200}, false},
		{Person{"Ana", "ana@例え.テスト", 2001}, true},
	}
	for i, tc := range cases {
		t.Run("EX05_ValidatePerson_"+string(rune('A'+i)), func(t *testing.T) {
			got, err := ValidatePerson(tc.in)
			if tc.ok && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if tc.ok {
				if got.Name == "" {
					t.Fatalf("name should be non-empty")
				}
				if got.Email == "" {
					t.Fatalf("email should be normalized")
				}
				if got.BirthYear < 1900 || got.BirthYear > 2100 {
					t.Fatalf("year out of range")
				}
			}
		})
	}
}

func TestEX05_UpdateEmail(t *testing.T) {
	p := Person{Name: "Bob", Email: "bob@EXAMPLE.com", BirthYear: 1980}
	if err := UpdateEmail(&p, "bob@example.com"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// invalid
	if err := UpdateEmail(&p, "bobexample.com"); err == nil {
		t.Fatalf("expected error for invalid email")
	}
	if err := UpdateEmail(nil, "x@y"); err == nil {
		t.Fatalf("expected error for nil person")
	}
}

func TestEX05_AgeAtYear(t *testing.T) {
	p := Person{Name: "A", Email: "a@b.c", BirthYear: 2000}
	cases := []struct {
		year int
		ok   bool
		want int
	}{
		{2000, true, 0},
		{2001, true, 1},
		{2100, true, 100},
		{1999, false, 0},
		{1800, false, 0},
		{2200, false, 0},
	}
	for i, tc := range cases {
		t.Run("EX05_AgeAtYear_"+string(rune('A'+i)), func(t *testing.T) {
			got, err := AgeAtYear(p, tc.year)
			if tc.ok && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error")
			}
			if tc.ok && got != tc.want {
				t.Fatalf("got %d want %d", got, tc.want)
			}
		})
	}
}
