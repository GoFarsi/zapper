package zapper

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	logFunc   func(zap *zap.SugaredLogger, args ...any)
	logfFunc  func(zap *zap.SugaredLogger, msg string, args ...any)
	logKvFunc func(zap *zap.SugaredLogger, msg string, keyAndValues ...any)
)

type Zapper struct {
	zap   *zap.Logger
	sugar *zap.SugaredLogger
	cores []zapcore.Core

	development bool
	timeFormat  TimeFormat
	level       Level

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

// New create new Zapper object
func New(development bool, opts ...Option) *Zapper {
	zapper := new(Zapper)

	zapper.development = development
	zapper.level = Info

	for _, opt := range opts {
		opt(zapper)
	}

	zapper.funcLoader()

	return zapper
}

// NewCore create cores for zapper
func (z *Zapper) NewCore(cores ...Core) error {
	for _, core := range cores {
		if err := core.init(z); err != nil {
			return err
		}
	}

	z.zap = zap.New(zapcore.NewTee(z.cores...), zap.AddCaller(), zap.AddStacktrace(Error.zapLevel()))
	z.sugar = z.zap.Sugar()

	return nil
}

func (z *Zapper) funcLoader() {
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
