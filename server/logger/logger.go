package logger

import (
	"log"
	"os"
	"io/ioutil"
	"io"
)

const loggerFlags = log.Ldate | log.Ltime


type Logger struct{
	debugLogger log.Logger
	infoLogger log.Logger
	warningLogger log.Logger
	errorLogger log.Logger
	fatalLogger log.Logger
}


func InitializeLogger(level uint8) *Logger {
	return &Logger{
        debugLogger:   *log.New(AssignWriter(1, level), "", loggerFlags),
        infoLogger:    *log.New(AssignWriter(2, level), "", loggerFlags),
        warningLogger: *log.New(AssignWriter(3, level), "", loggerFlags),
        errorLogger:   *log.New(AssignWriter(4, level), "", loggerFlags),
        fatalLogger:   *log.New(AssignWriter(5, level), "", loggerFlags),
    }
}

func AssignWriter(level uint8, givenLevel uint8) io.Writer{
	if givenLevel<level{
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
