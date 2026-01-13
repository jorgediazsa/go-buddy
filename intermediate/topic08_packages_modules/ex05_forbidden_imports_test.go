package topic08_packages_modules

import "testing"

func TestEX05_NoForbiddenImports(t *testing.T) {
	srcs := []string{
		`package x
         import "fmt"
         import f "os"
         import (
             _ "net/http"
             y "strings"
         )
         var _ = f.Stdout
         var _ = y.Builder{}
        `,
		`package y
         import (
             "path/filepath"
             "strings"
         )
        `,
		`package z
         // no imports
        `,
		`package a
         import (
             // comment
             "net/http"
             alias "encoding/json"
         )
         var _ = alias.Valid{}
        `,
	}
	forbidden := []string{"net/http", "encoding/json", "fmt"}
	wants := [][]string{
		{"fmt", "net/http"},
		{},
		{},
		{"encoding/json", "net/http"},
	}
	for i, src := range srcs {
		t.Run("EX05_Forbidden_"+string(rune('A'+i)), func(t *testing.T) {
			got, err := NoForbiddenImports(src, forbidden)
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			want := wants[i]
			if len(got) != len(want) {
				t.Fatalf("len=%d want %d, got=%v want=%v", len(got), len(want), got, want)
			}
			for j := range got {
				if got[j] != want[j] {
					t.Fatalf("at %d got %q want %q", j, got[j], want[j])
				}
			}
		})
	}
}
