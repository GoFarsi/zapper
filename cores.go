package zapper

import (
	"github.com/TheZeroSlave/zapsentry"
	"github.com/getsentry/sentry-go"
	color "github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Rotation config log file rotation in log path
type Rotation struct {
	MaxAge   int
	FileSize int
	Compress bool
}

// SentryConfig for sentry core set custom configs
type SentryConfig struct {
	AttachStacktrace  bool
	Debug             bool
	EnableTracing     bool
	Environment       string
	Dist              string
	EnableBreadcrumbs bool
	BreadcrumbLevel   Level
	MaxBreadcrumbs    int
	MaxSpans          int
	Tags              map[string]string
	MinLevel          Level
}

// Core zapper base abstract
type Core interface {
	init(*Zap) error
}

type core struct {
	do func(*Zap) (zapcore.Core, error)
}

// ConsoleWriterCore create console writer for zapper to show log in console
func ConsoleWriterCore(colorable bool) Core {
	return newCore(func(z *Zap) (zapcore.Core, error) {
		return zapcore.NewCore(encoder(z.development, colorable, z.timeFormat, func(cfg zapcore.EncoderConfig) zapcore.Encoder {
			return zapcore.NewConsoleEncoder(cfg)
		}), zapcore.AddSync(color.NewColorableStdout()), zap.LevelEnablerFunc(z.level)), nil
	})
}

// SentryCore send log into sentry service
func SentryCore(dsn string, serverName string, cfg *SentryConfig) Core {
	if cfg == nil {
		cfg = _defaultSentryConfig()
	}

	return newCore(func(zapper *Zap) (zapcore.Core, error) {
		s, err := sentry.NewClient(sentry.ClientOptions{
			Dsn:              dsn,
			AttachStacktrace: cfg.AttachStacktrace,
			ServerName:       serverName,
			Debug:            cfg.Debug,
			EnableTracing:    cfg.EnableTracing,
			Environment:      cfg.Environment,
			Dist:             cfg.Dist,
			MaxBreadcrumbs:   cfg.MaxBreadcrumbs,
			MaxSpans:         cfg.MaxSpans,
		})

		if err != nil {
			return nil, NewError("failed to create sentry client: %s", err.Error())
		}

		core, err := zapsentry.NewCore(zapsentry.Configuration{
			Tags:              cfg.Tags,
			Level:             cfg.MinLevel.zapLevel(),
			EnableBreadcrumbs: cfg.EnableBreadcrumbs,
			BreadcrumbLevel:   cfg.BreadcrumbLevel.zapLevel(),
		}, zapsentry.NewSentryClientFromClient(s))
		if err != nil {
			return nil, NewError("failed to create core sentry: %s", err.Error())
		}

		return core, nil
	})
}

func (z *core) init(zapper *Zap) error {
	core, err := z.do(zapper)
	if err != nil {
		return err
	}
	zapper.cores = append(zapper.cores, core)
	return nil
}

func newCore(f func(*Zap) (zapcore.Core, error)) *core {
	return &core{f}
}

func _defaultRotation() *Rotation {
	return &Rotation{
		MaxAge:   1,
		FileSize: 10,
		Compress: false,
	}
}

func _defaultSentryConfig() *SentryConfig {
	return &SentryConfig{
		AttachStacktrace:  true,
		EnableTracing:     false,
		Debug:             false,
		Environment:       "",
		Dist:              "",
		Tags:              make(map[string]string),
		MinLevel:          Error,
		EnableBreadcrumbs: true,
		BreadcrumbLevel:   Info,
	}
}
