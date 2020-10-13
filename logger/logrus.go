package logger

import (
	"context"
	"io"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/ingrowco/night-watch/configurator"
)

type loggerKey struct{}

// WithContext saves the logger inside of a given context
func WithContext(ctx context.Context) context.Context {
	if l := ctx.Value(loggerKey{}); l != nil {
		return ctx
	}

	return context.WithValue(ctx, loggerKey{}, newLogger(ctx, os.Stderr))
}

// FromContext returns a saved logger that saved inside a context. If there is no logger has been saved, it will be panic
func FromContext(ctx context.Context) *logrus.Logger {
	return ctx.Value(loggerKey{}).(*logrus.Logger)
}

// New creates an instance of logger
func New(output io.Writer, formatter logrus.Formatter, level string, reportCaller bool) *logrus.Logger {
	logger := &logrus.Logger{
		Out:          output,
		Formatter:    formatter,
		Hooks:        make(logrus.LevelHooks),
		Level:        getLevel(level),
		ExitFunc:     os.Exit,
		ReportCaller: reportCaller,
	}

	return logger
}

func newLogger(ctx context.Context, output io.Writer) *logrus.Logger {
	return New(
		output,
		new(logrus.TextFormatter),
		configurator.FromContext(ctx).GetString("log.level"),
		false,
	)
}

func getLevel(level string) logrus.Level {
	parsedLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return logrus.InfoLevel
	}

	return parsedLevel
}
