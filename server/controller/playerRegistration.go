package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"unicode"
)

const (
	methodNotAllowedMessage    = "method not allowed, use POST instead"
	messageIsNotCorrectJson    = "given message is not a valid json"
	playerNameTooLongMessage   = "player nickname is too long"
	playerNameTooShortMessage  = "player nickname is too short"
	maxAllowedNicknameLength   = 16
	minAllowedNicknameLength   = 2
	forbiddenCharsInPlayerName = "nickname contains forbidden characters, use alphanumeric ones and _"
)

type playerRequest struct {
	RoomCode string `json:"room_code"`
	Nickname string `json:"nickname"`
}

type playerResponse struct {
	GlobalId string `json:"global_id"`
	Gui      string `json:"gui"`
}

func (c *controller) registerNewPlayer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		response := new(struct{})
		jsonResponse(w, response, http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		response := errorResponse{Message: methodNotAllowedMessage}
		jsonResponse(w, response, http.StatusMethodNotAllowed)
		c.logger.Debug("Client tried to register player using wrong method")
		return
	}

	var payload playerRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response := errorResponse{Message: messageIsNotCorrectJson}
		jsonResponse(w, response, http.StatusBadRequest)
		c.logger.Debug("Client tried to register player with wrong payload format")
		return
	}

	if err := payload.isValid(); err != nil {
		response := errorResponse{Message: err.Error()}
		jsonResponse(w, response, http.StatusBadRequest)
		c.logger.Debug("Client tried to register player provided invalid payload")
		return
	}

	gameRoom, err := c.engine.GetRoom(payload.RoomCode)
	if err != nil {
		response := errorResponse{Message: err.Error()}
		jsonResponse(w, response, http.StatusForbidden)
		return
	}

	player, code := c.engine.CreateNewPlayer(payload.Nickname)

	if err := gameRoom.AddPlayer(player); err != nil {
		response := errorResponse{Message: err.Error()}
		jsonResponse(w, response, http.StatusForbidden)
		return
	}

	c.logger.Info(fmt.Sprintf("Player %v has been registered to room %v", payload.Nickname, payload.RoomCode))

	response := playerResponse{
		GlobalId: code,
		Gui:      gameRoom.Gui,
	}

	jsonResponse(w, response, http.StatusOK)
}

func (payload *playerRequest) isValid() error {
	if len(payload.Nickname) > maxAllowedNicknameLength {
		return errors.New(playerNameTooLongMessage)
	}
	if len(payload.Nickname) < minAllowedNicknameLength {
		return errors.New(playerNameTooShortMessage)
	}
	for _, char := range payload.Nickname {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '_' {
			return errors.New(forbiddenCharsInPlayerName)
		}
	}
	return nil
}
