package logger

import "errors"

// Logging errors
var ()

// Reflect, type assertion, and datatype errors
var (
	ErrLogStructNotStruct = errors.New("logger: provided data argument is not a struct")
	ErrLogStructNil       = errors.New("logger: provided data argument is nil")
	ErrLogSliceNotSlice   = errors.New("logger: provided data argument is not a slice")
	ErrLogSliceNil        = errors.New("logger: provided data argument is nil")
)
