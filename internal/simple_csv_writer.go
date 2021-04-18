package json_to_csv

type SimpleCsvWriter interface {
	Write([]string) error
	Flush()
}
