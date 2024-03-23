package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

type Logger interface {
	Info(msg string)
	Infof(msg string, args ...interface{})
	Error(msg string)
	Errorf(msg string, args ...interface{})
}

type logger struct {
	l *slog.Logger
}

func New(logWriters ...io.Writer) *logger {
	writer := io.MultiWriter(logWriters...)
	writer = io.MultiWriter(writer, os.Stdout)

	l := slog.New(slog.NewJSONHandler(writer, nil))

	return &logger{l: l}
}

func (l *logger) Info(msg string) {
	l.l.Info(msg)
}

func (l *logger) Infof(msg string, args ...interface{}) {
	l.l.Info(fmt.Sprintf(msg, args...))
}

func (l *logger) Error(msg string) {
	l.l.Error(msg)
}

func (l *logger) Errorf(msg string, args ...interface{}) {
	l.l.Error(fmt.Sprintf(msg, args...))
}

func (l *logger) SetServiceName(serviceName string) {
	l.l = l.l.WithGroup(serviceName)
}
