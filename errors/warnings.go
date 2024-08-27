package errutil

import "gopkg.hlmpn.dev/pkg/go-logger"

type Warning struct {
	Message string
}

func (w *Warning) Log() {
	logger.Warning(w.Message)
}

func (w *Warning) Notice() {
	logger.LogOrange("NOTICE: " + w.Message)
}
func (w *Warning) Warn() {
	logger.Warning(w.Message)
}

func NewNotice(msg string) *Warning {
	return &Warning{
		Message: "NOTICE: " + msg,
	}
}
