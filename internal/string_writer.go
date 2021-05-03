package json_to_csv

type StringWriter interface {
	WriteString(s string) (int, error)
	Flush() error
}
