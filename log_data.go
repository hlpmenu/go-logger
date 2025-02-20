package logger

import (
	"errors"
	"reflect"
)

func LogStruct(data interface{}) error {
	if data == nil {
		return errors.New("struct is nil")
	}

	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		return errors.New("not a struct")
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		Logf("%s: %v", field.Name, value.Interface())
	}
	return nil
}

func LogSlice(data interface{}) error {
	if data == nil {
		return errors.New("slice is nil")
	}

	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		return errors.New("not a slice")
	}

	for i := 0; i < v.Len(); i++ {
		Logf("[%d]: %v", i, v.Index(i).Interface())
	}
	return nil
}
