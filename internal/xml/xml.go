package xml

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"sort"
	"strings"
	swissjson "swiss/internal/json"

	xj "github.com/basgys/goxml2json"
)

func ToJSON(in string) (string, error) {
	xml := strings.NewReader(in)
	j, err := xj.Convert(xml)
	if err != nil {
		return "", fmt.Errorf("error occurred while marshalling xmlObj into json %v", err)
	}
	return j.String(), nil
}

func ToYAML(in string) (string, error) {
	j, err := ToJSON(in)
	if err != nil {
		return "", err
	}
	return swissjson.ToYAML(j)
}

func ToCSV(in string) (string, error) {
	j, err := ToJSON(in)
	if err != nil {
		return "", err
	}
	var csv map[string]map[string][]map[string]interface{}
	err = json.Unmarshal([]byte(j), &csv)
	if err != nil {
		return "", err
	}
	var columns []string
	var values []string
	for _, v := range csv {
		for _, items := range v {
			for i, item := range items {
				if i == 0 {
					for k, _ := range item {
						columns = append(columns, k)
					}
					sort.Slice(columns, func(i, j int) bool {
						return columns[i] < columns[j]
					})
				}
				var vals []string
				for _, col := range columns {
					vals = append(vals, fmt.Sprintf("%v", item[col]))
				}
				values = append(values, strings.Join(vals, ","))
			}
		}
	}

	return fmt.Sprintf("%s\n%s", strings.Join(columns, ","), strings.Join(values, "\n")), err
}

var xmlEscaper = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	"\"", "&quot;",
	"'", "&apos;",
)

var xmlUnescaper = strings.NewReplacer(
	"&amp;", "&",
	"&lt;", "<",
	"&gt;", ">",
	"&quot;", "\"",
	"&apos;", "'",
)

// Escape replaces XML-significant characters with their entity references.
func Escape(in string) (string, error) {
	return xmlEscaper.Replace(in), nil
}

// Unescape replaces XML entity references with their literal characters.
func Unescape(in string) (string, error) {
	return xmlUnescaper.Replace(in), nil
}

// Beautify re-renders an XML document with four-space indentation.
func Beautify(in string) (string, error) {
	return reformat(in, "    ")
}

// Uglify removes insignificant whitespace between XML elements.
func Uglify(in string) (string, error) {
	return reformat(in, "")
}

func reformat(in, indent string) (string, error) {
	dec := xml.NewDecoder(strings.NewReader(in))
	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	if indent != "" {
		enc.Indent("", indent)
	}
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("invalid xml: %v", err)
		}
		// Drop whitespace-only text nodes so indentation is driven solely by
		// the encoder instead of the original (arbitrary) formatting.
		if cd, ok := tok.(xml.CharData); ok && strings.TrimSpace(string(cd)) == "" {
			continue
		}
		if err := enc.EncodeToken(tok); err != nil {
			return "", err
		}
	}
	if err := enc.Flush(); err != nil {
		return "", err
	}
	return buf.String(), nil
}
