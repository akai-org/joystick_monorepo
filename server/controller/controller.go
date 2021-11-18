// The controller package provides RESTfull api for joystick engine
package controller

import (
	"akai.org.pl/joystick_server/engine"
	"akai.org.pl/joystick_server/room"
)

type Controller struct {
	engine engine.Engine
}

func New() *Controller {
	roomManager := room.New()
	return &Controller{
		engine: engine.NewEngine(roomManager),
	}
}
