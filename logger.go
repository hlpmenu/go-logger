package logger

import "fmt"

func Warning(filename string) {
	// Define ANSI escape codes for orange text and bold font
	orange := "\033[38;5;214m" // ANSI code for orange
	bold := "\033[1m"
	reset := "\033[0m"

	// Define a warning emoji for the log message
	emoji := "⚠️" // Warning emoji

	// Log a nicely formatted warning message with ASCII colors and emoji
	defaultLogger.Printf("\n%s%s%s Warning: %s %s\n",
		bold, orange, emoji, filename, reset)
}
func Warnf(format string, v ...interface{}) {
	// Define ANSI escape codes for orange text and bold font
	orange := "\033[38;5;214m" // ANSI code for orange
	bold := "\033[1m"
	reset := "\033[0m"

	msg := fmt.Sprintf(format, v...)

	// Define a warning emoji for the log message
	emoji := "⚠️" // Warning emoji

	// Log a nicely formatted warning message with ASCII colors and emoji
	defaultLogger.Printf("\n%s%s%s Warning: %s %s\n",
		bold, orange, emoji, msg, reset)
}

func LogRed(msg string) {
	// Define ANSI escape codes for red text and bold font
	red := "\033[31m"
	bold := "\033[1m"
	reset := "\033[0m"

	// Define the emoji for the error message
	emoji := emojiMap["fail"]

	// Log a nicely formatted error message with ASCII colors and emoji
	defaultLogger.Printf("\n\n%s%s========================================%s\n%s%s%s ERROR: %s%s\n%s%s========================================%s\n\n",
		bold, red, reset,
		bold, red, emoji, msg, reset,
		bold, red, reset)
}

func LogPurple(msg string) {
	// Define ANSI escape codes for dark purple text and bold font
	darkPurple := "\033[35m"
	bold := "\033[1m"
	reset := "\033[0m"

	// Log a nicely formatted file change detected message with ASCII colors and emoji
	defaultLogger.Printf("\n%s%s %s %s\n",
		bold, darkPurple, msg, reset)
}

// LogOrange logs a message with orange text and bold font
func LogOrange(msg string) {
	// Define ANSI escape codes for orange text and bold font
	orange := "\033[38;5;214m" // 214 is the color code for orange in 256-color mode
	bold := "\033[1m"
	reset := "\033[0m"

	// Log a nicely formatted file change detected message with ASCII colors and emoji
	defaultLogger.Printf("\n%s%s %s %s\n",
		bold, orange, msg, reset)
}
func NoteF(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	orange := "\033[38;5;214m" // 214 is the color code for orange in 256-color mode
	bold := "\033[1m"
	reset := "\033[0m"

	// Log a nicely formatted file change detected message with ASCII colors and emoji
	defaultLogger.Printf("\n%s%s %s %s\n",
		bold, orange, msg, reset)
}

func Note(msg string) {
	orange := "\033[38;5;214m" // 214 is the color code for orange in 256-color mode
	bold := "\033[1m"
	reset := "\033[0m"

	// Log a nicely formatted file change detected message with ASCII colors and emoji
	defaultLogger.Printf("\n%s%s %s %s\n",
		bold, orange, msg, reset)
}
