package logger

var lnLogger = newStdLnLogger()

// Println logs a message with a newline
// Println does not include any timestamps or prefixes
func Println(v ...any) {
	lnLogger.Println(v...)
}

func Printfln(format string, v ...any) {
	lnLogger.Printf(format, v...)
}

func NewStdLnLogger() *Logger {
	return newStdLnLogger()
}
