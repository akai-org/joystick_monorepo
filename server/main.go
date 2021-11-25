package main

import (
	"akai.org.pl/joystick_server/controller"
)

func main() {
	ctrl := controller.New()
	ctrl.Listen()
}
