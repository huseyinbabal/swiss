package sqlfmt

import (
	"strings"
	"testing"
)

func TestFormat(t *testing.T) {
	res, err := Format("select a from t where x=1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(res, "SELECT") {
		t.Fatalf("expected SELECT in output:\n%s", res)
	}
	if !strings.Contains(res, "\nFROM") {
		t.Fatalf("expected newline before FROM in output:\n%s", res)
	}
}

func TestMinify(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"newlines", "select\n  a\nfrom t", "select a from t"},
		{"spaces", "a    b   c", "a b c"},
		{"trim", "  x  ", "x"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Minify(tt.in)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if res != tt.want {
				t.Fatalf("got %q, want %q", res, tt.want)
			}
		})
	}
}
