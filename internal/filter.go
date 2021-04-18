package json_to_csv

type Filter interface {
	Check(map[string]interface{}) bool
}
