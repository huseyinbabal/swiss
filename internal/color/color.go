package color

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ToRGB parses a hex color and returns "rgb(r, g, b)".
func ToRGB(in string) (string, error) {
	r, g, b, err := parseHex(in)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("rgb(%d, %d, %d)", r, g, b), nil
}

// ToHex parses "rgb(r, g, b)" or "r,g,b" and returns "#rrggbb".
func ToHex(in string) (string, error) {
	s := strings.TrimSpace(in)
	s = strings.TrimPrefix(s, "rgb(")
	s = strings.TrimPrefix(s, "rgba(")
	s = strings.TrimSuffix(s, ")")
	parts := strings.Split(s, ",")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid rgb input: %q", in)
	}
	vals := make([]int, 3)
	for i, p := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			return "", fmt.Errorf("invalid rgb component %q: %w", p, err)
		}
		if n < 0 || n > 255 {
			return "", fmt.Errorf("rgb component out of range: %d", n)
		}
		vals[i] = n
	}
	return fmt.Sprintf("#%02x%02x%02x", vals[0], vals[1], vals[2]), nil
}

// ToHSL parses a hex color and returns "hsl(h, s%, l%)".
func ToHSL(in string) (string, error) {
	ri, gi, bi, err := parseHex(in)
	if err != nil {
		return "", err
	}
	r := float64(ri) / 255.0
	g := float64(gi) / 255.0
	b := float64(bi) / 255.0

	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	l := (max + min) / 2.0

	var h, s float64
	if max == min {
		h = 0
		s = 0
	} else {
		d := max - min
		if l > 0.5 {
			s = d / (2.0 - max - min)
		} else {
			s = d / (max + min)
		}
		switch max {
		case r:
			h = (g - b) / d
			if g < b {
				h += 6
			}
		case g:
			h = (b-r)/d + 2
		case b:
			h = (r-g)/d + 4
		}
		h /= 6
	}

	hDeg := math.Round(h * 360)
	sPct := math.Round(s * 100)
	lPct := math.Round(l * 100)
	return fmt.Sprintf("hsl(%d, %d%%, %d%%)", int(hDeg), int(sPct), int(lPct)), nil
}

func parseHex(in string) (int, int, int, error) {
	s := strings.TrimSpace(in)
	s = strings.TrimPrefix(s, "#")
	s = strings.ToLower(s)
	if len(s) == 3 {
		s = string([]byte{s[0], s[0], s[1], s[1], s[2], s[2]})
	}
	if len(s) != 6 {
		return 0, 0, 0, fmt.Errorf("invalid hex color: %q", in)
	}
	r, err := strconv.ParseInt(s[0:2], 16, 0)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid hex color: %q", in)
	}
	g, err := strconv.ParseInt(s[2:4], 16, 0)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid hex color: %q", in)
	}
	b, err := strconv.ParseInt(s[4:6], 16, 0)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid hex color: %q", in)
	}
	return int(r), int(g), int(b), nil
}
