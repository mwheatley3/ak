package pg

import (
	"github.com/Sirupsen/logrus"
)

type logger struct {
	l *logrus.Logger
}

func (l *logger) Debug(msg string, ctx ...interface{}) {
	ctx = append([]interface{}{msg}, ctx...)
	l.l.Debug(ctx...)
}

func (l *logger) Info(msg string, ctx ...interface{}) {
	ctx = append([]interface{}{msg}, ctx...)
	l.l.Info(ctx...)
}

func (l *logger) Warn(msg string, ctx ...interface{}) {
	ctx = append([]interface{}{msg}, ctx...)
	l.l.Warn(ctx...)
}

func (l *logger) Error(msg string, ctx ...interface{}) {
	ctx = append([]interface{}{msg}, ctx...)
	l.l.Error(ctx...)
}
