package controller

import (
	"akai.org.pl/joystick_server/game"
	"encoding/json"
	"net/http"
)

type roomCreateRequest struct {
	Gui       string `json:"interface"`
	MaxPlayer int    `json:"max_players"`
}

type roomCreateResponse struct {
	Address string `json:"address"`
	Code    string `json:"code"`
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	var payload roomCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response := errorResponse{Message: messageIsNotCorrectJson}
		jsonResponse(w, response, http.StatusBadRequest)
		return
	}

	newRoom := game.NewRoom(payload.Gui, payload.MaxPlayer)

	code := eng.GenerateCode()
	if err := eng.AppendRoomWithCode(newRoom, code); err != nil {
		response := errorResponse{Message: err.Error()}
		jsonResponse(w, response, http.StatusBadRequest)
		return
	}

	response := roomCreateResponse{
		Address: "asdasdadadsa",
		Code: code,
 	}

	jsonResponse(w, response, http.StatusOK)
	return
}
