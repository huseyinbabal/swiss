package slug

import "testing"

func TestMake(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"simple", "Hello World", "hello-world"},
		{"punctuation", "Hello, World!", "hello-world"},
		{"leading trailing", "  --Hello--  ", "hello"},
		{"multiple separators", "a   b___c", "a-b-c"},
		{"alphanumeric", "Go 1.26 Release", "go-1-26-release"},
		{"empty", "", ""},
		{"only symbols", "!!!", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Make(tt.in); got != tt.want {
				t.Errorf("Make(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
