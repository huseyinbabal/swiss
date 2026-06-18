package lorem

import (
	"strings"
	"testing"
)

func TestWords(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"five", 5, 5},
		{"zero", 0, 0},
		{"negative", -3, 0},
		{"cycle", len(bank) + 3, len(bank) + 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Words(tt.n)
			if tt.want == 0 {
				if got != "" {
					t.Errorf("Words(%d) = %q, want empty", tt.n, got)
				}
				return
			}
			if c := len(strings.Fields(got)); c != tt.want {
				t.Errorf("Words(%d) has %d words, want %d", tt.n, c, tt.want)
			}
		})
	}
}

func TestSentences(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{"two", 2},
		{"five", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sentences(tt.n)
			if c := strings.Count(got, "."); c != tt.n {
				t.Errorf("Sentences(%d) has %d periods, want %d", tt.n, c, tt.n)
			}
			if got != "" && got[0] < 'A' || got != "" && got[0] > 'Z' {
				t.Errorf("Sentences(%d) does not start capitalized: %q", tt.n, got)
			}
		})
	}
}

func TestSentencesEmpty(t *testing.T) {
	if got := Sentences(0); got != "" {
		t.Errorf("Sentences(0) = %q, want empty", got)
	}
}

func TestParagraphs(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{"one", 1},
		{"three", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Paragraphs(tt.n)
			blocks := strings.Split(got, "\n\n")
			if len(blocks) != tt.n {
				t.Errorf("Paragraphs(%d) produced %d blocks, want %d", tt.n, len(blocks), tt.n)
			}
			for _, b := range blocks {
				c := strings.Count(b, ".")
				if c < 3 || c > 5 {
					t.Errorf("paragraph has %d sentences, want 3-5: %q", c, b)
				}
			}
		})
	}
}

func TestParagraphsEmpty(t *testing.T) {
	if got := Paragraphs(0); got != "" {
		t.Errorf("Paragraphs(0) = %q, want empty", got)
	}
}

func TestDeterministic(t *testing.T) {
	if Paragraphs(2) != Paragraphs(2) {
		t.Error("Paragraphs is not deterministic")
	}
}
