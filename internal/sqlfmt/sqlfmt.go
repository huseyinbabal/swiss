package sqlfmt

import (
	"regexp"
	"strings"
)

var wsRe = regexp.MustCompile(`\s+`)

// keywords to uppercase when they appear as whole words.
var keywords = []string{
	"SELECT", "FROM", "WHERE", "AND", "OR", "JOIN", "LEFT", "RIGHT",
	"INNER", "OUTER", "ON", "GROUP BY", "ORDER BY", "HAVING", "LIMIT",
	"INSERT", "INTO", "VALUES", "UPDATE", "SET", "DELETE",
}

// major clause keywords that get a newline before them.
var clauses = []string{
	"LEFT JOIN", "RIGHT JOIN", "INNER JOIN",
	"SELECT", "FROM", "WHERE", "GROUP BY", "ORDER BY", "HAVING", "LIMIT", "JOIN",
}

// Format is a lightweight SQL formatter: it uppercases known keywords and
// places newlines before major clause keywords.
func Format(in string) (string, error) {
	s := strings.TrimSpace(wsRe.ReplaceAllString(in, " "))

	for _, kw := range keywords {
		re := regexp.MustCompile(`(?i)\b` + regexp.QuoteMeta(kw) + `\b`)
		s = re.ReplaceAllString(s, kw)
	}

	for _, cl := range clauses {
		re := regexp.MustCompile(`\b` + regexp.QuoteMeta(cl) + `\b`)
		s = re.ReplaceAllString(s, "\n"+cl)
	}

	s = strings.TrimSpace(s)
	// normalize lines: trim trailing spaces and drop a leading blank line.
	lines := strings.Split(s, "\n")
	out := make([]string, 0, len(lines))
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		out = append(out, l)
	}

	return strings.Join(out, "\n"), nil
}

// Minify collapses all whitespace (including newlines) to single spaces.
func Minify(in string) (string, error) {
	return strings.TrimSpace(wsRe.ReplaceAllString(in, " ")), nil
}
