package timeconv

import (
	"fmt"
	"strconv"
	"time"
)

// Now returns the current time in several formats.
func Now() (string, error) {
	t := time.Now()
	return fmt.Sprintf("Unix: %d\nUnixMilli: %d\nISO: %s\nUTC: %s",
		t.Unix(),
		t.UnixMilli(),
		t.Format(time.RFC3339),
		t.UTC().Format(time.RFC3339),
	), nil
}

// ToUnix parses a date string and returns its unix seconds as a string.
func ToUnix(in string) (string, error) {
	layouts := []string{
		time.RFC3339,
		"2006-01-02 15:04:05",
		"2006-01-02",
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, in); err == nil {
			return strconv.FormatInt(t.Unix(), 10), nil
		}
	}
	return "", fmt.Errorf("could not parse date: %q", in)
}

// FromUnix parses unix seconds and returns an RFC3339 UTC string.
func FromUnix(in string) (string, error) {
	n, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return "", fmt.Errorf("invalid unix seconds: %q", in)
	}
	return time.Unix(n, 0).UTC().Format(time.RFC3339), nil
}
