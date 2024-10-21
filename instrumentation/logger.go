package instrumentation

import (
	"log/slog"
	"os"
)

type GoLogger struct {
	logger *slog.Logger
}

func InitInstruments() GoLogger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return GoLogger{logger}
}

func (l *GoLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *GoLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *GoLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *GoLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}
