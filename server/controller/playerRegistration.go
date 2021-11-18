package controller

import (
	"akai.org.pl/joystick_server/games"
	"encoding/json"
	"errors"
	"net/http"
	"unicode"
)

const (
	methodNotAllowedMessage    = "method not allowed, use POST instead"
	messageIsNotCorrectJson    = "given message is not a valid json"
	playerNameTooLongMessage   = "player nickname is too long"
	maxAllowedNicknameLength   = 16
	forbiddenCharsInPlayerName = "nickname contains forbidden characters, use alphanumeric ones and _"
)

type playerRequest struct {
	RoomCode string `json:"room_code"`
	Nickname string `json:"nickname"`
}

type playerResponse struct {
	Address   string `json:"address"`
	Interface string `json:"interface"`
}

func (c *Controller) RegisterNewPlayer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		response := new(struct{})
		jsonResponse(w, response, http.StatusNoContent)
		return
	}

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

	player := &games.Player{
		Nickname: payload.Nickname,
	}

	gameRoom, err := c.engine.RoomManager.GetRoom(payload.RoomCode)
	if err != nil {
		response := errorResponse{Message: err.Error()}
		jsonResponse(w, response, http.StatusForbidden)
		return
	}

	if err := gameRoom.AddPlayer(player); err != nil {
		response := errorResponse{Message: err.Error()}
		jsonResponse(w, response, http.StatusForbidden)
		return
	}

	response := playerResponse{
		Address:   "Somethingthatwillbereplaced",
		Interface: gameRoom.Gui,
	}

	jsonResponse(w, response, http.StatusOK)
	return
}

func (payload *playerRequest) isValid() error {
	if len(payload.Nickname) > maxAllowedNicknameLength {
		return errors.New(playerNameTooLongMessage)
	}
	for _, char := range payload.Nickname {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '_' {
			return errors.New(forbiddenCharsInPlayerName)
		}
	}
	return nil
}
