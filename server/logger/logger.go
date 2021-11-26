package logger

import (
	"log"
	"os"
	"io/ioutil"
	"io"
)


const loggerFlags = log.Ldate | log.Ltime

const (
    DebugLevel uint8 = iota+1
    InfoLevel
    WarningLevel
    ErrorLevel
    FatalLevel
)


type Logger struct{
	debugLogger log.Logger
	infoLogger log.Logger
	warningLogger log.Logger
	errorLogger log.Logger
	fatalLogger log.Logger
}


func New(level uint8) *Logger {
	return &Logger{
        debugLogger:   *log.New(assignWriter(DebugLevel, level), "", loggerFlags),
        infoLogger:    *log.New(assignWriter(InfoLevel, level), "", loggerFlags),
        warningLogger: *log.New(assignWriter(WarningLevel, level), "", loggerFlags),
        errorLogger:   *log.New(assignWriter(ErrorLevel, level), "", loggerFlags),
        fatalLogger:   *log.New(assignWriter(FatalLevel, level), "", loggerFlags),
    }
}


func assignWriter(level uint8, givenLevel uint8) io.Writer{
	if givenLevel>level{
		return ioutil.Discard
	}
	return os.Stderr
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
