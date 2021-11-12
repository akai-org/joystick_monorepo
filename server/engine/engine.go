// Package engine provides tools to run and manage joystick by combining all the implemented features.
package engine

import (
	"akai.org.pl/joystick_server/room"
)

type Engine struct {
	roomManager room.RoomManager
	//TODO playerManager player.Manager, etc.
}

func NewEngine(rm room.RoomManager) Engine {
	return Engine{rm}
}

func (e *Engine) CreateRoom(gui string, maxPlayers int) (code int) {
	return e.roomManager.CreateRoom(gui, maxPlayers)
}
