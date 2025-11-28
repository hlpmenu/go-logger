package logger_test

import (
	"io"
	"testing"

	"gopkg.hlmpn.dev/pkg/go-logger"
)

func SetDefaultLoggerOutput(t *testing.T, w io.Writer) {
	logger.SetDefaultLoggerOutput(t, w)
}
