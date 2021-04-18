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

func ConvertToCsv(input io.Reader, output SimpleCsvWriter, keys []string, filter Filter, header bool) error {
	if header && keys != nil {
		err := output.Write(keys)
		if err != nil {
			return err
		}
	}
	scanner := bufio.NewScanner(input)
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
			values = append(values, fmt.Sprint(data[key]))
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
