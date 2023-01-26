package zapper

import (
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap/zapcore"
)

//go:generate stringer -type=Level

// Level zapper levels
type Level zapcore.Level

const (
	Debug  Level = iota - 1 // Debug logs are typically voluminous, and are usually disabled in production
	Info                    // Info is the default logging priority
	Warn                    // Warn logs are more important than Info, but don't need individual human review
	Error                   // Error logs are high-priority. If an application is running smoothly, it shouldn't generate any error-level logs
	DPanic                  // DPanic logs are particularly important errors. In development the logger panics after writing the message
	Panic                   // Panic logs a message, then panics
	Fatal                   // Fatal logs a message, then calls os.Exit(1)
)

func (l Level) zapLevel() zapcore.Level {
	return zapcore.Level(l)
}

func (l Level) sentryLevel() sentry.Level {
	switch l {
	case Debug:
		return sentry.LevelDebug
	case Info:
		return sentry.LevelInfo
	case Warn:
		return sentry.LevelWarning
	case Error:
		return sentry.LevelError
	default:
		return sentry.LevelFatal
	}
}
