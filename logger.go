package kura

import (
	"fmt"
	"io"

	"github.com/y-yagi/color"
)

// Logger is a logger for kura.
type Logger struct {
	w io.Writer
}

var (
	green = color.New(color.FgGreen, color.Bold).SprintFunc()
)

// NewLogger creates a new Logger.
func NewLogger(w io.Writer) *Logger {
	l := &Logger{w: w}
	return l
}

// Printf print log with format.
func (l *Logger) Printf(action, format string, a ...interface{}) {
	log := fmt.Sprintf("%s ", green(action))
	log += fmt.Sprintf(format, a...)
	fmt.Fprint(l.w, log)
}
