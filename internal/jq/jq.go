package jq

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type step struct {
	key   string
	index int
	isKey bool
}

func tokenize(path string) ([]step, error) {
	path = strings.TrimPrefix(path, ".")
	var steps []step
	if path == "" {
		return steps, nil
	}
	parts := strings.Split(path, ".")
	for _, part := range parts {
		if part == "" {
			continue
		}
		// Split off any trailing [n] indices.
		name := part
		if i := strings.IndexByte(part, '['); i >= 0 {
			name = part[:i]
		}
		if name != "" {
			steps = append(steps, step{key: name, isKey: true})
		}
		rest := part[len(name):]
		for len(rest) > 0 {
			if rest[0] != '[' {
				return nil, fmt.Errorf("invalid path segment: %q", part)
			}
			end := strings.IndexByte(rest, ']')
			if end < 0 {
				return nil, fmt.Errorf("invalid path segment: %q", part)
			}
			idxStr := rest[1:end]
			idx, err := strconv.Atoi(idxStr)
			if err != nil {
				return nil, fmt.Errorf("invalid index: %q", idxStr)
			}
			steps = append(steps, step{index: idx, isKey: false})
			rest = rest[end+1:]
		}
	}
	return steps, nil
}

// Query parses value as JSON and walks the given path, returning the
// resulting sub-value re-marshalled as indented JSON.
func Query(path, value string) (string, error) {
	var data interface{}
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return "", err
	}
	steps, err := tokenize(path)
	if err != nil {
		return "", err
	}
	cur := data
	for _, s := range steps {
		if s.isKey {
			m, ok := cur.(map[string]interface{})
			if !ok {
				return "", errors.New("path not found")
			}
			v, ok := m[s.key]
			if !ok {
				return "", errors.New("path not found")
			}
			cur = v
		} else {
			arr, ok := cur.([]interface{})
			if !ok {
				return "", errors.New("path not found")
			}
			if s.index < 0 || s.index >= len(arr) {
				return "", errors.New("path not found")
			}
			cur = arr[s.index]
		}
	}
	b, err := json.MarshalIndent(cur, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
