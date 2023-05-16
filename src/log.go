package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

type LogSeverity int

type SCLogger interface {
	Init() error
	Term()
	PrintLog(value string)
}

var gLogLevel LogSeverity = LogSeverityInfo

const (
	LogSeverityDebug LogSeverity = iota
	LogSeverityStats
	LogSeverityInfo
	LogSeverityWarn
	LogSeverityError
	LogSeverityFatal
)

var logSeverityNames = [...]string{
	"debug",
	"stats",
	"info",
	"warn",
	"error",
	"fatal",
}

const (
	LoggerTimeFormat     = "2006-01-02T15:04:05.000"
	LoggerFileTimeFormat = "06_01_02__15_04_05"
)

// Default Logger

var logServerHandler []SCLogger

func (l LogSeverity) String() string {

	return logSeverityNames[int(l)]
}

func AddLogger(l SCLogger) {

	if err := l.Init(); err != nil {
		LogWarn("function=AddLogger, error=%v", err)
		return
	}

	logServerHandler = append(logServerHandler, l)
}

func SetLogLevelString(level string) {

	switch level {
	case "DEBUG":
		gLogLevel = LogSeverityDebug
	case "STATS":
		gLogLevel = LogSeverityStats
	case "INFO":
		gLogLevel = LogSeverityInfo
	case "WARN":
		LogInfo("function=SetLogLevelConfig, message=log level set to %v", level)
		gLogLevel = LogSeverityWarn
	case "ERROR":
		LogInfo("function=SetLogLevelConfig, message=log level set to %v", level)
		gLogLevel = LogSeverityError
	case "FATAL":
		LogInfo("function=SetLogLevelConfig, message=log level set to %v", level)
		gLogLevel = LogSeverityFatal
	default:
		LogWarn("function=SetLogLevelConfig, logLevel=%v message=invalid log level, log level set to default : DEBUG", level)
		gLogLevel = LogSeverityDebug
		return
	}

	LogInfo("function=SetLogLevelConfig, message=log level set to %v", level)
}

func SetLogLevel(level LogSeverity) {
	gLogLevel = level
}

func TerminateLog() {

	for _, handler := range logServerHandler {
		handler.Term()
	}

	// clear logs
	logServerHandler = []SCLogger{}
}

func Log(severity LogSeverity, fn string, line int, format string, a ...interface{}) {
	if severity < gLogLevel {
		return
	}

	b := strings.Builder{}
	b.WriteString(time.Now().UTC().Format(LoggerTimeFormat))
	b.WriteString(" : ")
	b.WriteString(fmt.Sprintf("%s:%v", fn, line))
	b.WriteString(" : ")
	b.WriteString(severity.String())
	b.WriteString(" : ")

	b.WriteString(fmt.Sprintf(format, a...))

	str := b.String()

	// Temp
	fmt.Println(str)
	for _, handler := range logServerHandler {
		handler.PrintLog(str)
	}
}

func LogInfo(format string, a ...interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	Log(LogSeverityInfo, fn, line, format, a...)
}

func LogWarn(format string, a ...interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	Log(LogSeverityWarn, fn, line, format, a...)
}

func LogError(format string, a ...interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	Log(LogSeverityError, fn, line, format, a...)
}

func LogDebug(format string, a ...interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	Log(LogSeverityDebug, fn, line, format, a...)
}

func LogFatal(format string, a ...interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	Log(LogSeverityFatal, fn, line, format, a...)
}
