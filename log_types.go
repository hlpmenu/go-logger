package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Logging flags
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

// Global logger instance
var defaultLogger = NewLogger()

// Logger provides formatted logging with thread-safe output
type Logger struct {
	mu        sync.Mutex
	out       io.Writer              // destination for output
	prefix    atomic.Pointer[string] // prefix to identify the logger
	flag      atomic.Int32           // properties
	isDiscard atomic.Bool
}

var bufferPool = sync.Pool{New: func() any { return new([]byte) }}

func getBuffer() *[]byte {
	v := bufferPool.Get()
	p, ok := v.(*[]byte)
	if !ok {
		return new([]byte)
	}
	*p = (*p)[:0]
	return p
}

func putBuffer(p *[]byte) {
	if cap(*p) > 64<<10 {
		*p = nil
	}
	bufferPool.Put(p)
}

// Cheap integer to fixed-width decimal ASCII
func itoa(buf *[]byte, i int, wid int) {
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

// NewLogger creates a new Logger
func NewLogger() *Logger {
	l := &Logger{
		out: os.Stdout,
	}
	l.SetFlags(LstdFlags)
	return l
}

func newStdLnLogger() *Logger {
	l := &Logger{
		out: os.Stdout,
	}
	return l
}

// formatHeader writes log header to buf in following order:
// - prefix (if not blank and Lmsgprefix is unset)
// - date and/or time (if corresponding flags are provided)
// - file and line number (if corresponding flags are provided)
// - prefix (if not blank and Lmsgprefix is set)
func formatHeader(buf *[]byte, t time.Time, prefix string, flag int, file string, line int) {
	if flag&Lmsgprefix == 0 {
		*buf = append(*buf, prefix...)
	}

	if flag&(Ldate|Ltime|Lmicroseconds) != 0 {
		if flag&LUTC != 0 {
			t = t.UTC()
		}
		if flag&Ldate != 0 {
			year, month, day := t.Date()
			itoa(buf, year, 4)
			*buf = append(*buf, '/')
			itoa(buf, int(month), 2)
			*buf = append(*buf, '/')
			itoa(buf, day, 2)
			*buf = append(*buf, ' ')
		}
		if flag&(Ltime|Lmicroseconds) != 0 {
			hour, minute, sec := t.Clock()
			itoa(buf, hour, 2)
			*buf = append(*buf, ':')
			itoa(buf, minute, 2)
			*buf = append(*buf, ':')
			itoa(buf, sec, 2)
			if flag&Lmicroseconds != 0 {
				*buf = append(*buf, '.')
				itoa(buf, t.Nanosecond()/1e3, 6)
			}
			*buf = append(*buf, ' ')
		}
	}

	if flag&(Lshortfile|Llongfile) != 0 {
		if flag&Lshortfile != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		*buf = append(*buf, file...)
		*buf = append(*buf, ':')
		itoa(buf, line, -1)
		*buf = append(*buf, ": "...)
	}

	if flag&Lmsgprefix != 0 {
		*buf = append(*buf, prefix...)
	}
}

// output formats and writes the output for a logging event.
//
// This is an internal method that handles the core logging functionality:
// 1. Retrieves a buffer from the pool (getBuffer guarantees non-nil return)
// 2. Formats the log header with timestamp, prefix, file, line, etc.
// 3. Appends the actual log message via the appendOutput function
// 4. Ensures the message ends with a newline
// 5. Writes the message to the configured output destination
//
// The error returned is from the underlying Write operation to l.out.
// In most cases, callers can safely ignore this error since:
// - For common io.Writers like os.Stdout, errors are rare and typically fatal
// - For memory-based writers like bytes.Buffer, errors will never occur
// - If logging fails, there's often nowhere else to report the failure
// =========== Above commment is ai generated since i kept getting scared of the no-nil-checks lol ==========
func (l *Logger) output(pc uintptr, calldepth int, appendOutput func([]byte) []byte) {
	if l.isDiscard.Load() {
		return
	}

	now := time.Now()

	// Load prefix and flag once for consistency
	prefix := l.Prefix()
	flag := l.Flags()

	var file string
	var line int
	if flag&(Lshortfile|Llongfile) != 0 {
		if pc == 0 {
			var ok bool
			_, file, line, ok = runtime.Caller(calldepth)
			if !ok {
				file = "???"
				line = 0
			}
		} else {
			fs := runtime.CallersFrames([]uintptr{pc})
			f, _ := fs.Next()
			file = f.File
			if file == "" {
				file = "???"
			}
			line = f.Line
		}
	}

	buf := getBuffer() // Guaranteed not nil, it checks pool result and never returns nil

	defer putBuffer(buf) // Return the buffer to the pool after use

	formatHeader(buf, now, prefix, flag, file, line)
	*buf = appendOutput(*buf)
	if len(*buf) == 0 || (*buf)[len(*buf)-1] != '\n' {
		*buf = append(*buf, '\n')
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	// Removed error check
	//#nosec G104 -- see below
	l.out.Write(*buf) //nolint:errcheck //
	// Since were simply using io.Writer over append() for performance reasons
	// Originally returning a error was a mistake.
	// Compared with bytes.Buffer, which while it does return a error
	// Has "return nil" hardcoded as the only return.
	// Also as were just doing the same operation as a append() call wouldve done
	// which also do not return an error. Theres no reason to do so.
	// Lastly, the only places this internal method is used, has no use for the errors.
}

// SetOutput sets the output destination
func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out = w
	l.isDiscard.Store(w == io.Discard)
}

// SetPrefix sets the prefix for log messages
func (l *Logger) SetPrefix(prefix string) {
	l.prefix.Store(&prefix)
}

// SetFlags sets the logging flags
func (l *Logger) SetFlags(flag int) {
	if flag < 0 || flag > (1<<31-1) {
		return
	}
	l.flag.Store(int32(flag))
}

// Flags returns the output flags
func (l *Logger) Flags() int {
	return int(l.flag.Load())
}

// Prefix returns the output prefix
func (l *Logger) Prefix() string {
	if p := l.prefix.Load(); p != nil {
		return *p
	}
	return ""
}

// Print logs a message
func (l *Logger) Print(v ...interface{}) {
	l.output(0, 2, func(b []byte) []byte {
		return fmt.Append(b, v...)
	})
}

// Printf logs a formatted message
func (l *Logger) Printf(format string, v ...interface{}) {
	l.output(0, 2, func(b []byte) []byte {
		return fmt.Appendf(b, format, v...)
	})
}

// Println logs a message with a newline
func (l *Logger) Println(v ...interface{}) {
	l.output(0, 2, func(b []byte) []byte {
		return fmt.Appendln(b, v...)
	})
}

// Fatal logs a message and then calls os.Exit(1)
func (l *Logger) Fatal(v ...interface{}) {
	l.output(0, 2, func(b []byte) []byte {
		return fmt.Append(b, v...)
	})
	os.Exit(1)
}

// Fatalf logs a formatted message and then calls os.Exit(1)
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.output(0, 2, func(b []byte) []byte {
		return fmt.Appendf(b, format, v...)
	})
	os.Exit(1)
}

