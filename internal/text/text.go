package text

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Count returns a multi-line report about the given input string.
func Count(in string) string {
	chars := utf8.RuneCountInString(in)
	bytes := len(in)
	words := len(strings.Fields(in))

	lines := strings.Count(in, "\n")
	if in != "" && !strings.HasSuffix(in, "\n") {
		lines++
	}

	readingTime := (words + 199) / 200

	var b strings.Builder
	fmt.Fprintf(&b, "Characters:   %d\n", chars)
	fmt.Fprintf(&b, "Bytes:        %d\n", bytes)
	fmt.Fprintf(&b, "Words:        %d\n", words)
	fmt.Fprintf(&b, "Lines:        %d\n", lines)
	fmt.Fprintf(&b, "Reading time: %d min", readingTime)
	return b.String()
}
