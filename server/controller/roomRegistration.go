package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type roomCreateRequest struct {
	Gui       string `json:"gui"`
	MaxPlayer int    `json:"max_players"`
}

type roomCreateResponse struct {
	Code string `json:"code"`
}

const (
	roomCreatingRequestDataInvalid = "Given room creating data is invalid"
)

var interfaceNames = []string{"arrows-horizontal", "arrows-vertical", "arrows-vertical-1ab", "cross-arrows", "cross-arrows-1ab"}

func (c *controller) createRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		response := new(struct{})
		jsonResponse(w, response, http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		response := errorResponse{Message: methodNotAllowedMessage}
		jsonResponse(w, response, http.StatusMethodNotAllowed)
		c.logger.Debug("Wrong method given at room creating endpoint")
		return
	}

	c.logger.Debug("Received request for creating new room")
	var payload roomCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response := errorResponse{Message: messageIsNotCorrectJson}
		jsonResponse(w, response, http.StatusBadRequest)
		c.logger.Warning("Room not created due to wrong request payload format")
		return
	}

	if payload.invalid() {
		response := errorResponse{Message: roomCreatingRequestDataInvalid}
		jsonResponse(w, response, http.StatusBadRequest)
		c.logger.Warning("Request for creating new room is invalid")
		return
	}
	c.logger.Debug("Request for creating new room is valid")

	if code, err := c.engine.CreateNewRoom(payload.Gui, payload.MaxPlayer); err != nil {
		response := errorResponse{Message: err.Error()}
		jsonResponse(w, response, http.StatusBadRequest)
		c.logger.Warning("Failed to create new room")
		return
	} else {
		response := roomCreateResponse{
			Code: code,
		}
		jsonResponse(w, response, http.StatusOK)
		c.logger.Info(fmt.Sprintf("Created new room with code: {%v}", code))
		return
	}
}

func (r *roomCreateRequest) invalid() bool {
	for _, name:= range interfaceNames{
		if name == r.Gui && (r.MaxPlayer>0 && r.MaxPlayer<=255){
			return true
		}
	}
	return false
}
