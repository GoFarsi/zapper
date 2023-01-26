package zapper

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type encoderFunc func(zapcore.EncoderConfig) zapcore.Encoder

func encoder(development, colorable bool, timeFormat TimeFormat, encoderFunc encoderFunc) zapcore.Encoder {
	var cfg zapcore.EncoderConfig

	if development {
		cfg = zap.NewDevelopmentEncoderConfig()
		cfg.EncodeCaller = zapcore.FullCallerEncoder
	} else {
		cfg = zap.NewProductionEncoderConfig()
		if colorable {
			cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		} else {
			cfg.EncodeLevel = zapcore.CapitalLevelEncoder
		}
	}

	switch timeFormat {
	case ISO8601:
		cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	case RFC3339:
		cfg.EncodeTime = zapcore.RFC3339TimeEncoder
	case RFC3339NANO:
		cfg.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	}

	cfg.StacktraceKey = "stacktrace"
	cfg.FunctionKey = "function"

	return encoderFunc(cfg)
}
