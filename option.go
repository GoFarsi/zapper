package zapper

import (
	"go.uber.org/zap/zapcore"
)

type Option func(*Zap)

// WithDebugLevel enable debug level for logging
func WithDebugLevel() Option {
	return func(z *Zap) {
		z.level = func(lvl zapcore.Level) bool {
			return lvl <= Fatal.zapLevel()
		}
	}
}

// WithTimeFormat set custom time format for zapper logs
func WithTimeFormat(format TimeFormat) Option {
	return func(z *Zap) {
		z.timeFormat = format
	}
}

// WithCustomStackTraceLevel set custom level for show stacktrace, min level warn
func WithCustomStackTraceLevel(level Level) Option {
	return func(z *Zap) {
		if level < Warn {
			level = Error
		}

		z.stackLevel = level
	}
}
