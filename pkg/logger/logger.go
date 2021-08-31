package logger

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

type Option func(l *Logger)

type Logger struct {
	logr logr.Logger
	cnf  zap.Config
}

func New(options ...Option) *Logger {
	lgr := &Logger{
		cnf: zap.NewDevelopmentConfig(),
	}
	for _, opt := range options {
		opt(lgr)
	}
	l, err := lgr.cnf.Build()
	if err != nil {
		panic(err)
	}
	lgr.logr = zapr.NewLogger(l)
	return lgr
}

func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	l.logr.Info(msg, keysAndValues...)
}

func (l *Logger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.logr.Error(err, msg, keysAndValues...)
}

func (l *Logger) Fork(keysAndValues ...interface{}) *Logger {
	n := l.logr.WithValues(keysAndValues...)
	return &Logger{
		logr: n,
		cnf:  l.cnf,
	}
}

func WithoutStd() Option {
	return func(l *Logger) {
		l.cnf.OutputPaths = []string{}
		l.cnf.ErrorOutputPaths = []string{}
	}
}

func WithFile(p string) Option {
	return func(l *Logger) {
		l.cnf.OutputPaths = append(l.cnf.OutputPaths, p)
		l.cnf.ErrorOutputPaths = append(l.cnf.ErrorOutputPaths, p)
	}
}
