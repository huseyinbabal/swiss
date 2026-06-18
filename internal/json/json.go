package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"swiss/internal/convert"
	yaml2 "swiss/internal/yaml"

	"gopkg.in/yaml.v3"
)

func Beautify(uglyJson string) (string, error) {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(uglyJson), "", "    ")
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func Uglify(in string) (string, error) {
	var out bytes.Buffer
	err := json.Compact(&out, []byte(in))
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func ToYAML(in string) (string, error) {
	var j interface{}

	err := yaml.Unmarshal([]byte(in), &j)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshall json to object. Err: %v", err)
	}

	marshal, err := yaml.Marshal(j)
	if err != nil {
		return "", fmt.Errorf("failed to marshall json object to yaml. Err: %v", err)
	}
	out, err := yaml2.Beautify(string(marshal))
	if err != nil {
		return "", fmt.Errorf("failed to beautify yaml. Err: %v", err)
	}
	return out, nil
}

// ToXML converts a JSON document into an indented XML document.
func ToXML(in string) (string, error) {
	obj, err := decode(in)
	if err != nil {
		return "", err
	}
	return convert.ObjectToXML(obj), nil
}

// ToCSV converts a JSON array of objects (or a single object) into CSV.
func ToCSV(in string) (string, error) {
	obj, err := decode(in)
	if err != nil {
		return "", err
	}
	return convert.ObjectToDelimited(obj, ',')
}

// ToTSV converts a JSON array of objects (or a single object) into TSV.
func ToTSV(in string) (string, error) {
	obj, err := decode(in)
	if err != nil {
		return "", err
	}
	return convert.ObjectToDelimited(obj, '\t')
}

// ToGoStruct generates a Go struct definition matching the shape of the JSON.
func ToGoStruct(in string) (string, error) {
	obj, err := decode(in)
	if err != nil {
		return "", err
	}
	return convert.ObjectToGoStruct(obj), nil
}

// Escape returns the JSON-escaped form of a raw string, without the surrounding
// quotes, so it can be embedded into a JSON document.
func Escape(in string) (string, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return "", err
	}
	return string(b[1 : len(b)-1]), nil
}

// Unescape reverses Escape, accepting the value with or without surrounding
// double quotes.
func Unescape(in string) (string, error) {
	quoted := in
	if !strings.HasPrefix(quoted, "\"") || !strings.HasSuffix(quoted, "\"") {
		quoted = "\"" + quoted + "\""
	}
	var out string
	if err := json.Unmarshal([]byte(quoted), &out); err != nil {
		return "", fmt.Errorf("failed to unescape json string. Err: %v", err)
	}
	return out, nil
}

func decode(in string) (interface{}, error) {
	var obj interface{}
	if err := json.Unmarshal([]byte(in), &obj); err != nil {
		return nil, fmt.Errorf("failed to unmarshall json. Err: %v", err)
	}
	return obj, nil
}
