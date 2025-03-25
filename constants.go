package logger

// ANSI escape codes for colored text and formatting
const (
	// Colors
	red        = "\033[31m"
	darkPurple = "\033[35m"
	orange     = "\033[38;5;214m"
	green      = "\033[32m"
	blue       = "\033[34m"
	cyan       = "\033[1;36m"
	yellow     = "\033[1;33m"
	lightGray  = "\033[37m"
	darkGray   = "\033[90m"
	brightBlue = "\033[94m"
)

const (
	// Formatting
	bold  = "\033[1m"
	reset = "\033[0m"
)

const (
	// Special sequences
	clearScreen = "\033[H\033[2J"
)

// Snippets for formatting
const (
	dividerLine = "========================================" //nolint:all //
)
