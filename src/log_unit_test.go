package main

import (
	"errors"
	"strings"
	"testing"
)

type TestLogger struct {
	messages []string
}

func (t *TestLogger) Init() error {
	t.messages = []string{}
	return nil
}

func (t *TestLogger) Term() {
	t.messages = []string{}
}

func (t *TestLogger) PrintLog(value string) {
	t.messages = append(t.messages, value)
}

func TestLog(t *testing.T) {
	tLogger := &TestLogger{}
	AddLogger(tLogger)

	SetLogLevel(LogSeverityDebug)

	LogDebug("This is a debug log")
	if !strings.Contains(tLogger.messages[len(tLogger.messages)-1], "This is a debug log") {
		t.Error(errors.New("Debug log message not found"))
	}

	LogInfo("This is an info log")
	if !strings.Contains(tLogger.messages[len(tLogger.messages)-1], "This is an info log") {
		t.Error(errors.New("Info log message not found"))
	}

	LogWarn("This is a warning log")
	if !strings.Contains(tLogger.messages[len(tLogger.messages)-1], "This is a warning log") {
		t.Error(errors.New("Warning log message not found"))
	}

	LogError("This is an error log")
	if !strings.Contains(tLogger.messages[len(tLogger.messages)-1], "This is an error log") {
		t.Error(errors.New("Error log message not found"))
	}

	LogFatal("This is a fatal log")
	if !strings.Contains(tLogger.messages[len(tLogger.messages)-1], "This is a fatal log") {
		t.Error(errors.New("Fatal log message not found"))
	}

	SetLogLevelString("WARN")
	if gLogLevel != LogSeverityWarn {
		t.Error(errors.New("Log level not set to WARN"))
	}

	SetLogLevelString("INVALID")
	if gLogLevel != LogSeverityDebug {
		t.Error(errors.New("Invalid log level should set level to DEBUG"))
	}
}
