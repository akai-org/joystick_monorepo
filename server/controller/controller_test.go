package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"akai.org.pl/joystick_server/engine"
	"akai.org.pl/joystick_server/room"
)

func TestCreateRoomAndReturnCode(t *testing.T) {
	rm := room.RoomManager{}

	c := Controller{engine.Engine{RoomManager: rm}}

	req, err := http.NewRequest("POST", "/api/room", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandleFunc(c.CreateRoom)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var actualNewRoom room.Room
	err = json.Unmarshal([]byte(rr.Body.String()), &actualNewRoom)
	if err != nil {
		t.Fatal(err)
	}

	expectedNewRoom := room.Room{Code: 0, GUI: "basic", MaxPlayers: 3}

	if actualNewRoom != expectedNewRoom {
		t.Errorf("expected room %v, but created one is %v", expectedNewRoom, actualNewRoom)
	}
}
