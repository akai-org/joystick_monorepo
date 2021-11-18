package logger

import (
	"flag"
	"log"
	"os"
)

var (
	isVerbose bool
	logger *log.Logger
)

func InitializeLogger() {
	verbose := flag.Bool("verbose", false, "Enables verbose logger")
	flag.Parse()
	if *verbose {
		isVerbose = true
	} else {
		isVerbose = false
	}
}

func Debug(message string) {
	if isVerbose{
		logger = log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds)
		logger.Print(message)
	}
}

func Info(message string) {
	if isVerbose{
		logger = log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds)
		logger.Print(message)
	}
}

func Warning(message string) {
	if isVerbose{
		logger = log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds)
		logger.Print(message)
	}
}

func Error(message string) {
	logger = log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds)
	logger.Print(message)
}

func Fatal(message string) {
	logger = log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds)
	logger.Print(message)
}