// Fatalln logs a message with a newline and then calls os.Exit(1)
func (l *Logger) Fatalln(v ...interface{}) {
	l.output(0, 2, func(b []byte) []byte {
		return fmt.Appendln(b, v...)
	})
	os.Exit(1)
}

// Panic logs a message and then panics
func (l *Logger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.output(0, 2, func(b []byte) []byte {
		return append(b, s...)
	})
	panic(s)
}

// Panicf logs a formatted message and then panics
func (l *Logger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.output(0, 2, func(b []byte) []byte {
		return append(b, s...)
	})
	panic(s)
}

// Panicln logs a message with a newline and then panics
func (l *Logger) Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	l.output(0, 2, func(b []byte) []byte {
		return append(b, s...)
	})
	panic(s)
}

// Output writes the output for a logging event
func (l *Logger) Output(calldepth int, s string) {
	calldepth++ // +1 for this frame
	l.output(0, calldepth, func(b []byte) []byte {
		return append(b, s...)
	})
}

// Package-level functions that use the default logger

func Print(v ...interface{}) {
	defaultLogger.Print(v...)
}

func Printf(format string, v ...interface{}) {
	defaultLogger.Printf(format, v...)
}

func Fatal(v ...interface{}) {
	defaultLogger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	defaultLogger.Fatalln(v...)
}

func Panic(v ...interface{}) {
	defaultLogger.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	defaultLogger.Panicf(format, v...)
}

func Panicln(v ...interface{}) {
	defaultLogger.Panicln(v...)
}
