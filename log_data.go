package logger

import (
	"reflect"
)

func LogStruct(data any) error {
	v := reflect.ValueOf(data)

	switch {
	case data == nil:
		return ErrLogStructNil
	case v.Kind() != reflect.Struct:
		return ErrLogStructNotStruct
	}

	t := v.Type()
	for i := range v.NumField() {
		field := t.Field(i)
		value := v.Field(i)
		Logf("%s: %v", field.Name, value.Interface())
	}

	return nil
}

func LogSlice(data any) error {
	v := reflect.ValueOf(data)

	switch {
	case data == nil:
		return ErrLogSliceNil
	case v.Kind() != reflect.Slice:
		return ErrLogSliceNotSlice
	}

	for i := range v.Len() {
		Logf("[%d]: %v", i, v.Index(i).Interface())
	}
	return nil
}
