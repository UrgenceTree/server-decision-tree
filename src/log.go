package main

import (
	"flag"
	"fmt"
	"runtime"
	"strings"
	"time"
)

type LogSeverity int

type Logger interface {
	Init() error
	Term()
	PrintLog(value string)
}

var gNoTimestamp bool = false
var gNoFilePath bool = false
var gNoPartielFilePath bool = true
var gNoColor bool = false
var gLogLevel LogSeverity = LogSeverityInfo

const (
	LogSeverityDebug LogSeverity = iota
	LogSeverityStats
	LogSeverityInfo
	LogSeveritySuccess
	LogSeverityWarn
	LogSeverityError
	LogSeverityFatal
)

var logSeverityNames = [...]string{
	"[ debug ]",
	"[ stats ]",
	"[ info  ]",
	"[ success ]",
	"[ warn  ]",
	"[ error ]",
	"[ fatal ]",
}

const (
	LoggerTimeFormat     = "2006-01-02T15:04:05.000"
	LoggerFileTimeFormat = "06_01_02__15_04_05"
)

// Default Logger

var logHandler []Logger

func (l LogSeverity) String() string {

	return logSeverityNames[int(l)]
}

func AddLogger(l Logger) {

	if err := l.Init(); err != nil {
		LogWarn("function=AddLogger, error=%v", err)
		return
	}

	logHandler = append(logHandler, l)
}

func SetLogLevelString(level string) {

	switch level {
	case "DEBUG":
		gLogLevel = LogSeverityDebug
	case "STATS":
		gLogLevel = LogSeverityStats
	case "INFO":
		gLogLevel = LogSeverityInfo
	case "SUCCESS":
		gLogLevel = LogSeveritySuccess
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

	for _, handler := range logHandler {
		handler.Term()
	}

	// clear logs
	logHandler = []Logger{}
}

func Log(severity LogSeverity, fn string, line int, format string, a ...interface{}) {

	if severity < gLogLevel {
		return
	}

	b := strings.Builder{}

	if !gNoTimestamp {
		b.WriteString(time.Now().UTC().Format(LoggerTimeFormat))
		b.WriteString(" : ")
	}
	if !gNoPartielFilePath {
		b.WriteString(fmt.Sprintf("%s:%v", fn[strings.LastIndex(fn, "/")+1:], line))
		b.WriteString(" : ")
	} else if !gNoFilePath {
		b.WriteString(fmt.Sprintf("%s:%v", fn, line))
		b.WriteString(" : ")
	}
	if !gNoColor {
		switch severity {
		case LogSeverityDebug:
			b.WriteString("\x1b[0m" + severity.String() + "\x1b[0m") // white
		case LogSeverityStats:
			b.WriteString("\x1b[0m" + severity.String() + "\x1b[0m") // white
		case LogSeverityInfo:
			b.WriteString("\x1b[0m" + severity.String() + "\x1b[0m") // white
		case LogSeveritySuccess:
			b.WriteString("\x1b[32m" + severity.String() + "\x1b[0m") // green
		case LogSeverityWarn:
			b.WriteString("\x1b[33m" + severity.String() + "\x1b[0m") // yellow
		case LogSeverityError:
			b.WriteString("\x1b[31m" + severity.String() + "\x1b[0m") // red
		case LogSeverityFatal:
			b.WriteString(severity.String()) // special case
		}
	}
	if gNoColor {
		println("no color")
		b.WriteString(severity.String())
	}

	b.WriteString(" : ")

	b.WriteString(fmt.Sprintf(format, a...))

	str := b.String()

	if severity == LogSeverityFatal && !gNoColor {
		str = "\x1b[1;31m" + str + "\x1b[0m" // red
	}

	fmt.Println(str)
	for _, handler := range logHandler {
		handler.PrintLog(str)
	}
}

func LogInfo(format string, a ...interface{}) {

	_, fn, line, _ := runtime.Caller(1)
	Log(LogSeverityInfo, fn, line, format, a...)
}

func LogSuccess(format string, a ...interface{}) {

	_, fn, line, _ := runtime.Caller(1)
	Log(LogSeveritySuccess, fn, line, format, a...)
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

func LogStats(format string, a ...interface{}) {

	_, fn, line, _ := runtime.Caller(1)
	Log(LogSeverityStats, fn, line, format, a...)
}

func LogFatal(format string, a ...interface{}) {

	_, fn, line, _ := runtime.Caller(1)
	Log(LogSeverityFatal, fn, line, format, a...)
}

func InitLogger() {

	// Define the logger flag with default values
	loggerFlag := flag.Bool("logger", false, "Enable custom logger settings")
	noFilePath := flag.Bool("no_file_path", false, "Disable file path logging")
	noPartialFilepath := flag.Bool("partial_file_path", false, "Disable partial file path logging")
	noTimestamp := flag.Bool("no_timestamp", false, "Disable timestamp logging")
	noColor := flag.Bool("no_color", false, "Disable color logging")
	logLevel := flag.String("log_level", "INFO", "Set log level")

	// Parse the flags
	flag.Parse()

	// If the logger flag was provided, check for additional flags
	if *loggerFlag {
		LogDebug("function=main, message=Logger flag provided")

		gNoFilePath = *noFilePath
		gNoColor = *noColor
		gNoTimestamp = *noTimestamp
		gNoPartielFilePath = *noPartialFilepath

		println("no partial file path", gNoPartielFilePath)

		SetLogLevelString(*logLevel)
	}
}
