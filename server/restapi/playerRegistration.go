package restapi

import (
	"akai.org.pl/joystick_server/games"
	"encoding/json"
	"errors"
	"net/http"
)

const (
	methodNotAllowedMessage  = "method not allowed, use POST instead"
	messageIsNotCorrectJson  = "given message is not a valid json"
	playerNameTooLongMessage = "player nickname is too long1"
	maxAllowedNicknameLength = 16
	gameDoesntExistMessage   = "room with such code doesn't exist"
)

func RegisterNewPlayer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := errorResponse{Message: methodNotAllowedMessage}
		jsonResponse(w, response, http.StatusMethodNotAllowed)
		return
	}

	var payload playerRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response := errorResponse{Message: messageIsNotCorrectJson}
		jsonResponse(w, response, http.StatusBadRequest)
		return
	}

	if err := payload.isValid(); err != nil {
		response := errorResponse{Message: err.Error()}
		jsonResponse(w, response, http.StatusBadRequest)
		return
	}

	player := games.Player{
		Nickname: payload.Nickname,
	}

	if err := games.AddPlayerToRoom(player, payload.RoomCode); err != nil {
		response := errorResponse{Message: err.Error()}
		jsonResponse(w, response, http.StatusForbidden)
		return
	}

	response := playerResponse{
		Address: "Somethingthatwillbereplaced",
		Interface: games.GetRoomInterface(payload.RoomCode),
	}

	jsonResponse(w, response, http.StatusOK)
	return
}

func (payload *playerRequest) isValid() error {
	if len(payload.Nickname) > maxAllowedNicknameLength {
		return errors.New(playerNameTooLongMessage)
	}
	if !games.RoomExists(payload.RoomCode) {
		return errors.New(gameDoesntExistMessage)
	}
	return nil
}
