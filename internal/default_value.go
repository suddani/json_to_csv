package json_to_csv

import (
	"reflect"
	"strings"
)

type DefaultValue interface {
	DefaultInterface(string, interface{}) interface{}
	DefaultString(string, string) string
}

type emptyDefaults struct{}

func (*emptyDefaults) DefaultInterface(key string, value interface{}) interface{} {
	return value
}
func (*emptyDefaults) DefaultString(key string, value string) string {
	return value
}

type mapDefaults struct {
	values map[string]string
}

func (d *mapDefaults) DefaultInterface(key string, value interface{}) interface{} {
	val, ok := d.values[key]
	if ok && (value == nil || (reflect.TypeOf(value).String() == "string" && reflect.ValueOf(value).String() == "")) {
		return val
	}

	return value
}
func (d *mapDefaults) DefaultString(key string, value string) string {
	val, ok := d.values[key]
	if ok && (value == "" || value == "<nil>") {
		return val
	}
	return value
}

func NewDefaultValue(f []string) DefaultValue {
	if len(f) == 0 {
		return &emptyDefaults{}
	}
	values := map[string]string{}
	for _, valueString := range f {
		data := strings.SplitN(valueString, ":", 2)
		values[data[0]] = data[1]
	}

	return &mapDefaults{
		values: values,
	}
}
