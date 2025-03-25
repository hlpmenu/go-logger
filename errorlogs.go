package logger

import "gopkg.hlmpn.dev/pkg/xprint"

// LogErrorf logs a formatted error message in red with a bordered format and exits the program.
// Output format:
// ========================================
// ❌ ERROR: your formatted message here
// ========================================
// After logging, the program will exit with a non-zero status code.
func LogErrorf(format string, args ...any) {
	// Define ANSI escape codes for red text and bold font
	red := "\033[31m"
	bold := "\033[1m"
	reset := "\033[0m"

	// Define the emoji for the error message
	emoji := emojiMap["fail"]

	// Log a nicely formatted error message with ASCII colors and emoji
	defaultLogger.Fatalf("\n\n%s%s========================================%s\n%s%s%s ERROR: %s%s\n%s%s========================================%s\n\n",
		bold, red, reset,
		bold, red, emoji, xprint.Printf(format, args...), reset,
		bold, red, reset)
}

// LogError logs an error message in red with a bordered format and exits the program.
// Output format:
// ========================================
// ❌ ERROR: your message here
// ========================================
// After logging, the program will exit with a non-zero status code.
func LogError(msg string) {
	// Define ANSI escape codes for red text and bold font
	red := "\033[31m"
	bold := "\033[1m"
	reset := "\033[0m"

	// Define the emoji for the error message
	emoji := emojiMap["fail"]

	// Log a nicely formatted error message with ASCII colors and emoji
	defaultLogger.Fatalf("\n\n%s%s========================================%s\n%s%s%s ERROR: %s%s\n%s%s========================================%s\n\n",
		bold, red, reset,
		bold, red, emoji, msg, reset,
		bold, red, reset)
}
