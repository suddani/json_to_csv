package json_to_csv

import (
	"bufio"
	"encoding/json"
	"io"
	"sort"
)

func FindAllKeys(input io.Reader, output SimpleCsvWriter, bufferLimit BufferLimit) error {
	keyMap := map[string]bool{}
	scanner := bufio.NewScanner(input)
	if bufferLimit.Valid {
		initialScanBuffer := make([]byte, 0, bufferLimit.Default)
		scanner.Buffer(initialScanBuffer, bufferLimit.Max)
	}
	for scanner.Scan() {
		data := map[string]interface{}{}
		err := json.Unmarshal(scanner.Bytes(), &data)
		if err != nil {
			return err
		}
		for key, _ := range data {
			keyMap[key] = true
		}
	}
	keys := []string{}
	for key, _ := range keyMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		err := output.Write([]string{key})
		if err != nil {
			return err
		}
	}
	return nil
}
