package logger

import (
	"io"
	"io/ioutil"
	"log"
)

const loggerFlags = log.Ldate | log.Ltime | log.Lmsgprefix

const (
	DebugLevel uint8 = iota + 1
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

type Logger struct {
	debugLogger   log.Logger
	infoLogger    log.Logger
	warningLogger log.Logger
	errorLogger   log.Logger
	fatalLogger   log.Logger
}

func New(level uint8, writer io.Writer) *Logger {
	return &Logger{
		debugLogger:   *log.New(assignWriter(DebugLevel, level, writer), "DEBUG | ", loggerFlags),
		infoLogger:    *log.New(assignWriter(InfoLevel, level, writer), "INFO  | ", loggerFlags),
		warningLogger: *log.New(assignWriter(WarningLevel, level, writer), "WARN  | ", loggerFlags),
		errorLogger:   *log.New(assignWriter(ErrorLevel, level, writer), "ERROR | ", loggerFlags),
		fatalLogger:   *log.New(assignWriter(FatalLevel, level, writer), "PANIC | ", loggerFlags),
	}
}

func assignWriter(level uint8, givenLevel uint8, writer io.Writer) io.Writer {
	if givenLevel > level {
		return ioutil.Discard
	}
	return writer
}

func (logger *Logger) Debug(message string) {
	logger.debugLogger.Print(message)
}

func (logger *Logger) Info(message string) {
	logger.infoLogger.Print(message)
}

func (logger *Logger) Warning(message string) {
	logger.warningLogger.Print(message)
}

func (logger *Logger) Error(message string) {
	logger.errorLogger.Print(message)
}

func (logger *Logger) Fatal(message string) {
	logger.fatalLogger.Print(message)
}
