package textcase

import (
	"strings"
	"unicode"
)

// tokenize splits the input into lowercase words, breaking on spaces,
// underscores, hyphens and camelCase boundaries.
func tokenize(in string) []string {
	var words []string
	var current []rune

	flush := func() {
		if len(current) > 0 {
			words = append(words, strings.ToLower(string(current)))
			current = current[:0]
		}
	}

	runes := []rune(in)
	for i, r := range runes {
		switch {
		case r == ' ' || r == '_' || r == '-':
			flush()
		case unicode.IsUpper(r):
			// camelCase boundary: previous rune is lower/digit, or next rune is lower
			if i > 0 {
				prev := runes[i-1]
				if unicode.IsLower(prev) || unicode.IsDigit(prev) {
					flush()
				} else if unicode.IsUpper(prev) && i+1 < len(runes) && unicode.IsLower(runes[i+1]) {
					flush()
				}
			}
			current = append(current, r)
		default:
			current = append(current, r)
		}
	}
	flush()
	return words
}

func capitalize(w string) string {
	if w == "" {
		return ""
	}
	r := []rune(w)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// ToCamel converts the input to camelCase.
func ToCamel(in string) string {
	words := tokenize(in)
	var b strings.Builder
	for i, w := range words {
		if i == 0 {
			b.WriteString(w)
		} else {
			b.WriteString(capitalize(w))
		}
	}
	return b.String()
}

// ToPascal converts the input to PascalCase.
func ToPascal(in string) string {
	words := tokenize(in)
	var b strings.Builder
	for _, w := range words {
		b.WriteString(capitalize(w))
	}
	return b.String()
}

// ToSnake converts the input to snake_case.
func ToSnake(in string) string {
	return strings.Join(tokenize(in), "_")
}

// ToKebab converts the input to kebab-case.
func ToKebab(in string) string {
	return strings.Join(tokenize(in), "-")
}

// ToTitle converts the input to Title Case.
func ToTitle(in string) string {
	words := tokenize(in)
	out := make([]string, len(words))
	for i, w := range words {
		out[i] = capitalize(w)
	}
	return strings.Join(out, " ")
}

// ToUpper upper-cases the entire input.
func ToUpper(in string) string {
	return strings.ToUpper(in)
}

// ToLower lower-cases the entire input.
func ToLower(in string) string {
	return strings.ToLower(in)
}
