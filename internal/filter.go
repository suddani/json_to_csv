package json_to_csv

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"strings"
)

type Filter interface {
	Check(map[string]interface{}) bool
}

func FilterProcess(input io.Reader, output StringWriter, filter Filter) error {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		data := map[string]interface{}{}

		d := json.NewDecoder(strings.NewReader(scanner.Text()))
		d.UseNumber()
		err := d.Decode(&data)
		if err != nil {
			return err
		}

		if filter != nil && !filter.Check(data) {
			continue
		}

		_, err = output.WriteString(scanner.Text())
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
