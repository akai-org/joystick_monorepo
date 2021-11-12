// Package engine provides tools to run and manage joystick by combining all the implemented features.
package engine

import (
	"akai.org.pl/joystick_server/room"
)

type Engine struct {
	RoomManager room.RoomManager
	//TODO playerManager player.Manager, etc.
}
