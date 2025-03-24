package logger

import (
	"io"
	"testing"
)

func SetDefaultLoggerOutput(t *testing.T, w io.Writer) {
	defaultLogger.SetOutput(w)
}
