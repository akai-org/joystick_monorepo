// The controller package provides RESTfull api for joystick engine
package controller

import (
	"encoding/json"
	"net/http"

	"akai.org.pl/joystick_server/engine"
	"akai.org.pl/joystick_server/room"
)

type Controller struct {
	engine engine.Engine
}

func (c *Controller) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var newRoom room.Room

	err := json.NewDecoder(r.Body).Decode(&newRoom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newRoom.Code = c.engine.CreateRoom(newRoom.GUI, newRoom.MaxPlayers)

	newRoomJson, err := json.Marshal(newRoom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(newRoomJson)
}
