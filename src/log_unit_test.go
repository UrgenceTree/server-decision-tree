package main

import (
	"testing"
)

type testLogger struct {
	LastLog string
}

func (t *testLogger) Init() error {
	return nil
}

func (t *testLogger) Term() {}

func (t *testLogger) PrintLog(value string) {
	t.LastLog = value
}

func TestAddLogger(t *testing.T) {
	tl := &testLogger{}
	AddLogger(tl)

	if len(logHandler) != 1 {
		t.Errorf("Expected 1 logger, got %d", len(logHandler))
	}
}

func TestSetLogLevelString(t *testing.T) {
	SetLogLevelString("DEBUG")
	if gLogLevel != LogSeverityDebug {
		t.Errorf("Expected LogSeverityDebug, got %d", gLogLevel)
	}
	// Repeat for all valid log levels.
	// Add a test case for invalid log level, which should default to DEBUG.
}

func TestLog(t *testing.T) {
	tl := &testLogger{}
	AddLogger(tl)
	SetLogLevelString("DEBUG")

	Log(LogSeverityDebug, "TestFunc", 1, "test message")

	if tl.LastLog == "" {
		t.Errorf("Expected log message, got empty string")
	}
}
