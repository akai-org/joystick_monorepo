package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type initialWsPlayerMessage struct {
	GlobalID string `json:"global_id"`
}

func (c *controller) playerSocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := c.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	var payload initialWsPlayerMessage
	messageType, message, err := conn.ReadMessage()
	if err != nil {
		log.Print("Error druing message reading:", err)
		return
	}
	if messageType != websocket.TextMessage {
		log.Print("First message is not text type")
		return
	}
	messageReader := bytes.NewReader(message)
	if err := json.NewDecoder(messageReader).Decode(&payload); err != nil {
		log.Print("First message is not valid json")
		return
	}
	player, err := c.engine.GetPlayer(payload.GlobalID)
	if err != nil {
		log.Print("Such player does not exist")
		return
	}
	log.Printf("Player %s connected", player.Nickname)

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		if messageType != websocket.BinaryMessage {
			log.Println("Message received is not of binary type! Disconnecting")
			break
		}
		player.SendMessageToRoom(message)
		log.Printf("%s sent %s", player.Nickname, message)
	}
}
