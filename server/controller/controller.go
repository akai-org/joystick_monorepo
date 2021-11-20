package controller

import (
	"akai.org.pl/joystick_server/engine"
	"akai.org.pl/joystick_server/game"
)

type controller struct {
	engine engine.Engine
}

func New() *controller {
	return &controller{
		engine: engine.NewEngine(game.NewRoomManager()),
	}
}
