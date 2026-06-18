package color

import "testing"

func TestToRGB(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{"white", "#ffffff", "rgb(255, 255, 255)", false},
		{"shorthand", "#fff", "rgb(255, 255, 255)", false},
		{"no hash", "ff0000", "rgb(255, 0, 0)", false},
		{"invalid", "#zzz", "", true},
		{"bad length", "#ff", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToRGB(tt.in)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ToRGB(%q) err=%v wantErr=%v", tt.in, err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("ToRGB(%q)=%q want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestToHex(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{"red rgb", "rgb(255,0,0)", "#ff0000", false},
		{"plain", "0,255,0", "#00ff00", false},
		{"out of range", "rgb(256,0,0)", "", true},
		{"invalid", "rgb(a,b,c)", "", true},
		{"too few", "1,2", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToHex(tt.in)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ToHex(%q) err=%v wantErr=%v", tt.in, err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("ToHex(%q)=%q want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestToHSL(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{"white", "#ffffff", "hsl(0, 0%, 100%)", false},
		{"black", "#000000", "hsl(0, 0%, 0%)", false},
		{"red", "#ff0000", "hsl(0, 100%, 50%)", false},
		{"invalid", "nope", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToHSL(tt.in)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ToHSL(%q) err=%v wantErr=%v", tt.in, err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("ToHSL(%q)=%q want %q", tt.in, got, tt.want)
			}
		})
	}
}
