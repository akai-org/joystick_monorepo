package controller

import (
	"encoding/json"
	"net/http"
)

type roomCreateRequest struct {
	Gui       string `json:"gui"`
	MaxPlayer int    `json:"max_players"`
}

type roomCreateResponse struct {
	Code    string `json:"code"`
}

const (
	roomCreatingRequestDataInvalid = "Given room creating data is invalid"
)

func (c *controller) createRoom(w http.ResponseWriter, r *http.Request) {
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

	var payload roomCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response := errorResponse{Message: messageIsNotCorrectJson}
		jsonResponse(w, response, http.StatusBadRequest)
		return
	}

	if payload.invalid() {
		response := errorResponse{Message: roomCreatingRequestDataInvalid}
		jsonResponse(w, response, http.StatusBadRequest)
		return
	}

	if code, err := c.engine.CreateNewRoom(payload.Gui, payload.MaxPlayer); err != nil {
		response := errorResponse{Message: err.Error()}
		jsonResponse(w, response, http.StatusBadRequest)
		return
	} else {
		response := roomCreateResponse{
			Code:    code,
		}
		jsonResponse(w, response, http.StatusOK)
		return
	}
}

func (r *roomCreateRequest) invalid() bool {
	return r.Gui == "" || r.MaxPlayer <= 0
}
