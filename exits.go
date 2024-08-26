package logger

import "os"

func Exit(msg string, code int) {
	defaultLogger.Print(msg)
	os.Exit(code)
}

func Exitf(format string, code int, v ...interface{}) {
	msg := format
	if len(v) > 0 {
		msg = format
	}
	defaultLogger.Print(msg)
	os.Exit(code)
}

func LogURL(url string) {
	emoji := "🔗"
	defaultLogger.Printf("%s %s", emoji, url)
}

func Clear() {
	defaultLogger.Print("\033[H\033[2J")
}

func Log(msg string) {
	defaultLogger.Print(msg)
}

func Logf(format string, v ...interface{}) {
	defaultLogger.Printf(format, v...)
}
