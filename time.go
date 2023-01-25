package zapper

//go:generate stringer -type=TimeFormat

// TimeFormat set custom time format for zap log
type TimeFormat int

const (
	ISO8601     TimeFormat = iota // ISO8601 serializes a time.Time to an ISO8601-formatted string with millisecond precision
	RFC3339                       // RFC3339 serializes a time.Time to an RFC3339-formatted string
	RFC3339NANO                   // RFC3339NANO  serializes a time.Time to an RFC3339-formatted string with nanosecond precision
)
