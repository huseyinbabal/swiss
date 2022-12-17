package json

import (
	"bytes"
	"encoding/json"
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
