package json_to_csv

import "strings"

func SliceString(s string) []string {
	if s == "" {
		return nil
	}

	array := strings.Split(s, ",")
	if len(array) == 0 {
		return nil
	}

	return array
}
