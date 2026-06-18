package html

import (
	stdhtml "html"
	"regexp"
)

var tagRe = regexp.MustCompile(`<[^>]*>`)

// Escape converts special characters to their HTML entity equivalents.
func Escape(in string) (string, error) {
	return stdhtml.EscapeString(in), nil
}

// Unescape converts HTML entities back to their plain characters.
func Unescape(in string) (string, error) {
	return stdhtml.UnescapeString(in), nil
}

// Strip removes HTML tags from the input.
func Strip(in string) (string, error) {
	return tagRe.ReplaceAllString(in, ""), nil
}
