package zapper

type Option func(*Zapper)

// WithMinLevel set minimum level logger
func WithMinLevel(min Level) Option {
	return func(z *Zapper) {
		z.level = min
	}
}

// WithTimeFormat set custom time format for zapper logs
func WithTimeFormat(format TimeFormat) Option {
	return func(z *Zapper) {
		z.timeFormat = format
	}
}
