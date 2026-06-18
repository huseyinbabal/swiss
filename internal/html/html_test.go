package html

import "testing"

func TestEscapeUnescapeRoundTrip(t *testing.T) {
	tests := []string{
		`<a href="x">link</a>`,
		`a & b < c > d`,
		`plain text`,
	}

	for _, in := range tests {
		t.Run(in, func(t *testing.T) {
			esc, err := Escape(in)
			if err != nil {
				t.Fatalf("escape error: %v", err)
			}
			out, err := Unescape(esc)
			if err != nil {
				t.Fatalf("unescape error: %v", err)
			}
			if out != in {
				t.Fatalf("round-trip failed: got %q, want %q", out, in)
			}
		})
	}
}

func TestStrip(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"anchor", `<a href="x">link</a>`, "link"},
		{"nested", `<p><b>hi</b></p>`, "hi"},
		{"no tags", "plain", "plain"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Strip(tt.in)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if res != tt.want {
				t.Fatalf("got %q, want %q", res, tt.want)
			}
		})
	}
}
