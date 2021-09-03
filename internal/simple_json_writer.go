package json_to_csv

type SimpleJsonWriter interface {
	Write(interface{}) error
	Flush()
}
