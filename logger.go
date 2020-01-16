// Package hclog provides a Logur adapter for hclog.
package hclog

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"logur.dev/logur"

	"logur.dev/adapter/hclog/internal/keyvals"
)

// Logger is a Logur adapter for hclog.
type Logger struct {
	logger hclog.Logger
}

// New returns a new Logur logger.
// If logger is nil, a default instance is created.
func New(logger hclog.Logger) *Logger {
	if logger == nil {
		logger = hclog.Default()
	}

	return &Logger{
		logger: logger,
	}
}

// Trace implements the Logur Logger interface.
func (l *Logger) Trace(msg string, fields ...map[string]interface{}) {
	if !l.logger.IsTrace() {
		return
	}

	l.logger.Trace(msg, l.keyvals(fields)...)
}

// Debug implements the Logur Logger interface.
func (l *Logger) Debug(msg string, fields ...map[string]interface{}) {
	if !l.logger.IsDebug() {
		return
	}

	l.logger.Debug(msg, l.keyvals(fields)...)
}

// Info implements the Logur Logger interface.
func (l *Logger) Info(msg string, fields ...map[string]interface{}) {
	if !l.logger.IsInfo() {
		return
	}

	l.logger.Info(msg, l.keyvals(fields)...)
}

// Warn implements the Logur Logger interface.
func (l *Logger) Warn(msg string, fields ...map[string]interface{}) {
	if !l.logger.IsWarn() {
		return
	}

	l.logger.Warn(msg, l.keyvals(fields)...)
}

// Error implements the Logur Logger interface.
func (l *Logger) Error(msg string, fields ...map[string]interface{}) {
	if !l.logger.IsError() {
		return
	}

	l.logger.Error(msg, l.keyvals(fields)...)
}

func (l *Logger) keyvals(fields []map[string]interface{}) []interface{} {
	var kvs []interface{}
	if len(fields) > 0 {
		kvs = keyvals.FromMap(fields[0])
	}

	return kvs
}

func (l *Logger) TraceContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Trace(msg, fields...)
}

func (l *Logger) DebugContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Debug(msg, fields...)
}

func (l *Logger) InfoContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Info(msg, fields...)
}

func (l *Logger) WarnContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Warn(msg, fields...)
}

func (l *Logger) ErrorContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Error(msg, fields...)
}

// LevelEnabled implements the Logur LevelEnabler interface.
func (l *Logger) LevelEnabled(level logur.Level) bool {
	switch level {
	case logur.Trace:
		return l.logger.IsTrace()
	case logur.Debug:
		return l.logger.IsDebug()
	case logur.Info:
		return l.logger.IsInfo()
	case logur.Warn:
		return l.logger.IsWarn()
	case logur.Error:
		return l.logger.IsError()
	}

	return true
}
