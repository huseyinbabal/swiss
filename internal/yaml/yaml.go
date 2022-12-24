package yaml

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"strings"
)

func Format(in string) (string, error) {
	d := yaml.NewDecoder(strings.NewReader(in))
	node := yaml.Node{
		Kind:        0,
		Style:       0,
		Tag:         "",
		Value:       "",
		Anchor:      "",
		Alias:       &yaml.Node{},
		Content:     []*yaml.Node{},
		HeadComment: "",
		LineComment: "",
		FootComment: "",
		Line:        0,
		Column:      0,
	}
	err := d.Decode(&node)
	var out bytes.Buffer
	for err == nil {
		e := yaml.NewEncoder(&out)
		e.SetIndent(4)
		if err := e.Encode(&node); err != nil {
			log.Fatal(err)
		}
		e.Close()

		err = d.Decode(&node)
		if err == nil {
			fmt.Fprintln(&out, "---")
		}
	}

	if err != nil && err != io.EOF {
		return "", err
	}
	return out.String(), nil
}

func Beautify(in string) (string, error) {
	in = strings.TrimLeft(in, "{")
	in = strings.TrimRight(in, "}")
	count := strings.Count(in, "{")
	curlyBrackets := 0
	for i, ch := range in {
		if ch == '{' {
			curlyBrackets++
		}
		if ch == ',' {
			in = in[:i] + fmt.Sprintf("\n%s", strings.Repeat("\t", curlyBrackets)) + in[i+1:]
		}
	}
	for i := 0; i < count; i++ {
		in = strings.Replace(in, "{", fmt.Sprintf("\n%s", strings.Repeat("\t", i+1)), 1)
	}
	in = strings.ReplaceAll(in, "}", "")
	return in, nil
}

func ToJson(in string) (string, error) {
	var yamlObj interface{}
	err := yaml.Unmarshal([]byte(in), &yamlObj)
	if err != nil {
		return "", fmt.Errorf("error occurred while unmarshalling json %v", err)
	}
	j, err := json.Marshal(yamlObj)
	if err != nil {
		return "", fmt.Errorf("error occurred while marshalling into json %v", err)
	}
	return string(j), nil
}
