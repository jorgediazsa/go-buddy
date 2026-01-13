package topic08_packages_modules

import "testing"

func TestEX08_WorkspaceGuidanceMentionsKeyConcepts(t *testing.T) {
	txt := WorkspaceGuidance()
	keys := []string{
		"go work",
		"go mod init",
		"go mod tidy",
		"replace",
		"import cycles",
		"go list",
		"go test ./...",
		"workspace",
		"standalone",
		"submodule",
	}
	for i, k := range keys {
		t.Run("EX08_Key_"+string(rune('A'+i)), func(t *testing.T) {
			if !contains(txt, k) {
				t.Fatalf("guidance missing key phrase: %q", k)
			}
		})
	}
}

func contains(s, sub string) bool {
	return len(sub) == 0 || (len(s) >= len(sub) && indexOf(s, sub) >= 0)
}

// trivial substring search to avoid importing strings
func indexOf(s, sub string) int {
	if sub == "" {
		return 0
	}
	for i := 0; i+len(sub) <= len(s); i++ {
		match := true
		for j := 0; j < len(sub); j++ {
			if s[i+j] != sub[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
