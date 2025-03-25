package logger

import (
	"gopkg.hlmpn.dev/pkg/xprint"
)

// Warn logs a warning message in orange with a warning triangle emoji (⚠️)
func Warn(msg string) {
	// Define a warning emoji for the log message
	emoji := "⚠️" // Warning emoji

	// Log a nicely formatted warning message with ASCII colors and emoji
	defaultLogger.Printf("\n%s%s%s Warning: %s %s\n",
		bold, orange, emoji, msg, reset)
}

// Warnf logs a formatted warning message in orange with a warning triangle emoji (⚠️)
func Warnf(format string, v ...interface{}) {
	msg := xprint.Printf(format, v...)

	// Define a warning emoji for the log message
	emoji := "⚠️" // Warning emoji

	// Log a nicely formatted warning message with ASCII colors and emoji
	defaultLogger.Printf("\n%s%s%s Warning: %s %s\n",
		bold, orange, emoji, msg, reset)
}

// LogRed logs an error message in red with a bordered format.
// Output format:
// ========================================
// ❌ ERROR: your message here
// ========================================
func LogRed(msg string) {
	// Define the emoji for the error message
	emoji := emojiMap["fail"]

	// Log a nicely formatted error message with ASCII colors and emoji
	defaultLogger.Printf("\n\n%s%s========================================%s\n%s%s%s ERROR: %s%s\n%s%s========================================%s\n\n",
		bold, red, reset,
		bold, red, emoji, msg, reset,
		bold, red, reset)
}

// LogPurple logs a message in purple
func LogPurple(msg string) {
	// Log a nicely formatted file change detected message with ASCII colors and emoji
	defaultLogger.Printf("\n%s%s %s %s\n",
		bold, darkPurple, msg, reset)
}

// LogOrange logs a message in orange
func LogOrange(msg string) {
	// Log a nicely formatted file change detected message with ASCII colors and emoji
	defaultLogger.Printf("\n%s%s %s %s\n",
		bold, orange, msg, reset)
}

// LogRedf logs a formatted error message in red with a bordered format.
// Output format:
// ========================================
// ❌ ERROR: your formatted message here
// ========================================
func LogRedf(format string, args ...interface{}) {
	// Define the emoji for the error message
	emoji := emojiMap["fail"]

	// Log a nicely formatted error message with ASCII colors and emoji
	defaultLogger.Printf("\n\n%s%s========================================%s\n%s%s%s ERROR: %s%s\n%s%s========================================%s\n\n",
		bold, red, reset,
		bold, red, emoji, xprint.Printf(format, args...), reset,
		bold, red, reset)
}

// LogPurplef logs a formatted message in purple
func LogPurplef(format string, args ...interface{}) {
	// Log a nicely formatted file change detected message with ASCII colors and emoji
	defaultLogger.Printf("\n%s%s %s %s\n",
		bold, darkPurple, xprint.Printf(format, args...), reset)
}

// LogOrangef logs a formatted message in orange
func LogOrangef(format string, args ...interface{}) {
	// Log a nicely formatted file change detected message with ASCII colors and emoji
	defaultLogger.Printf("\n%s%s %s %s\n",
		bold, orange, xprint.Printf(format, args...), reset)
}

// NoteF logs a formatted note message in orange
func NoteF(format string, v ...interface{}) {
	msg := xprint.Printf(format, v...)
	defaultLogger.Printf("\n%s%s %s %s\n",
		bold, orange, msg, reset)
}

// Note logs a note message in orange
func Note(msg string) {
	defaultLogger.Printf("\n%s%s %s %s\n",
		bold, orange, msg, reset)
}

// LogInfof logs a formatted info message with an info emoji (ℹ️)
func LogInfof(format string, args ...interface{}) {
	msg := xprint.Printf(format, args...)

	// Define an info emoji for the log message
	emoji := emojiMap["info"]

	// Log a nicely formatted info message with ASCII colors and emoji
	defaultLogger.Printf("\n%s%s%s Info: %s %s\n",
		bold, blue, emoji, msg, reset)
}

// LogInfo logs an info message with an info emoji (ℹ️)
func LogInfo(msg string) {
	// Define an info emoji for the log message
	emoji := emojiMap["info"]

	// Log a nicely formatted info message with ASCII colors and emoji
	defaultLogger.Printf("\n%s%s%s Info: %s %s\n",
		bold, blue, emoji, msg, reset)
}

// LogGreenf logs a formatted message in green
func LogGreenf(format string, args ...interface{}) {
	// Log a nicely formatted message with ASCII colors
	defaultLogger.Printf("\n%s%s %s %s\n",
		bold, green, xprint.Printf(format, args...), reset)
}

// LogGreen logs a message in green
func LogGreen(msg string) {
	// Log a nicely formatted message with ASCII colors
	defaultLogger.Printf("\n%s%s %s %s\n",
		bold, green, msg, reset)
}

// LogSuccess logs a success message in green with a bordered format and checkmark emoji
func LogSuccess(msg string) {
	// Define the emoji for the success message
	emoji := emojiMap["success"]

	// Log a nicely formatted success message with ASCII colors and emoji
	defaultLogger.Printf("\n\n%s%s========================================%s\n%s%s%s SUCCESS: %s%s\n%s%s========================================%s\n\n",
		bold, green, reset,
		bold, green, emoji, msg, reset,
		bold, green, reset)
}

// LogSuccessf logs a formatted success message in green with a bordered format and checkmark emoji
func LogSuccessf(format string, args ...interface{}) {
	// Define the emoji for the success message
	emoji := emojiMap["success"]

	// Log a nicely formatted success message with ASCII colors and emoji
	defaultLogger.Printf("\n\n%s%s========================================%s\n%s%s%s SUCCESS: %s%s\n%s%s========================================%s\n\n",
		bold, green, reset,
		bold, green, emoji, xprint.Printf(format, args...), reset,
		bold, green, reset)
}
