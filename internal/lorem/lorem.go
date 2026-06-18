package lorem

import "strings"

var bank = []string{
	"lorem", "ipsum", "dolor", "sit", "amet", "consectetur",
	"adipiscing", "elit", "sed", "do", "eiusmod", "tempor",
	"incididunt", "ut", "labore", "et", "dolore", "magna",
	"aliqua", "enim", "ad", "minim", "veniam", "quis",
	"nostrud", "exercitation", "ullamco", "laboris",
}

// Words returns n space-joined words, cycling the word bank deterministically.
func Words(n int) string {
	if n <= 0 {
		return ""
	}
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = bank[i%len(bank)]
	}
	return strings.Join(words, " ")
}

// capitalize upper-cases the first letter of a word.
func capitalize(w string) string {
	if w == "" {
		return ""
	}
	return strings.ToUpper(w[:1]) + w[1:]
}

// sentence builds a single capitalized sentence of 6-12 words, deterministically
// based on the provided index.
func sentence(idx int) string {
	count := 6 + (idx % 7) // 6..12
	words := make([]string, count)
	for i := 0; i < count; i++ {
		words[i] = bank[(idx+i)%len(bank)]
	}
	words[0] = capitalize(words[0])
	return strings.Join(words, " ") + "."
}

// Sentences returns n sentences, each capitalized and ending in a period,
// joined by spaces.
func Sentences(n int) string {
	if n <= 0 {
		return ""
	}
	sentences := make([]string, n)
	for i := 0; i < n; i++ {
		sentences[i] = sentence(i)
	}
	return strings.Join(sentences, " ")
}

// Paragraphs returns n paragraphs, each made of 3-5 sentences, separated by
// blank lines.
func Paragraphs(n int) string {
	if n <= 0 {
		return ""
	}
	paragraphs := make([]string, n)
	idx := 0
	for p := 0; p < n; p++ {
		count := 3 + (p % 3) // 3..5
		sentences := make([]string, count)
		for i := 0; i < count; i++ {
			sentences[i] = sentence(idx)
			idx++
		}
		paragraphs[p] = strings.Join(sentences, " ")
	}
	return strings.Join(paragraphs, "\n\n")
}
