package json_to_csv

import (
	"encoding/json"
	"errors"
	"io"
)

type JsonWriter struct {
	Writer  io.Writer
	Encoder *json.Encoder
	Limit   int
	count   int
}

func (w *JsonWriter) Write(row interface{}) error {
	err := w.Encoder.Encode(row)
	if err != nil {
		return err
	}
	if w.Limit <= 0 {
		return nil
	}
	w.count += 1
	if w.count >= w.Limit {
		return errors.New("maximum number of rows reached")
	}
	return nil
}

func (w *JsonWriter) Flush() {
}

func NewJsonWriter(writer io.Writer, limit int) *JsonWriter {
	return &JsonWriter{
		Writer:  writer,
		Encoder: json.NewEncoder(writer),
		Limit:   limit,
	}
}
