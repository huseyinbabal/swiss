package xml

import (
	"encoding/json"
	"fmt"
	xj "github.com/basgys/goxml2json"
	"sort"
	"strings"
	swissjson "swiss/internal/json"
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
