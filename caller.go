package zapper

type Caller interface {
	Debug(...any)
	DebugF(string, ...any)
	DebugW(string, ...any)

	Info(...any)
	InfoF(string, ...any)
	InfoW(string, ...any)

	Warn(...any)
	WarnF(string, ...any)
	WarnW(string, ...any)

	Error(...any)
	ErrorF(string, ...any)
	ErrorW(string, ...any)

	DPanic(...any)
	DPanicF(string, ...any)
	DPanicW(string, ...any)

	Panic(...any)
	PanicF(string, ...any)
	PanicW(string, ...any)

	Fatal(...any)
	FatalF(string, ...any)
	FatalW(string, ...any)
}

// Debug uses fmt.Sprint to construct and log a message
func (z *Zap) Debug(args ...any) {
	z.debug(z.sugar, args...)
}

// DebugF uses fmt.Sprintf to log a templated message
func (z *Zap) DebugF(message string, args ...any) {
	z.debugF(z.sugar, message, args...)
}

// DebugW logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(keysAndValues).Debug(msg)
func (z *Zap) DebugW(message string, keyAndValues ...any) {
	z.debugW(z.sugar, message, keyAndValues...)
}

// Info uses fmt.Sprint to construct and log a message
func (z *Zap) Info(args ...any) {
	z.info(z.sugar, args...)
}

// InfoF uses fmt.Sprintf to log a templated message
func (z *Zap) InfoF(message string, args ...any) {
	z.infoF(z.sugar, message, args...)
}

// InfoW logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With
func (z *Zap) InfoW(message string, keyAndValues ...any) {
	z.infoW(z.sugar, message, keyAndValues...)
}

// Warn uses fmt.Sprint to construct and log a message
func (z *Zap) Warn(args ...any) {
	z.warn(z.sugar, args...)
}

// WarnF uses fmt.Sprintf to log a templated message
func (z *Zap) WarnF(message string, args ...any) {
	z.warnF(z.sugar, message, args...)
}

// WarnW logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With
func (z *Zap) WarnW(message string, keyAndValues ...any) {
	z.warnW(z.sugar, message, keyAndValues...)
}

// Error uses fmt.Sprint to construct and log a message
func (z *Zap) Error(args ...any) {
	z.error(z.sugar, args...)
}

// ErrorF uses fmt.Sprintf to log a templated message
func (z *Zap) ErrorF(message string, args ...any) {
	z.errorF(z.sugar, message, args...)
}

// ErrorW logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With
func (z *Zap) ErrorW(message string, keyAndValues ...any) {
	z.errorW(z.sugar, message, keyAndValues...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (z *Zap) DPanic(args ...any) {
	z.dPanic(z.sugar, args...)
}

// DPanicF uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (z *Zap) DPanicF(message string, args ...any) {
	z.dPanicF(z.sugar, message, args...)
}

// DPanicW logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With
func (z *Zap) DPanicW(message string, keyAndValues ...any) {
	z.dPanicW(z.sugar, message, keyAndValues...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (z *Zap) Panic(args ...any) {
	z.panic(z.sugar, args...)
}

// PanicF uses fmt.Sprintf to log a templated message, then panics
func (z *Zap) PanicF(message string, args ...any) {
	z.panicF(z.sugar, message, args...)
}

// PanicW logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With
func (z *Zap) PanicW(message string, keyAndValues ...any) {
	z.panicW(z.sugar, message, keyAndValues...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit
func (z *Zap) Fatal(args ...any) {
	z.fatal(z.sugar, args...)
}

// FatalF uses fmt.Sprintf to log a templated message, then calls os.Exit
func (z *Zap) FatalF(message string, args ...any) {
	z.fatalF(z.sugar, message, args...)
}

// FatalW logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With
func (z *Zap) FatalW(message string, keyAndValues ...any) {
	z.fatalW(z.sugar, message, keyAndValues...)
}
