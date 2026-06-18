package cron

import (
	"strings"
	"testing"
)

func TestExplain(t *testing.T) {
	tests := []struct {
		name    string
		expr    string
		want    string
		wantErr bool
	}{
		{"every 5", "*/5 * * * *", "5", false},
		{"wrong count", "a b c", "", true},
		{"out of range", "99 * * * *", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Explain(tt.expr)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Explain() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !strings.Contains(got, tt.want) {
				t.Errorf("Explain() = %q, want contains %q", got, tt.want)
			}
		})
	}
}

func TestNext(t *testing.T) {
	tests := []struct {
		name    string
		expr    string
		count   int
		want    int
		wantErr bool
	}{
		{"every minute", "* * * * *", 3, 3, false},
		{"bad expr", "a b c", 3, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Next(tt.expr, tt.count)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Next() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			lines := strings.Split(strings.TrimSpace(got), "\n")
			if len(lines) != tt.want {
				t.Errorf("Next() returned %d lines, want %d", len(lines), tt.want)
			}
		})
	}
}
