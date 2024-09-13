package logger

import (
	"fmt"
)

func LogErrorf(format string, args ...interface{}) {
	// Define ANSI escape codes for red text and bold font
	red := "\033[31m"
	bold := "\033[1m"
	reset := "\033[0m"

	// Define the emoji for the error message
	emoji := emojiMap["fail"]

	// Log a nicely formatted error message with ASCII colors and emoji
	defaultLogger.Fatalf("\n\n%s%s========================================%s\n%s%s%s ERROR: %s%s\n%s%s========================================%s\n\n",
		bold, red, reset,
		bold, red, emoji, fmt.Sprintf(format, args...), reset,
		bold, red, reset)
}

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
