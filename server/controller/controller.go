package controller

import (
	"akai.org.pl/joystick_server/engine"
	"akai.org.pl/joystick_server/game"
	"akai.org.pl/joystick_server/logger"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type controller struct {
	engine   engine.Engine
	upgrader websocket.Upgrader
	logger   *logger.Logger
	conn     *websocket.Conn
}

func New(logger *logger.Logger) *controller {
	logger.Debug("Creating controller")
	upgrader := websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	return &controller{
		engine:   engine.NewEngine(game.NewRoomManager(logger), game.NewPlayerManager(logger)),
		upgrader: upgrader,
		logger:   logger,
	}
}

func (c *controller) Listen() {
	c.logger.Info("Began listening...")
	http.HandleFunc("/player/socket", c.playerSocketHandler)
	http.HandleFunc("/room/socket", c.roomSocketHandler)
	http.HandleFunc("/join", c.registerNewPlayer)
	http.HandleFunc("/create", c.createRoom)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
