package errutil

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.hlmpn.dev/pkg/go-logger"
)

// CustomError is a custom error type that satisfies the error interface
type Error struct {
	Message string
}

func (e *Error) Unwrap() error {
	return errors.New(e.Message)
}

// Error implements the error interface for CustomError
func (e *Error) Error() string {
	return e.Message
}

// Logs the error message to the console
func (e *Error) Log() {
	log.Print(e.Message)
}

// Logs the error message to the console and exits the program
func (e *Error) LogFatal() {
	log.Fatal(e.Message)
}
func (e *Error) LogFatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
func (e *Error) Panic() {
	panic(e.Message)
}

// Logs the error message to the console and exits the program with a custom exit code
func (e *Error) LogExit(code int) {
	log.Fatalf("%s, Exiting with code %d", e.Message, code)
}

func (e *Error) LogError() {
	logger.LogRed(e.Message)
}
func (e *Error) LogErrorFatal() {
	logger.LogRed(e.Message)
	e.LogFatalf("")
}
func (e *Error) LogErrorFatalf(format string, args ...interface{}) {
	constmsg := fmt.Sprintf(format+fmt.Sprintf("Error: %s", e.Message), args...)
	logger.LogRed(constmsg)
	os.Exit(1)
}

func (e *Error) Warning() *Warning {
	return &Warning{
		Message: e.Message,
	}
}
func (e *Error) LogWarning() {
	w := Warning{
		Message: e.Message,
	}
	w.Warn()
}

// New creates a new custom error with the provided message
func New(msg string) *Error {
	return &Error{
		Message: msg,
	}
}
func E(msg string) error {
	e := errors.New(msg)
	return e
}

// Errorf creates a new custom error with a formatted message
func Errorf(format string, args ...interface{}) *Error {
	return &Error{
		Message: fmt.Sprintf(format, args...),
	}
}

// TrError creates a formatted error message from the provided strings and error
func TrErr(msg string, err error, tags ...string) *Error {
	// Prefix each tag with "@"
	for i, tag := range tags {
		tags[i] = "@" + tag
	}
	var errMsg string
	if err == nil {
		errMsg = ""
	} else {
		errMsg = err.Error()
	}

	tagsFormatted := strings.Join(tags, "")
	return &Error{
		Message: fmt.Sprintf("%s: %s, Error: %s", tagsFormatted, msg, errMsg),
	}
}

func TrErrorf(functionname string, format string, args ...interface{}) *Error {
	return &Error{
		Message: fmt.Sprintf("%s: %s", functionname, fmt.Sprintf(format, args...)),
	}
}
