package logger

import (
	"context"
	"io"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/ingrowco/night-watch/configurator"
)

type loggerKey struct{}

func WithContext(ctx context.Context) context.Context {
	if l := ctx.Value(loggerKey{}); l != nil {
		return ctx
	}

	return context.WithValue(ctx, loggerKey{}, newLogger(ctx, os.Stderr))
}

func FromContext(ctx context.Context) *logrus.Logger {
	return ctx.Value(loggerKey{}).(*logrus.Logger)
}

func New(output io.Writer, formatter logrus.Formatter, hooks logrus.LevelHooks, level string, exitFunc func(int), reportCaller bool) *logrus.Logger {
	logger := &logrus.Logger{
		Out:          output,
		Formatter:    formatter,
		Hooks:        hooks,
		Level:        getLevel(level),
		ExitFunc:     exitFunc,
		ReportCaller: reportCaller,
	}

	return logger
}

func newLogger(ctx context.Context, output io.Writer) *logrus.Logger {
	return New(
		output,
		new(logrus.TextFormatter),
		make(logrus.LevelHooks),
		configurator.FromContext(ctx).GetString("log.level"),
		os.Exit,
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
