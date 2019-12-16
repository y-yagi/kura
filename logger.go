package kura

import (
	"fmt"
	"io"

	"github.com/y-yagi/color"
)

type KuraLogger struct {
	w io.Writer
}

var (
	green = color.New(color.FgGreen, color.Bold).SprintFunc()
)

// NewLogger creates a new Logger.
func NewLogger(w io.Writer) *KuraLogger {
	l := &KuraLogger{w: w}
	return l
}

// Printf print log with format.
func (l *KuraLogger) Printf(action, format string, a ...interface{}) {
	log := fmt.Sprintf("%s ", green(action))
	log += fmt.Sprintf(format, a...)
	fmt.Fprintf(l.w, log)
}
