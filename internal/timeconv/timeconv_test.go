package timeconv

import (
	"strings"
	"testing"
)

func TestNow(t *testing.T) {
	got, err := Now()
	if err != nil {
		t.Fatalf("Now() err=%v", err)
	}
	if got == "" {
		t.Errorf("Now() returned empty string")
	}
}

func TestFromUnix(t *testing.T) {
	tests := []struct {
		name       string
		in         string
		wantPrefix string
		wantErr    bool
	}{
		{"epoch", "0", "1970-01-01", false},
		{"non numeric", "abc", "", true},
		{"empty", "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromUnix(tt.in)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromUnix(%q) err=%v wantErr=%v", tt.in, err, tt.wantErr)
			}
			if !tt.wantErr && !strings.HasPrefix(got, tt.wantPrefix) {
				t.Errorf("FromUnix(%q)=%q want prefix %q", tt.in, got, tt.wantPrefix)
			}
		})
	}
}

func TestToUnixRoundTrip(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		wantErr bool
	}{
		{"date", "1970-01-01", false},
		{"datetime", "2000-01-02 03:04:05", false},
		{"invalid", "not-a-date", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unix, err := ToUnix(tt.in)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ToUnix(%q) err=%v wantErr=%v", tt.in, err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			back, err := FromUnix(unix)
			if err != nil {
				t.Fatalf("FromUnix(%q) err=%v", unix, err)
			}
			datePart := strings.SplitN(tt.in, " ", 2)[0]
			if !strings.HasPrefix(back, datePart) {
				t.Errorf("round trip of %q got %q want prefix %q", tt.in, back, datePart)
			}
		})
	}
}
