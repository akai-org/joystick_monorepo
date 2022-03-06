package controller

import (
	"akai.org.pl/joystick_server/game"
	"akai.org.pl/joystick_server/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

type initialWsRoomMessage struct {
	Code string `json:"code"`
}

func (c *controller) roomSocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := c.upgrader.Upgrade(w, r, nil)
	if err != nil {
		c.logger.Warning(fmt.Sprintf("Failed to upgrade connection: %v", err))
		return
	}
	defer conn.Close()

	c.logger.Debug("Successfuly established websocket connection with room")
	var payload initialWsRoomMessage
	messageType, message, err := conn.ReadMessage()
	if err != nil {
		c.logger.Warning(fmt.Sprintf("Error druing message reading: %v", err))
		return
	}
	if messageType != websocket.TextMessage {
		c.logger.Warning("First message is not text type")
		return
	}
	messageReader := bytes.NewReader(message)
	if err := json.NewDecoder(messageReader).Decode(&payload); err != nil {
		c.logger.Warning("First message is not valid json")
		return
	}
	fmt.Println(payload.Code)
	room, err := c.engine.GetRoom(payload.Code)
	if err != nil {
		c.logger.Warning("Such room does not exist")
		return
	}
	// jeżeli pokój jest już połączony to wypierdol eerrror
	c.logger.Info(fmt.Sprintf("Game host connected: %s", payload.Code))
	defer func() {
		c.logger.Debug(fmt.Sprintf("Communication loop with game %v host has been broken. Removing that room now.", payload.Code))
		c.engine.RemoveRoom(payload.Code)
	}()

	c.logger.Debug(fmt.Sprintf("Listening for player input on channel %v", room.PlayerChannel))
	connectionClose := make(chan interface{})
	wg := sync.WaitGroup{}
	wg.Add(2)
	go communicateWithHost(room, conn, c.logger, connectionClose, &wg)
	go ping(connectionClose, conn, &wg, c.logger)
	wg.Wait()
}

func ping(connectionClose chan interface{}, conn *websocket.Conn, wg *sync.WaitGroup, logger *logger.Logger) {
	tickerTime := time.Second
	ticker := time.NewTicker(tickerTime)
	defer func() {
		ticker.Stop()
		connectionClose <- true
		wg.Done()
	}()
	for {
		select {
		case <-ticker.C:
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func communicateWithHost(room *game.Room, conn *websocket.Conn, logger *logger.Logger, gameHostClosedConnection <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	messageType := websocket.BinaryMessage
	for {
		select {
		case playerInput := <-room.PlayerChannel:
			logger.Debug(fmt.Sprintf("From player %d received %d\n", playerInput[0], playerInput[1]))

			err := conn.WriteMessage(messageType, playerInput)
			if err != nil {
				logger.Warning(fmt.Sprintf("Error during message writing: %v", err))
			}
		case <-gameHostClosedConnection:
			return
		}
	}
}
