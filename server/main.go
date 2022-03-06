package main

import (
	"akai.org.pl/joystick_server/controller"
	"akai.org.pl/joystick_server/logger"
	"os"
)

func main() {
	logFileWriter, err := os.Create("logs")
	if err != nil {
		panic(err)
	}
	logg := logger.New(logger.DebugLevel, logFileWriter)
	ctrl := controller.New(logg)
	ctrl.Listen()
}
