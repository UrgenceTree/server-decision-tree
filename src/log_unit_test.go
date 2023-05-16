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

	TerminateLog()

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

func TestAddLogger(t *testing.T) {

	TerminateLog()
	tLogger := &TestLogger{}

	AddLogger(tLogger)
	if len(logServerHandler) != 1 {
		t.Errorf("First logger was not added correctly : %v", len(logServerHandler))
	}

	AddLogger(tLogger)
	if len(logServerHandler) != 2 {
		t.Errorf("Second logger was not added correctly : %v", len(logServerHandler))
	}
}

func TestTerminateLog(t *testing.T) {

	TerminateLog()

	tLogger := &TestLogger{}

	AddLogger(tLogger)
	AddLogger(tLogger)

	// TerminateLog()
	// if len(logServerHandler) != 0 {
	// 	t.Error(errors.New("Loggers were not terminated correctly"))
	// }
}

func TestLog2(t *testing.T) {

	TerminateLog()

	tLogger := &TestLogger{}
	AddLogger(tLogger)

	SetLogLevel(LogSeverityDebug)

	// Test log with lower severity
	Log(LogSeverityDebug-1, "function", 1, "This log should not appear")
	if len(tLogger.messages) != 0 {
		t.Error(errors.New("Logger incorrectly logged message with lower severity"))
	}

	// Test log with equal severity
	Log(LogSeverityDebug, "function", 1, "This log should appear")
	if !strings.Contains(tLogger.messages[len(tLogger.messages)-1], "This log should appear") {
		t.Error(errors.New("Logger did not log message with equal severity"))
	}

	// Test log with higher severity
	Log(LogSeverityDebug+1, "function", 1, "This log should also appear")
	if !strings.Contains(tLogger.messages[len(tLogger.messages)-1], "This log should also appear") {
		t.Error(errors.New("Logger did not log message with higher severity"))
	}
}
