// Package engine provides tools to run and manage joystick by combining all the implemented features.
package engine

import (
	"akai.org.pl/joystick_server/game"
)

type Engine struct {
	*game.Manager
}

func NewEngine(rm *game.Manager) Engine {
	return Engine{rm}
}
