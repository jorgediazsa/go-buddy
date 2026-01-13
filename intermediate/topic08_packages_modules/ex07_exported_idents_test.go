package topic08_packages_modules

import "testing"

func TestEX07_ListExportedIdents(t *testing.T) {
	srcs := []struct {
		src  string
		want []string
	}{
		{`package a
          const A = 1
          var B, c = 2, 3
          type T struct{}
          func F() {}
          func g() {}
        `, []string{"A", "B", "F", "T"}},
		{`package b
          import "fmt"
          const (
             X = 1
             y = 2
          )
          var (
             Z int
             zz int
          )
          type (
             AA int
             bb int
          )
          func H(){}
          func (t AA) Method(){}
        `, []string{"AA", "H", "X", "Z"}},
		{`package c
          // no exported
          var x int
          type t int
          func f(){}
        `, []string{}},
		{`package d
          const (
             A1 = iota
             A2
          )
          var (
             V1, V2 = 1,2
          )
          type (
             T1 struct{}
             T2 interface{}
          )
          func Fn(){}
        `, []string{"A1", "A2", "Fn", "T1", "T2", "V1", "V2"}},
		{`package e
          // Methods shouldn't add exported identifier names beyond func decl name
          type T struct{}
          func (T) M(){}
        `, []string{"T"}},
		{`package f
          const (
             ALPHA=1
          )
          var (
             BETA string
          )
          type GAMMA int
          func DELTA(){}
        `, []string{"ALPHA", "BETA", "DELTA", "GAMMA"}},
	}
	for i, tc := range srcs {
		t.Run("EX07_List_"+string(rune('A'+i)), func(t *testing.T) {
			got, err := ListExportedIdents(tc.src)
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			if len(got) != len(tc.want) {
				t.Fatalf("len=%d want %d, got=%v want=%v", len(got), len(tc.want), got, tc.want)
			}
			for j := range got {
				if got[j] != tc.want[j] {
					t.Fatalf("at %d got %q want %q", j, got[j], tc.want[j])
				}
			}
		})
	}
}
