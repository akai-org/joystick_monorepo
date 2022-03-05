package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
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

	c.logger.Debug(fmt.Sprintf("Listening for player input on channel %v", room.PlayerChannel))
	playerInput := <-room.PlayerChannel
	c.logger.Debug(fmt.Sprintf("From player %d received %d\n", playerInput[0], playerInput[1]))
	messageType = websocket.BinaryMessage

	err = conn.WriteMessage(messageType, playerInput)
	if err != nil {
		c.logger.Warning(fmt.Sprintf("Error during message writing: %v", err))
	}
	c.logger.Info(fmt.Sprintf("Connection to game host %v broken", payload.Code))
}
