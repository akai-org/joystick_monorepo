package controller

import (
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
	var m sync.Mutex
	if err != nil {
		c.logger.Warning(fmt.Sprintf("Failed to upgrade connection: %v", err))
		return
	}
	c.conn = conn
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
	go ping(connectionClose, conn, &m)
	for {
		select {
		case playerInput := <-room.PlayerChannel:
			go func() {
				c.logger.Debug(fmt.Sprintf("From player %d received %d\n", playerInput[0], playerInput[1]))
				messageType = websocket.BinaryMessage
				err := WriteMessage(conn, &m, messageType, playerInput)
				if err != nil {
					c.logger.Warning(fmt.Sprintf("Error during message writing: %v", err))
				}
			}()
		case serverNotification := <-room.CommunicationChannel:
			c.logger.Debug(fmt.Sprintf("Sending system communication to game host: %v", serverNotification))
			err := WriteMessage(conn, &m, messageType, []byte(serverNotification))
			messageType = websocket.TextMessage
			if err != nil {
				c.logger.Warning(fmt.Sprintf("Error during notification sending: %v", err))
			}
		case <-connectionClose:
			return
		}
	}
}

func ping(connectionClose chan interface{}, conn *websocket.Conn, mutex *sync.Mutex) {
	tickerTime := time.Second
	ticker := time.NewTicker(tickerTime)
	defer func() {
		ticker.Stop()
		connectionClose <- true
	}()
	for {
		select {
		case <-ticker.C:
			mutex.Lock()
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				mutex.Unlock()
				return
			}
			mutex.Unlock()
		}
	}
}

func WriteMessage(conn *websocket.Conn, mutex *sync.Mutex, messageType int, data []byte) error {
	mutex.Lock()
	defer mutex.Unlock()
	return conn.WriteMessage(messageType, data)
}
