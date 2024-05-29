package json_to_csv

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
)

func digMap(data map[string]interface{}, key string) (string, error) {
	var output interface{}
	value := data
	keys := strings.Split(key, ".")
	for _, part := range keys {
		if value == nil {
			return "", fmt.Errorf("Could not find key")
		}
		output = value[part]
		switch output.(type) {
		case map[string]interface{}:
			value = output.(map[string]interface{})
		default:
			value = nil
		}
	}
	return fmt.Sprint(output), nil
}
func ConvertToCsv(input io.Reader, output SimpleCsvWriter, keys []string, filter Filter, header bool, bufferLimit BufferLimit) error {
	if header && keys != nil {
		err := output.Write(keys)
		if err != nil {
			return err
		}
	}
	scanner := bufio.NewScanner(input)
	if bufferLimit.Valid {
		initialScanBuffer := make([]byte, 0, bufferLimit.Default)
		scanner.Buffer(initialScanBuffer, bufferLimit.Max)
	}
	for scanner.Scan() {
		data := map[string]interface{}{}

		d := json.NewDecoder(strings.NewReader(scanner.Text()))
		d.UseNumber()
		err := d.Decode(&data)
		if err != nil {
			return err
		}

		if keys == nil {
			for key, _ := range data {
				keys = append(keys, key)
			}
			sort.Strings(keys)
			if header {
				err := output.Write(keys)
				if err != nil {
					return err
				}
			}
		}

		if filter != nil && !filter.Check(data) {
			continue
		}

		values := []string{}
		for _, key := range keys {
			value, err := digMap(data, key)
			if err != nil {
				values = append(values, "")
			} else {
				values = append(values, value)
			}
		}

		err = output.Write(values)
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
