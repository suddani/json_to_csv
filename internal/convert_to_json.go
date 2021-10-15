package json_to_csv

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"strings"
)

func ConvertToJson(input io.Reader, output SimpleJsonWriter, keys []string, filter Filter) error {
	keyMap := map[string]bool{}
	if keys != nil {
		for _, key := range keys {
			keyMap[key] = true
		}
	}
	reader := csv.NewReader(input)
	header := []string{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if len(header) == 0 {
			header = append(header, record...)
			continue
		}

		data := map[string]interface{}{}
		for index, name := range header {
			if keys == nil || keyMap[name] {

				var entry interface{}
				d := json.NewDecoder(strings.NewReader(record[index]))
				d.UseNumber()
				err := d.Decode(&entry)
				if err != nil {
					data[name] = record[index]
				} else {
					data[name] = entry
				}
			}
		}

		if filter != nil && !filter.Check(data) {
			continue
		}
		err = output.Write(data)
		if err != nil {
			return err
		}
	}
	return nil
}
