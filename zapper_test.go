package zapper

import "testing"

func Test_ConsoleWriterCore(t *testing.T) {
	tests := []struct {
		name        string
		development bool
		debug       bool
		colorable   bool
		timeFormat  TimeFormat
		stackLevel  Level
	}{
		{
			name:        "default",
			development: false,
			debug:       false,
			colorable:   false,
			timeFormat:  ISO8601,
		},
		{
			name:        "development mode",
			development: true,
			debug:       false,
			colorable:   false,
			timeFormat:  RFC3339,
		},
		{
			name:        "with debug level",
			development: false,
			debug:       true,
			colorable:   false,
			timeFormat:  RFC3339NANO,
		},
		{
			name:        "colorable levels",
			development: false,
			debug:       false,
			colorable:   true,
		},
		{
			name:        "development mode with debug",
			development: true,
			debug:       true,
			colorable:   false,
		},
		{
			name:        "debug with colorable levels",
			development: false,
			debug:       true,
			colorable:   true,
			stackLevel:  Warn,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := make([]Option, 0)
			opts = append(opts, WithTimeFormat(tt.timeFormat))
			opts = append(opts, WithCustomStackTraceLevel(tt.stackLevel))
			if tt.debug {
				opts = append(opts, WithDebugLevel())
			}
			zapper := New(tt.development, opts...)
			if err := zapper.NewCore(ConsoleWriterCore(tt.colorable)); err != nil {
				t.Fatal(err)
			}

			zapper.Debug("test debug")
			zapper.Info("test info")
			zapper.Warn("test warn")
			zapper.Error("test error")
		})
	}
}

func Test_SentryCore(t *testing.T) {
	zapper := New(false, WithCustomStackTraceLevel(Warn))
	if err := zapper.NewCore(SentryCore("https://43bc037c7b504219ac3f3cca1f113ea4@o334462.ingest.sentry.io/4504569607159808", "test", nil)); err != nil {
		t.Fatal(err)
	}

	zapper.Error("sentry error 2")
}
