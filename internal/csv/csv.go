package csv

import (
	"bytes"
	encodingcsv "encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	swissjson "swiss/internal/json"
)

// ToJSON parses CSV (first row is treated as the header) into a JSON array of
// objects.
func ToJSON(in string) (string, error) {
	records, err := read(in)
	if err != nil {
		return "", err
	}
	if len(records) == 0 {
		return "[]", nil
	}
	headers := records[0]
	rows := make([]map[string]string, 0, len(records)-1)
	for _, rec := range records[1:] {
		row := make(map[string]string, len(headers))
		for i, h := range headers {
			if i < len(rec) {
				row[h] = rec[i]
			} else {
				row[h] = ""
			}
		}
		rows = append(rows, row)
	}
	b, err := json.Marshal(rows)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ToXML converts CSV into an indented XML document.
func ToXML(in string) (string, error) {
	j, err := ToJSON(in)
	if err != nil {
		return "", err
	}
	return swissjson.ToXML(j)
}

// ToYAML converts CSV into YAML.
func ToYAML(in string) (string, error) {
	j, err := ToJSON(in)
	if err != nil {
		return "", err
	}
	return swissjson.ToYAML(j)
}

// Escape quotes a single CSV field per RFC 4180 when it contains a comma,
// quote or newline.
func Escape(in string) (string, error) {
	var buf bytes.Buffer
	w := encodingcsv.NewWriter(&buf)
	if err := w.Write([]string{in}); err != nil {
		return "", err
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return "", err
	}
	return strings.TrimRight(buf.String(), "\n"), nil
}

// Unescape returns the literal value of the first field of a CSV record.
func Unescape(in string) (string, error) {
	r := encodingcsv.NewReader(strings.NewReader(in))
	rec, err := r.Read()
	if err != nil {
		if err == io.EOF {
			return "", nil
		}
		return "", err
	}
	if len(rec) == 0 {
		return "", nil
	}
	return rec[0], nil
}

func read(in string) ([][]string, error) {
	r := encodingcsv.NewReader(strings.NewReader(in))
	r.TrimLeadingSpace = true
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("invalid csv: %v", err)
	}
	return records, nil
}
