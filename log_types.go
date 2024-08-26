package logger

import (
	"log"
	"os"
)

// Global logger instance
var defaultLogger = NewLogger()

// LoggerInterface defines the logging methods
type LoggerInterface interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
}

// Logger struct that implements the LoggerInterface
type Logger struct {
	logger *log.Logger
}

// NewLogger creates a new Logger
func NewLogger() *Logger {
	return &Logger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// Print logs a message
func (l *Logger) Print(v ...interface{}) {
	l.logger.Print(v...)
}

// Printf logs a formatted message
func (l *Logger) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

// Println logs a message with a newline
func (l *Logger) Println(v ...interface{}) {
	l.logger.Println(v...)
}

// Fatal logs a message and then calls os.Exit(1)
func (l *Logger) Fatal(v ...interface{}) {
	l.logger.Fatal(v...)
}

// Fatalf logs a formatted message and then calls os.Exit(1)
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.logger.Fatalf(format, v...)
}

// Fatalln logs a message with a newline and then calls os.Exit(1)
func (l *Logger) Fatalln(v ...interface{}) {
	l.logger.Fatalln(v...)
}

// Panic logs a message and then panics
func (l *Logger) Panic(v ...interface{}) {
	l.logger.Panic(v...)
}

// Panicf logs a formatted message and then panics
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.logger.Panicf(format, v...)
}

// Panicln logs a message with a newline and then panics
func (l *Logger) Panicln(v ...interface{}) {
	l.logger.Panicln(v...)
}

// Print logs a message using the global logger
func Print(v ...interface{}) {
	defaultLogger.Print(v...)
}

// Printf logs a formatted message using the global logger
func Printf(format string, v ...interface{}) {
	defaultLogger.Printf(format, v...)
}

// Println logs a message with a newline using the global logger
func Println(v ...interface{}) {
	defaultLogger.Println(v...)
}

// Fatal logs a message and then calls os.Exit(1) using the global logger
func Fatal(v ...interface{}) {
	defaultLogger.Fatal(v...)
}

// Fatalf logs a formatted message and then calls os.Exit(1) using the global logger
func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatalf(format, v...)
}

// Fatalln logs a message with a newline and then calls os.Exit(1) using the global logger
func Fatalln(v ...interface{}) {
	defaultLogger.Fatalln(v...)
}

// Panic logs a message and then panics using the global logger
func Panic(v ...interface{}) {
	defaultLogger.Panic(v...)
}

// Panicf logs a formatted message and then panics using the global logger
func Panicf(format string, v ...interface{}) {
	defaultLogger.Panicf(format, v...)
}

// Panicln logs a message with a newline and then panics using the global logger
func Panicln(v ...interface{}) {
	defaultLogger.Panicln(v...)
}
