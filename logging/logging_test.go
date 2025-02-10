package logging

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	f()
	return buf.String()
}

func TestLoggingLevels(t *testing.T) {
	tests := []struct {
		name      string
		logLevel  string
		logFunc   func()
		shouldLog bool
	}{
		{"DebugLevel_DebugLog", "debug", func() { Debug("debug message") }, true},
		{"InfoLevel_InfoLog", "info", func() { Info("info message") }, true},
		{"WarningLevel_WarnLog", "warning", func() { Warn("warning message") }, true},
		{"ErrorLevel_ErrorLog", "error", func() { Error("error message") }, true},
		{"FatalLevel_FatalLog", "fatal", func() { Fatal("fatal message") }, true},
		{"FatalLevel_DebugLog", "fatal", func() { Debug("should not log") }, false},
		{"InfoLevel_DebugLog", "info", func() { Debug("should not log") }, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("LOG_LEVEL", tt.logLevel)
			Init("")
			output := captureOutput(tt.logFunc)

			if tt.shouldLog && output == "" {
				t.Errorf("Expected log output, but got none")
			}
			if !tt.shouldLog && output != "" {
				t.Errorf("Expected no log output, but got: %s", output)
			}
		})
	}
}
