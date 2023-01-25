package zapper

import "go.uber.org/zap/zapcore"

//go:generate stringer -type=Level

// Level zapper levels
type Level zapcore.Level

const (
	Debug  Level = iota // Debug logs are typically voluminous, and are usually disabled in production
	Info                // Info is the default logging priority
	Warn                // Warn logs are more important than Info, but don't need individual human review
	Error               // Error logs are high-priority. If an application is running smoothly, it shouldn't generate any error-level logs
	DPanic              // DPanic logs are particularly important errors. In development the logger panics after writing the message
	Panic               // Panic logs a message, then panics
	Fatal               // Fatal logs a message, then calls os.Exit(1)
)

func (i Level) zapLevel() zapcore.Level {
	return zapcore.Level(i)
}
