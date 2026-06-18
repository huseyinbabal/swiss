package base58

import "testing"

func TestRoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "simple", input: "hello"},
		{name: "with spaces", input: "the quick brown fox"},
		{name: "leading zero byte", input: "\x00abc"},
		{name: "multiple leading zeros", input: "\x00\x00xy"},
		{name: "empty", input: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc, err := Encode(tt.input)
			if err != nil {
				t.Fatalf("Encode error: %v", err)
			}
			dec, err := Decode(enc)
			if err != nil {
				t.Fatalf("Decode error: %v", err)
			}
			if dec != tt.input {
				t.Fatalf("round trip mismatch: got %q, want %q", dec, tt.input)
			}
		})
	}
}

func TestDecodeInvalid(t *testing.T) {
	if _, err := Decode("0OIl"); err == nil {
		t.Fatalf("expected error for invalid base58 input, got nil")
	}
}
