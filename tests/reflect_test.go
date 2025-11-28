package logger_test

import (
	"errors"
	"io"
	"reflect"
	"testing"

	"gopkg.hlmpn.dev/pkg/go-logger"
)

type dummy struct {
	Name string
	Age  int
	City string
	Car  string
}

func Test(t *testing.T) {
	logger.Logf("Hello, world!")
	d := dummy{
		Name: "John",
		Age:  25,
		City: "New York",
		Car:  "BMW",
	}
	err := Reflect(d)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func Reflect(data any) error {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		return errors.New("not a struct")
	}

	for i := range v.NumField() {
		field := v.Type().Field(i)
		value := v.Field(i)
		logger.Logf("%s: %v", field.Name, value.Interface())
	}
	return nil

}

func TestShowcaseRange(t *testing.T) {

	var runs = 10
	for i := range runs {
		logger.Logf("hello world number %d", i)
	}

}

func TestIfRefelctValueOfPanicsOnNil(t *testing.T) {
	var d uintptr
	doesItPanic(d)
}

func doesItPanic(data any) {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		logger.Logf("reflect.kind = %v", v.Kind())
		logger.Log("not a struct")
		return
	}
	logger.Logf("reflect.kind = %v", v.Kind())
	logger.Logf("is a struct")

}

func TestSliceRangeLen(t *testing.T) {
	s := []string{"apple", "banana", "cherry", "date"}

	v := reflect.ValueOf(s)
	logger.Logf("Slice length: %d", v.Len())

	for i := range v.Len() {
		logger.Logf("[%d]: %v", i, v.Index(i).Interface())
	}
}

type emptyStruct struct{}

type testStruct struct {
	Name    string
	Age     int
	Active  bool
	Pointer *string
}

func TestLogStruct(t *testing.T) {
	tests := []struct {
		name        string
		input       any
		expectError bool
		errorIs     error
	}{
		{
			name:        "nil value",
			input:       nil,
			expectError: true,
			errorIs:     logger.ErrLogStructNil,
		},
		{
			name:        "non-struct type",
			input:       "not a struct",
			expectError: true,
			errorIs:     logger.ErrLogStructNotStruct,
		},
		{
			name:        "non-struct slice",
			input:       []int{1, 2, 3},
			expectError: true,
			errorIs:     logger.ErrLogStructNotStruct,
		},
		{
			name:        "empty struct",
			input:       emptyStruct{},
			expectError: false,
		},
		{
			name: "valid struct",
			input: testStruct{
				Name:   "John",
				Age:    30,
				Active: true,
			},
			expectError: false,
		},
		{
			name: "struct with nil pointer",
			input: testStruct{
				Name:    "Jane",
				Age:     25,
				Active:  false,
				Pointer: nil,
			},
			expectError: false,
		},
	}

	// Temporarily disable logging output for tests
	logger.SetDefaultLoggerOutput(t, io.Discard)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := logger.LogStruct(tt.input)

			if tt.expectError && err == nil {
				t.Errorf("expected error but got nil")
			}

			if !tt.expectError && err != nil {
				t.Errorf("expected no error but got: %v", err)
			}

			if tt.errorIs != nil && !errors.Is(err, tt.errorIs) {
				t.Errorf("expected error to be %v but got: %v", tt.errorIs, err)
			}
		})
	}
}

func TestLogSlice(t *testing.T) {
	// Create a slice with a nil element
	strSlice := make([]*string, 3)
	hello := "hello"
	world := "world"
	strSlice[0] = &hello
	strSlice[1] = nil
	strSlice[2] = &world

	tests := []struct {
		name        string
		input       any
		expectError bool
		errorIs     error
	}{
		{
			name:        "nil value",
			input:       nil,
			expectError: true,
			errorIs:     logger.ErrLogSliceNil,
		},
		{
			name:        "non-slice type",
			input:       "not a slice",
			expectError: true,
			errorIs:     logger.ErrLogSliceNotSlice,
		},
		{
			name:        "non-slice struct",
			input:       testStruct{},
			expectError: true,
			errorIs:     logger.ErrLogSliceNotSlice,
		},
		{
			name:        "empty slice",
			input:       []int{},
			expectError: false,
		},
		{
			name:        "int slice",
			input:       []int{1, 2, 3, 4, 5},
			expectError: false,
		},
		{
			name:        "slice with nil elements",
			input:       strSlice,
			expectError: false,
		},
	}

	// Temporarily disable logging output for tests
	logger.SetDefaultLoggerOutput(t, io.Discard)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := logger.LogSlice(tt.input)

			if tt.expectError && err == nil {
				t.Errorf("expected error but got nil")
			}

			if !tt.expectError && err != nil {
				t.Errorf("expected no error but got: %v", err)
			}

			if tt.errorIs != nil && !errors.Is(err, tt.errorIs) {
				t.Errorf("expected error to be %v but got: %v", tt.errorIs, err)
			}
		})
	}
}
