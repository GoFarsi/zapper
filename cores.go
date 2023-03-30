package zapper

import (
	"fmt"
	"github.com/TheZeroSlave/zapsentry"
	"github.com/getsentry/sentry-go"
	sentryotel "github.com/getsentry/sentry-go/otel"
	color "github.com/mattn/go-colorable"
	"github.com/natefinch/lumberjack"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
	"time"
)

//go:generate stringer -type=coreType

type coreType int

const (
	CONSOLE coreType = iota
	SENTRY
	FILE
	JSON
)

// Rotation config log file rotation in log path
type Rotation struct {
	MaxAge   int  // MaxAge is the maximum number of days to retain old log files based on the timestamp encoded in their filename. Note that a day is defined as 24 hours and may not exactly correspond to calendar days due to daylight savings, leap seconds, etc. The default is not to remove old log files based on age.
	FileSize int  // FileSize is the maximum size in megabytes of the log file before it gets rotated. It defaults to 100 megabytes
	Compress bool // Compress determines if the rotated log files should be compressed using gzip. The default is not to perform compression
}

// SentryConfig for sentry core set custom configs
type SentryConfig struct {
	AttachStacktrace  bool // AttachStacktrace attach stacktrace to event
	Debug             bool // Debug add debug data to event
	EnableTracing     bool
	Environment       string
	Dist              string
	EnableBreadcrumbs bool
	BreadcrumbLevel   Level
	MaxBreadcrumbs    int
	MaxSpans          int
	Tags              map[string]string
	MinLevel          Level
	FlushTimeout      time.Duration
}

// Core zapper base abstract
type Core interface {
	init(*Zap) error
}

type core struct {
	coreType coreType
	do       func(*Zap) (zapcore.Core, error)
}

// ConsoleWriterCore create console writer for zapper to show log in console
func ConsoleWriterCore(colorable bool) Core {
	return newCore(CONSOLE, func(z *Zap) (zapcore.Core, error) {
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

	return newCore(SENTRY, func(zapper *Zap) (zapcore.Core, error) {
		s, err := sentry.NewClient(sentry.ClientOptions{
			Dsn:              dsn,
			AttachStacktrace: cfg.AttachStacktrace,
			ServerName:       serverName,
			Debug:            cfg.Debug,
			EnableTracing:    cfg.EnableTracing,
			TracesSampleRate: 1.0,
			Environment:      cfg.Environment,
			Dist:             cfg.Dist,
			MaxBreadcrumbs:   cfg.MaxBreadcrumbs,
			MaxSpans:         cfg.MaxSpans,
		})

		if cfg.EnableTracing {
			tp := sdktrace.NewTracerProvider(
				sdktrace.WithSpanProcessor(sentryotel.NewSentrySpanProcessor()),
			)
			otel.SetTracerProvider(tp)
			otel.SetTextMapPropagator(sentryotel.NewSentryPropagator())
		}

		if err != nil {
			return nil, NewError("failed to create sentry client: %s", err.Error())
		}

		c, err := zapsentry.NewCore(zapsentry.Configuration{
			Tags:              cfg.Tags,
			Level:             cfg.MinLevel.zapLevel(),
			EnableBreadcrumbs: cfg.EnableBreadcrumbs,
			BreadcrumbLevel:   cfg.BreadcrumbLevel.zapLevel(),
			DisableStacktrace: false,
			FlushTimeout:      cfg.FlushTimeout,
		}, zapsentry.NewSentryClientFromClient(s))
		if err != nil {
			return nil, NewError("failed to create core sentry: %s", err.Error())
		}

		return c, nil
	})
}

// FileWriterCore write logs into file
func FileWriterCore(logPath string, rotation *Rotation) Core {
	return newCore(FILE, func(z *Zap) (zapcore.Core, error) {
		if _, err := filepath.Abs(logPath); err != nil {
			return nil, NewError("logPath is invalid absolute path: %s", err.Error())
		}

		if rotation == nil {
			rotation = _defaultRotation()
		}

		syncer, err := fileWriteSyncer(logPath, ".log", z.service.Name, rotation)
		if err != nil {
			return nil, NewError("failed to create log path: %s", err.Error())
		}

		return zapcore.NewCore(encoder(z.development, false, z.timeFormat, func(cfg zapcore.EncoderConfig) zapcore.Encoder {
			return zapcore.NewConsoleEncoder(cfg)
		}), zapcore.Lock(syncer), zap.LevelEnablerFunc(z.level)), nil
	})
}

// JsonWriterCore write logs with json format, fileExtension is for set output file extension json,log and etc
func JsonWriterCore(logPath, fileExtension string, rotation *Rotation) Core {
	return newCore(JSON, func(z *Zap) (zapcore.Core, error) {
		if _, err := filepath.Abs(logPath); err != nil {
			return nil, NewError("logPath is invalid absolute path: %s", err.Error())
		}

		if rotation == nil {
			rotation = _defaultRotation()
		}

		syncer, err := fileWriteSyncer(logPath, fileExtension, z.service.Name, rotation)
		if err != nil {
			return nil, NewError("failed to create log path: %s", err.Error())
		}

		return zapcore.NewCore(encoder(z.development, false, z.timeFormat, func(cfg zapcore.EncoderConfig) zapcore.Encoder {
			return zapcore.NewJSONEncoder(cfg)
		}), zapcore.Lock(syncer), zap.LevelEnablerFunc(z.level)), nil
	})
}

func (z *core) init(zapper *Zap) error {
	c, err := z.do(zapper)
	if err != nil {
		return err
	}

	if z.coreType == SENTRY {
		zapper.sentryCore = c
		return nil
	}

	zapper.cores = append(zapper.cores, c)
	return nil
}

func newCore(coreType coreType, f func(*Zap) (zapcore.Core, error)) *core {
	return &core{
		coreType: coreType,
		do:       f,
	}
}

func fileWriteSyncer(logPath, extFile, serviceName string, rotation *Rotation) (zapcore.WriteSyncer, error) {
	if len(serviceName) != 0 {
		serviceName = fmt.Sprintf("%s_", serviceName)
	}

	path := filepath.Join(logPath, fmt.Sprintf("%s%s%s", serviceName, time.Now().Format("2006-01-02_15-04-05"), extFile))

	if len(filepath.Ext(path)) == 0 {
		return nil, NewError("file extension is invalid")
	}

	r := new(lumberjack.Logger)
	r.Filename = path
	r.MaxAge = rotation.MaxAge
	r.MaxSize = rotation.FileSize
	r.Compress = rotation.Compress
	return zapcore.AddSync(r), nil
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
		AttachStacktrace: true,
		Tags: map[string]string{
			"component": "system",
		},
		MinLevel:          Error,
		EnableBreadcrumbs: true,
		BreadcrumbLevel:   Info,
		FlushTimeout:      2 * time.Second,
	}
}
