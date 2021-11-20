package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var counter uint8 = 0
var upgrader = websocket.Upgrader{} // use default options

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	id := counter
	counter++
	defer conn.Close()
	// The event loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Received: %s from conn %d", message, id)
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}
}

func prepareSocket() {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
}