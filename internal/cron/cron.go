package cron

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type fieldSpec struct {
	name   string
	min    int
	max    int
	values map[int]bool
	star   bool
	step   int // 0 if not a pure */n
}

var fieldDefs = []struct {
	name string
	min  int
	max  int
}{
	{"Minute", 0, 59},
	{"Hour", 0, 23},
	{"Day of month", 1, 31},
	{"Month", 1, 12},
	{"Day of week", 0, 6},
}

func parse(expr string) ([]fieldSpec, error) {
	parts := strings.Fields(strings.TrimSpace(expr))
	if len(parts) != 5 {
		return nil, fmt.Errorf("expected 5 fields, got %d", len(parts))
	}
	specs := make([]fieldSpec, 5)
	for i, p := range parts {
		def := fieldDefs[i]
		fs, err := parseField(p, def.name, def.min, def.max)
		if err != nil {
			return nil, err
		}
		specs[i] = fs
	}
	return specs, nil
}

func parseField(s, name string, min, max int) (fieldSpec, error) {
	fs := fieldSpec{name: name, min: min, max: max, values: map[int]bool{}}
	for _, token := range strings.Split(s, ",") {
		if err := parseToken(token, &fs); err != nil {
			return fs, err
		}
	}
	return fs, nil
}

func parseToken(token string, fs *fieldSpec) error {
	step := 1
	rangePart := token
	if idx := strings.IndexByte(token, '/'); idx >= 0 {
		rangePart = token[:idx]
		stepStr := token[idx+1:]
		n, err := strconv.Atoi(stepStr)
		if err != nil || n <= 0 {
			return fmt.Errorf("%s: invalid step %q", fs.name, stepStr)
		}
		step = n
	}

	var lo, hi int
	if rangePart == "*" {
		lo, hi = fs.min, fs.max
		if step != 1 {
			fs.step = step
		} else {
			fs.star = true
		}
	} else if dash := strings.IndexByte(rangePart, '-'); dash >= 0 {
		a, err := strconv.Atoi(rangePart[:dash])
		if err != nil {
			return fmt.Errorf("%s: invalid value %q", fs.name, rangePart[:dash])
		}
		b, err := strconv.Atoi(rangePart[dash+1:])
		if err != nil {
			return fmt.Errorf("%s: invalid value %q", fs.name, rangePart[dash+1:])
		}
		lo, hi = a, b
	} else {
		v, err := strconv.Atoi(rangePart)
		if err != nil {
			return fmt.Errorf("%s: invalid value %q", fs.name, rangePart)
		}
		lo, hi = v, v
	}

	if lo < fs.min || hi > fs.max || lo > hi {
		return fmt.Errorf("%s: value out of range %d-%d", fs.name, fs.min, fs.max)
	}
	for v := lo; v <= hi; v += step {
		fs.values[v] = true
	}
	return nil
}

func (fs fieldSpec) matches(v int) bool {
	if fs.star {
		return true
	}
	if fs.step > 0 {
		return (v-fs.min)%fs.step == 0
	}
	return fs.values[v]
}

func (fs fieldSpec) describe() string {
	if fs.star {
		return "every"
	}
	if fs.step > 0 {
		return fmt.Sprintf("every %d", fs.step)
	}
	vals := make([]int, 0, len(fs.values))
	for v := fs.min; v <= fs.max; v++ {
		if fs.values[v] {
			vals = append(vals, v)
		}
	}
	parts := make([]string, len(vals))
	for i, v := range vals {
		parts[i] = strconv.Itoa(v)
	}
	return "at " + strings.Join(parts, ",")
}

// Explain returns a human-readable, field-by-field description of a cron expression.
func Explain(expr string) (string, error) {
	specs, err := parse(expr)
	if err != nil {
		return "", err
	}
	lines := make([]string, len(specs))
	for i, fs := range specs {
		lines[i] = fmt.Sprintf("%s: %s", fs.name, fs.describe())
	}
	return strings.Join(lines, "\n"), nil
}

// Next returns the next count times (from now) that match the cron expression,
// formatted as RFC3339 lines.
func Next(expr string, count int) (string, error) {
	specs, err := parse(expr)
	if err != nil {
		return "", err
	}
	if count <= 0 {
		return "", fmt.Errorf("count must be positive")
	}

	t := time.Now().Truncate(time.Minute).Add(time.Minute)
	var results []string
	maxIter := 525600 * 2
	for i := 0; i < maxIter && len(results) < count; i++ {
		if matchTime(specs, t) {
			results = append(results, t.Format(time.RFC3339))
		}
		t = t.Add(time.Minute)
	}
	if len(results) < count {
		return "", fmt.Errorf("could not find %d matches within search window", count)
	}
	return strings.Join(results, "\n"), nil
}

func matchTime(specs []fieldSpec, t time.Time) bool {
	return specs[0].matches(t.Minute()) &&
		specs[1].matches(t.Hour()) &&
		specs[2].matches(t.Day()) &&
		specs[3].matches(int(t.Month())) &&
		specs[4].matches(int(t.Weekday()))
}
