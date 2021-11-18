package restapi

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
	gameDoesntExistMessage     = "room with such code doesn't exist"
	forbiddenCharsInPlayerName = "nickname contains forbidden characters, use alphanumeric ones and _"
)

func RegisterNewPlayer(w http.ResponseWriter, r *http.Request) {
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

	player := games.Player{
		Nickname: payload.Nickname,
	}

	if err := games.AddPlayerToRoom(player, payload.RoomCode); err != nil {
		response := errorResponse{Message: err.Error()}
		jsonResponse(w, response, http.StatusForbidden)
		return
	}

	response := playerResponse{
		Address:   "Somethingthatwillbereplaced",
		Interface: games.GetRoomInterface(payload.RoomCode),
	}

	jsonResponse(w, response, http.StatusOK)
	return
}

func (payload *playerRequest) isValid() error {
	if err := validateNickname(payload.Nickname); err != nil {
		return err
	}
	if !games.RoomExists(payload.RoomCode) {
		return errors.New(gameDoesntExistMessage)
	}
	return nil
}

func validateNickname(nickname string) error {
	if len(nickname) > maxAllowedNicknameLength {
		return errors.New(playerNameTooLongMessage)
	}
	for _, char := range nickname {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '_' {
			return errors.New(forbiddenCharsInPlayerName)
		}
	}
	return nil
}
