package slug

import "strings"

// Make converts the input into a URL-friendly slug: lowercase, with any run of
// non-alphanumeric characters collapsed into a single "-", trimmed of leading
// and trailing hyphens.
func Make(in string) string {
	in = strings.ToLower(in)

	var b strings.Builder
	prevDash := false
	for _, r := range in {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
			prevDash = false
		} else if !prevDash {
			b.WriteByte('-')
			prevDash = true
		}
	}

	return strings.Trim(b.String(), "-")
}
