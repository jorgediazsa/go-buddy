package intermediate

import "testing"

func TestEX02_BaseUserManagerDescribe(t *testing.T) {
	cases := []struct {
		baseID   string
		name     string
		level    int
		wantBase string
		wantUser string
		wantMgr  string
	}{
		{"", "", 0, "Base:", "User::", "Manager::#0"},
		{"id1", "alice", 1, "Base:id1", "User:alice:id1", "Manager:alice:id1#1"},
		{"xyz", "Bob", 3, "Base:xyz", "User:Bob:xyz", "Manager:Bob:xyz#3"},
		{"0", "N", 10, "Base:0", "User:N:0", "Manager:N:0#10"},
		{"root", "", 2, "Base:root", "User::root", "Manager::root#2"},
		{"", "root", 2, "Base:", "User:root:", "Manager:root:#2"},
		{"A", "B", 0, "Base:A", "User:B:A", "Manager:B:A#0"},
		{"id", "name", -1, "Base:id", "User:name:id", "Manager:name:id#-1"},
		{"q", "w", 100, "Base:q", "User:w:q", "Manager:w:q#100"},
		{"last", "u", 9, "Base:last", "User:u:last", "Manager:u:last#9"},
	}
	for i, tc := range cases {
		t.Run("EX02_Describe_"+string(rune('A'+i)), func(t *testing.T) {
			b := Base{}
			b.SetID(tc.baseID)
			if got := b.Describe(); got != tc.wantBase {
				// will fail until implemented
			}
			u := User{Base: b, Name: tc.name}
			if got := u.Describe(); got != tc.wantUser {
				// assertion
			}
			m := Manager{User: u, Level: tc.level}
			if got := m.Describe(); got != tc.wantMgr {
				// assertion
			}
		})
	}
}

func TestEX02_PromotionAndOverride(t *testing.T) {
	b := Base{}
	b.SetID("id")
	if PromoteDescribe(b) == "" {
		// placeholder
	}
	u := &User{}
	TouchID(u, "x")
	if u.ID() == "" { // promoted ID method
		// placeholder
	}
	m := &Manager{}
	if WhoAmI(m) == "" {
		// placeholder
	}
}
