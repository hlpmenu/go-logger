package logger // nolint:testpackage //

import (
	"bytes"
	"testing"
)

func TestExportedLogFunctions(t *testing.T) {
	// Capture output to prevent test noise
	var buf bytes.Buffer
	defaultLogger.SetOutput(&buf)

	// Simple smoke tests - just ensure they don't panic
	tests := []struct {
		name string
		fn   func()
	}{
		{"LogRed", func() { LogRed("test") }},
		{"LogRedf", func() { LogRedf("test %s", "msg") }},
		{"LogPurple", func() { LogPurple("test") }},
		{"LogPurplef", func() { LogPurplef("test %s", "msg") }},
		{"LogOrange", func() { LogOrange("test") }},
		{"LogOrangef", func() { LogOrangef("test %s", "msg") }},
		{"LogSuccess", func() { LogSuccess("test") }},
		{"LogSuccessf", func() { LogSuccessf("test %s", "msg") }},
		{"Note", func() { Note("test") }},
		{"NoteF", func() { NoteF("test %s", "msg") }},
		{"Warn", func() { Warn("test") }},
		{"Warnf", func() { Warnf("test %s", "msg") }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			buf.Reset()
			tt.fn() // Should not panic
		})
	}
}
