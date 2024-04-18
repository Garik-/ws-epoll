package zlog

import (
	"errors"
	"syscall"

	"go.uber.org/zap"
)

type (
	Field  = zap.Field
	Logger = zap.Logger
)

var (
	// Int constructs a field with the given key and value.
	Int = zap.Int
	// String constructs a field with the given key and value.
	String = zap.String
	Err    = zap.Error
)

var defaultLogger *Logger

func init() {
	l, _ := zap.NewDevelopment()
	defaultLogger = l
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(msg string, fields ...Field) {
	defaultLogger.Info(msg, fields...)
}

// Sync calls the underlying Core's Sync method, flushing any buffered log
// entries. Applications should take care to call Sync before exiting.
func Sync() {
	err := defaultLogger.Sync()
	if err != nil && !errors.Is(err, syscall.ENOTTY) {
		panic(err)
	}
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(msg string, fields ...Field) {
	defaultLogger.Debug(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(msg string, fields ...Field) {
	defaultLogger.Error(msg, fields...)
}
