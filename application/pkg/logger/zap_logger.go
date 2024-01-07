package logger

import (
	"go.uber.org/zap"
)

type zapLogger struct {
	l *zap.Logger
}

func (z zapLogger) Debug(msg string, args ...Field) {
	z.l.Debug(msg, toZapFields(args...)...)
}

func (z zapLogger) Info(msg string, args ...Field) {
	z.l.Info(msg, toZapFields(args...)...)
}

func (z zapLogger) Warn(msg string, args ...Field) {
	z.l.Warn(msg, toZapFields(args...)...)
}

func (z zapLogger) Error(msg string, args ...Field) {
	z.l.Error(msg, toZapFields(args...)...)

}

func toZapFields(args ...Field) []zap.Field {
	var fields []zap.Field
	for _, arg := range args {
		fields = append(fields, zap.Any(arg.Key, arg.Value))
	}
	return fields
}
