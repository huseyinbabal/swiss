package hex

import "testing"

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"simple", "abc", "616263"},
		{"empty", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encode(tt.in)
			if err != nil {
				t.Fatalf("Encode() unexpected error = %v", err)
			}
			if got != tt.want {
				t.Errorf("Encode(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{"simple", "616263", "abc", false},
		{"empty", "", "", false},
		{"bad hex", "xyz", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.in)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("Decode(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
