package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type initialWsRoomMessage struct {
	Code string `json:"code"`
}

func (c *controller) roomSocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := c.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	var payload initialWsRoomMessage
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
	fmt.Println(payload.Code)
	room, err := c.engine.GetRoom(payload.Code)
	if err != nil {
		log.Print("Such room does not exist")
		return
	}
	// jeżeli pokój jest już połączony to wypierdol eerrror
	log.Printf("Game host connected: %s", payload.Code)

	for {
		var playerInput, serverInput, message []byte
		var messageType int
		select {
			case serverInput = <- room.ServerMessageChannel:
				log.Printf("From server received message: %s\n", serverInput)
				messageType = websocket.TextMessage
				message = serverInput

			case playerInput = <- room.PlayerChannel:
				log.Printf("From player %d received %d\n", playerInput[0], playerInput[1])
				messageType = websocket.BinaryMessage
				message = playerInput
		}
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}
}
