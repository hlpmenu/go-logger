package logger

import (
	"io"
	"reflect"

	"gopkg.hlmpn.dev/pkg/go-logger/internal/jsonformat"
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

func LogPrettyJSON(v any) error {
	r := jsonformat.PrettyPrintJSONIface(v)

	PrintRdr(r)
	return nil
}

func PrintPrettyJSONBytes(b []byte) error {
	r := jsonformat.PrettyPrintJSON(b)

	PrintRdr(r)
	return nil
}

func PrintPrettyJSONReader(r io.Reader) error {
	b := jsonformat.PrettyPrintJSONFromReader(r)

	PrintRdr(b)
	return nil
}
