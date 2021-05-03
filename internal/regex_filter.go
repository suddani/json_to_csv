package json_to_csv

import (
	"fmt"
	"regexp"
)

type regexFilter struct {
	key    string
	values []*regexp.Regexp
}

func (filter *regexFilter) Check(data map[string]interface{}) bool {
	value := data[filter.key]
	fmt.Printf("%v => %v\n", filter.key, value)
	for _, regex := range filter.values {
		if regex.MatchString(fmt.Sprint(value)) {
			return true
		}
	}
	return false
}

func NewRegexFilter(f []string, fileName string) Filter {
	keys, err := LoadKeys(f, fileName)
	if err != nil || keys == nil {
		return nil
	}

	values := []*regexp.Regexp{}

	for index, key := range keys {
		if index == 0 {
			continue
		}
		value := regexp.MustCompile(key)
		values = append(values, value)
	}

	return &regexFilter{
		key:    keys[0],
		values: values,
	}
}
