package communication

import (
	"akai.org.pl/joystick_server/games"
	"encoding/json"
	"errors"
	"net/http"
)

const (
	methodNotAllowedMessage = "method not allowed, use POST instead"
	messageIsNotCorrectJson = "given message is not a valid json"
	playerNameTooLongMessage = "player nickname is too long1"
	maxAllowedNicknameLength = 16
	gameDoesntExistMessage = "room with such code doesn't exist"
)


type playerRequest struct {
	RoomCode	string	`json:"room_code"`
	Nickname	string	`json:"nickname"`
}

func RegisterNewPlayer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, methodNotAllowedMessage, http.StatusMethodNotAllowed)
		return
	}

	var payload playerRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, messageIsNotCorrectJson, http.StatusBadRequest)
		return
	}

	if err := validatePayload(payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	player := games.Player{
		Nickname: payload.Nickname,
	}

	if err := games.AddPlayerToRoom(player, payload.RoomCode); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
}

func validatePayload(payload playerRequest) error {
	if len(payload.Nickname) > maxAllowedNicknameLength {
		return errors.New(playerNameTooLongMessage)
	}
	if !games.RoomExists(payload.RoomCode) {
		return errors.New(gameDoesntExistMessage)
	}
	return nil
}
