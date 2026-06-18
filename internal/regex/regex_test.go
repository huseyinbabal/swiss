package regex

import (
	"strings"
	"testing"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		value    string
		wantErr  bool
		contains []string
	}{
		{
			name:     "groups",
			pattern:  `(\w+)@(\w+)`,
			value:    "user@host",
			contains: []string{"user@host", "user", "host"},
		},
		{
			name:     "no match",
			pattern:  `xyz`,
			value:    "abc",
			contains: []string{"no match"},
		},
		{
			name:    "invalid pattern",
			pattern: `(`,
			value:   "abc",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Match(tt.pattern, tt.value)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			for _, c := range tt.contains {
				if !strings.Contains(res, c) {
					t.Fatalf("expected %q in output:\n%s", c, res)
				}
			}
		})
	}
}

func TestReplace(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		value   string
		repl    string
		want    string
		wantErr bool
	}{
		{
			name:    "simple",
			pattern: `a`,
			value:   "banana",
			repl:    "o",
			want:    "bonono",
		},
		{
			name:    "invalid",
			pattern: `(`,
			value:   "x",
			repl:    "y",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Replace(tt.pattern, tt.value, tt.repl)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if res != tt.want {
				t.Fatalf("got %q, want %q", res, tt.want)
			}
		})
	}
}
