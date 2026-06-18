package gzip

import "testing"

func TestRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		in   string
	}{
		{"simple", "hello world"},
		{"empty", ""},
		{"unicode", "héllo, wörld"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := Compress(tt.in)
			if err != nil {
				t.Fatalf("Compress() unexpected error = %v", err)
			}
			got, err := Decompress(c)
			if err != nil {
				t.Fatalf("Decompress() unexpected error = %v", err)
			}
			if got != tt.in {
				t.Errorf("round trip = %q, want %q", got, tt.in)
			}
		})
	}
}

func TestDecompressError(t *testing.T) {
	tests := []struct {
		name string
		in   string
	}{
		{"bad base64", "!!!not base64!!!"},
		{"valid base64 not gzip", "aGVsbG8="},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Decompress(tt.in); err == nil {
				t.Errorf("Decompress(%q) expected error, got nil", tt.in)
			}
		})
	}
}
