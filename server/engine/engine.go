// Package engine provides tools to run and manage joystick by combining all the implemented features.
package engine

import (
	"akai.org.pl/joystick_server/game"
)

type Engine struct {
	*game.RoomManager
	*game.PlayerManager
}

func NewEngine(rm *game.RoomManager, pm *game.PlayerManager) Engine {
	return Engine{rm, pm}
}
