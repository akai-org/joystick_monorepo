package controller

import (
	"akai.org.pl/joystick_server/engine"
	"akai.org.pl/joystick_server/game"
	"github.com/gorilla/websocket"
	"net/http"
)

type controller struct {
	engine engine.Engine
	upgrader websocket.Upgrader
}

func New() *controller {
	upgrader := websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	return &controller{
		engine: engine.NewEngine(game.NewRoomManager()),
		upgrader: upgrader,
	}
}
