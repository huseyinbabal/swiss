package toml

import (
	"bytes"
	"encoding/json"

	btoml "github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

// ToJSON decodes a TOML document and re-encodes it as indented JSON.
func ToJSON(in string) (string, error) {
	var m map[string]interface{}
	if _, err := btoml.Decode(in, &m); err != nil {
		return "", err
	}
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ToYAML decodes a TOML document and re-encodes it as YAML.
func ToYAML(in string) (string, error) {
	var m map[string]interface{}
	if _, err := btoml.Decode(in, &m); err != nil {
		return "", err
	}
	b, err := yaml.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// FromJSON parses a JSON document and re-encodes it as TOML.
func FromJSON(in string) (string, error) {
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(in), &m); err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := btoml.NewEncoder(&buf).Encode(m); err != nil {
		return "", err
	}
	return buf.String(), nil
}
