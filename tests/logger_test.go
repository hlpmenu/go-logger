package logger_test

import (
	"bytes"
	"testing"

	"gopkg.hlmpn.dev/pkg/go-logger"
)

func TestExportedLogFunctions(t *testing.T) {
	// Capture output to prevent test noise
	var buf bytes.Buffer
	logger.SetDefaultLoggerOutput(t, &buf)

	// Simple smoke tests - just ensure they don't panic
	tests := []struct {
		name string
		fn   func()
	}{
		{"LogRed", func() { logger.LogRed("test") }},
		{"LogRedf", func() { logger.LogRedf("test %s", "msg") }},
		{"LogPurple", func() { logger.LogPurple("test") }},
		{"LogPurplef", func() { logger.LogPurplef("test %s", "msg") }},
		{"LogOrange", func() { logger.LogOrange("test") }},
		{"LogOrangef", func() { logger.LogOrangef("test %s", "msg") }},
		{"LogSuccess", func() { logger.LogSuccess("test") }},
		{"LogSuccessf", func() { logger.LogSuccessf("test %s", "msg") }},
		{"Note", func() { logger.Note("test") }},
		{"NoteF", func() { logger.NoteF("test %s", "msg") }},
		{"Warn", func() { logger.Warn("test") }},
		{"Warnf", func() { logger.Warnf("test %s", "msg") }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			buf.Reset()
			tt.fn() // Should not panic
		})
	}
}

func TestPrintln(t *testing.T) {
	logger.Println("test")

}
