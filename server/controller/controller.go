package controller

import (
	"akai.org.pl/joystick_server/engine"
	"akai.org.pl/joystick_server/game"
)

var eng engine.Engine

func init() {
	eng = engine.NewEngine(game.NewRoomManager())
}
