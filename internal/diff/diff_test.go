package diff

import (
	"strings"
	"testing"
)

func TestLines(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		contains []string
	}{
		{
			name:     "single change",
			a:        "x\ny\nz",
			b:        "x\nY\nz",
			contains: []string{"  x", "- y", "+ Y", "  z"},
		},
		{
			name:     "identical",
			a:        "a\nb",
			b:        "a\nb",
			contains: []string{"  a", "  b"},
		},
		{
			name:     "addition",
			a:        "a",
			b:        "a\nb",
			contains: []string{"  a", "+ b"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Lines(tt.a, tt.b)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			lastIdx := -1
			for _, c := range tt.contains {
				idx := strings.Index(res, c)
				if idx < 0 {
					t.Fatalf("expected %q in output:\n%s", c, res)
				}
				if idx < lastIdx {
					t.Fatalf("expected %q after previous marker, output:\n%s", c, res)
				}
				lastIdx = idx
			}
		})
	}
}
