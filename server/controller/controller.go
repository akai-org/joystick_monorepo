package controller

import (
	"akai.org.pl/joystick_server/engine"
	"akai.org.pl/joystick_server/game"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
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
		engine: engine.NewEngine(game.NewRoomManager(), game.NewPlayerManager()),
		upgrader: upgrader,
	}
}

func (c *controller) Listen() {
	fmt.Println("Began listening...")
	http.HandleFunc("/player/socket", c.playerSocketHandler)
	http.HandleFunc("/join", c.registerNewPlayer)
	http.HandleFunc("/create", c.createRoom)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
