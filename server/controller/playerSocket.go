package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type initialWsPlayerMessage struct {
	GlobalID string `json:"global_id"`
}

func (c *controller) playerSocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := c.upgrader.Upgrade(w, r, nil)
	if err != nil {
		c.logger.Debug(fmt.Sprintf("Error during connection upgradation: %v", err))
		return
	}
	defer conn.Close()
	c.logger.Debug("Received correct player registration connection")
	var payload initialWsPlayerMessage
	messageType, message, err := conn.ReadMessage()
	if err != nil {
		c.logger.Debug(fmt.Sprintf("Error druing message reading: %v", err))
		return
	}
	if messageType != websocket.TextMessage {
		c.logger.Debug("First message is not text type")
		return
	}
	messageReader := bytes.NewReader(message)
	if err := json.NewDecoder(messageReader).Decode(&payload); err != nil {
		c.logger.Debug("First message is not valid json")
		return
	}
	player, err := c.engine.GetPlayer(payload.GlobalID)
	if err != nil {
		c.logger.Debug("Such player does not exist")
		return
	}
	c.logger.Debug(fmt.Sprintf("Player %s connected", player.Nickname))
	defer func() {
		c.logger.Debug(fmt.Sprintf("Player %s disconnected, removing that player", player.Nickname))
		c.engine.RemovePlayer(payload.GlobalID)
	}()

	connectionClose := make(chan interface{})
	go ping(connectionClose, conn)
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			c.logger.Warning(fmt.Sprintf("Error during message reading: %v", err))
			return
		}
		if messageType != websocket.BinaryMessage {
			c.logger.Debug(fmt.Sprintf("Message received from player %v is not of binary type!", player.Nickname))
			continue
		}
		player.SendMessageToRoom(message)
	}
}
