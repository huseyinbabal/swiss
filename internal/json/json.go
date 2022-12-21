package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	yaml2 "swiss/internal/yaml"
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
