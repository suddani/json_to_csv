package json_to_csv

import (
	"encoding/csv"
	"errors"
	"io"
)

type limitWriter struct {
	CSV   *csv.Writer
	Limit int
	count int
}

func (w *limitWriter) Write(row []string) error {
	err := w.CSV.Write(row)
	if err != nil {
		return err
	}
	w.count += 1
	if w.count >= w.Limit {
		return errors.New("maximum number of rows reached")
	}
	return nil
}

func (w *limitWriter) Flush() {
	w.CSV.Flush()
}

func NewLimitCsvWriter(writer io.Writer, limit int, header bool) SimpleCsvWriter {
	if limit <= 0 {
		return csv.NewWriter(writer)
	}
	if header {
		limit += 1
	}

	return &limitWriter{
		CSV:   csv.NewWriter(writer),
		Limit: limit,
	}
}
