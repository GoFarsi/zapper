package zapper

import (
	"github.com/TheZeroSlave/zapsentry"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	logFunc   func(zap *zap.SugaredLogger, args ...any)
	logfFunc  func(zap *zap.SugaredLogger, msg string, args ...any)
	logKvFunc func(zap *zap.SugaredLogger, msg string, keyAndValues ...any)
)

type Zapper interface {
	Caller

	NewCore(...Core) error
	GetServiceCode() uint
	GetServiceName() string
}

type Zap struct {
	zap        *zap.Logger
	sugar      *zap.SugaredLogger
	cores      []zapcore.Core
	sentryCore zapcore.Core

	development bool
	timeFormat  TimeFormat
	stackLevel  Level
	level       func(lvl zapcore.Level) bool
	service     *service

	debug  logFunc
	debugF logfFunc
	debugW logKvFunc

	info  logFunc
	infoF logfFunc
	infoW logKvFunc

	warn  logFunc
	warnF logfFunc
	warnW logKvFunc

	error  logFunc
	errorF logfFunc
	errorW logKvFunc

	dPanic  logFunc
	dPanicF logfFunc
	dPanicW logKvFunc

	panic  logFunc
	panicF logfFunc
	panicW logKvFunc

	fatal  logFunc
	fatalF logfFunc
	fatalW logKvFunc
}

type service struct {
	Code uint
	Name string
}

// New create new Zap object
func New(development bool, opts ...Option) Zapper {
	zapper := new(Zap)

	zapper.development = development
	zapper.service = new(service)
	zapper.stackLevel = Error
	zapper.level = func(lvl zapcore.Level) bool {
		return lvl >= Debug.zapLevel()
	}

	for _, opt := range opts {
		opt(zapper)
	}

	zapper.callersLoader()

	return zapper
}

// NewCore create cores for zapper
func (z *Zap) NewCore(cores ...Core) error {
	if len(cores) == 0 {
		cores = append(cores, _defaultCore())
	}

	useSentry := false

	for _, c := range cores {
		s, ok := c.(*core)
		if ok && s.coreType == SENTRY {
			useSentry = true
		}
		if err := c.init(z); err != nil {
			return err
		}
	}

	z.zap = zap.New(zapcore.NewTee(z.cores...), zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(z.stackLevel.zapLevel()))

	if useSentry {
		z.zap = zapsentry.AttachCoreToLogger(z.sentryCore, z.zap)
		z.zap = z.zap.With(zapsentry.NewScope())
	}

	z.sugar = z.zap.Sugar()

	return nil
}

func (z *Zap) GetServiceCode() uint {
	return z.service.Code
}

func (z *Zap) GetServiceName() string {
	return z.service.Name
}

func (z *Zap) GetZap() *zap.Logger {
	return z.zap
}

func (z *Zap) callersLoader() {
	z.debug = (*zap.SugaredLogger).Debug
	z.debugF = (*zap.SugaredLogger).Debugf
	z.debugW = (*zap.SugaredLogger).Debugw
	z.info = (*zap.SugaredLogger).Info
	z.infoF = (*zap.SugaredLogger).Infof
	z.infoW = (*zap.SugaredLogger).Infow
	z.warn = (*zap.SugaredLogger).Warn
	z.warnF = (*zap.SugaredLogger).Warnf
	z.warnW = (*zap.SugaredLogger).Warnw
	z.error = (*zap.SugaredLogger).Error
	z.errorF = (*zap.SugaredLogger).Errorf
	z.errorW = (*zap.SugaredLogger).Errorw
	z.dPanic = (*zap.SugaredLogger).DPanic
	z.dPanicF = (*zap.SugaredLogger).DPanicf
	z.dPanicW = (*zap.SugaredLogger).DPanicw
	z.panic = (*zap.SugaredLogger).Panic
	z.panicF = (*zap.SugaredLogger).Panicf
	z.panicW = (*zap.SugaredLogger).Panicw
	z.fatal = (*zap.SugaredLogger).Fatal
	z.fatalF = (*zap.SugaredLogger).Fatalf
	z.fatalW = (*zap.SugaredLogger).Fatalw
}

func _defaultCore() Core {
	return ConsoleWriterCore(true)
}
