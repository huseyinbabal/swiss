package ip

import (
	"strings"
	"testing"
)

func TestCIDR(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{"v4", "192.168.1.0/24", "192.168.1.255", false},
		{"invalid", "not-a-cidr", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CIDR(tt.in)
			if (err != nil) != tt.wantErr {
				t.Fatalf("CIDR() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !strings.Contains(got, tt.want) {
				t.Errorf("CIDR() = %q, want contains %q", got, tt.want)
			}
		})
	}
}

func TestRoundTrip(t *testing.T) {
	tests := []string{"10.0.0.1", "192.168.1.1", "0.0.0.0", "255.255.255.255"}
	for _, ip := range tests {
		t.Run(ip, func(t *testing.T) {
			n, err := ToInt(ip)
			if err != nil {
				t.Fatalf("ToInt() error = %v", err)
			}
			back, err := FromInt(n)
			if err != nil {
				t.Fatalf("FromInt() error = %v", err)
			}
			if back != ip {
				t.Errorf("round trip = %q, want %q", back, ip)
			}
		})
	}
}

func TestToIntErrors(t *testing.T) {
	if _, err := ToInt("not-an-ip"); err == nil {
		t.Errorf("expected error for invalid IP")
	}
}
