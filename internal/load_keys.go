package json_to_csv

import (
	"encoding/csv"
	"os"
)

func LoadKeys(keys []string, fileName string) ([]string, error) {
	if keys != nil {
		return keys, nil
	}
	if fileName == "" {
		return nil, nil
	}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		keys = append(keys, row[0])
	}

	return keys, nil
}
