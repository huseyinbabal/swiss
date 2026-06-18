package regex

import (
	"fmt"
	"regexp"
	"strings"
)

// Match compiles pattern and finds all matches (with capture groups) in value.
// Returns "no match" if nothing matches; errors on an invalid pattern.
func Match(pattern, value string) (string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}

	matches := re.FindAllStringSubmatch(value, -1)
	if len(matches) == 0 {
		return "no match", nil
	}

	var sb strings.Builder
	for i, m := range matches {
		sb.WriteString(fmt.Sprintf("match %d: %q\n", i+1, m[0]))
		for g := 1; g < len(m); g++ {
			sb.WriteString(fmt.Sprintf("  group %d: %q\n", g, m[g]))
		}
	}

	return strings.TrimRight(sb.String(), "\n"), nil
}

// Replace compiles pattern and replaces all matches in value with repl.
func Replace(pattern, value, repl string) (string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	return re.ReplaceAllString(value, repl), nil
}
