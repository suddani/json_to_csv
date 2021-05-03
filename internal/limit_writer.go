package json_to_csv

import (
	"bufio"
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

type stringWriter struct {
	Writer *bufio.Writer
	Limit  int
	Count  int
}

func (w *stringWriter) Flush() error {
	return w.Writer.Flush()
}

func (w *stringWriter) WriteString(s string) (int, error) {
	written, err := w.Writer.WriteString(s)
	if err != nil {
		return written, err
	}

	w.Count += 1
	if w.Count >= w.Limit {
		return written, errors.New("maximum number of strings reached")
	}
	return written, nil
}

func NewLimitWriter(writer io.Writer, limit int) StringWriter {
	if limit <= 0 {
		return bufio.NewWriter(writer)
	}

	return &stringWriter{
		Writer: bufio.NewWriter(writer),
		Limit:  limit,
	}
}
