package json_to_csv

import (
	"encoding/json"
	"fmt"
	"strings"
)

type arrayFilter struct {
	key    string
	values map[string]bool
}

func (filter *arrayFilter) Check(data map[string]interface{}) bool {
	value := data[filter.key]
	v, ok := filter.values[fmt.Sprint(value)]
	return v && ok
}

func NewArrayFilter(f []string, fileName string) Filter {
	keys, err := LoadKeys(f, fileName)
	if err != nil || keys == nil {
		return nil
	}

	values := map[string]bool{}
	for index, key := range keys {
		if index == 0 {
			continue
		}
		var value interface{}
		d := json.NewDecoder(strings.NewReader(key))
		d.UseNumber()
		err := d.Decode(&value)
		if err != nil {
			values[key] = true
		} else {
			values[fmt.Sprint(value)] = true
		}
	}

	return &arrayFilter{
		key:    keys[0],
		values: values,
	}
}
