package errutil_test

import (
	"testing"

	"gopkg.hlmpn.dev/pkg/go-logger/errutil"
)

const testErrMsg = "[TEST] This is a sample error message for testing purposes only"

func TestError(t *testing.T) {
	// Test New
	err := errutil.New(testErrMsg)
	if err.Message != testErrMsg {
		t.Errorf("New() = %v, want %v", err.Message, testErrMsg)
	}

	// Test Error() interface
	if err.Error() != testErrMsg {
		t.Errorf("Error() = %v, want %v", err.Error(), testErrMsg)
	}

	// Test Unwrap
	if err.Unwrap().Error() != testErrMsg {
		t.Errorf("Unwrap() = %v, want %v", err.Unwrap().Error(), testErrMsg)
	}

	// Test Log methods (just verify they don't panic)
	err.Log()
	err.LogError()
	err.LogWarning()

	// Test Warning conversion
	warning := err.Warning()
	if warning.Message != testErrMsg {
		t.Errorf("Warning() = %v, want %v", warning.Message, testErrMsg)
	}

	// Test E function
	e := errutil.E("simple error")
	if e.Error() != "simple error" {
		t.Errorf("E() = %v, want %v", e.Error(), "simple error")
	}

	// Test Errorf
	formattedErr := errutil.Errorf("[TEST] This is a sample %s for testing purposes only", "error message")
	if formattedErr.Error() != testErrMsg {
		t.Errorf("Errorf() = %v, want %v", formattedErr.Error(), testErrMsg)
	}

	// Test TrErr
	trErr := errutil.TrErr("test", nil, "tag1", "tag2")
	if trErr.Error() != "@tag1@tag2: test, Error: " {
		t.Errorf("TrErr() = %v, want %v", trErr.Error(), "@tag1@tag2: test, Error: ")
	}

	// Test TrErr with actual error
	actualErr := errutil.E("actual error")
	trErr = errutil.TrErr("test", actualErr, "tag1")
	if trErr.Error() != "@tag1: test, Error: actual error" {
		t.Errorf("TrErr() with error = %v, want %v", trErr.Error(), "@tag1: test, Error: actual error")
	}

	// Test TrErrorf
	trErr = errutil.TrErrorf("testFunc", "test %s", "error")
	if trErr.Error() != "testFunc: test error" {
		t.Errorf("TrErrorf() = %v, want %v", trErr.Error(), "testFunc: test error")
	}
}
