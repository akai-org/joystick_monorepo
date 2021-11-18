package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"akai.org.pl/joystick_server/engine"
	"akai.org.pl/joystick_server/room"
)

func TestCreateRoomAndReturnCode(t *testing.T) {
	rm := room.RoomManager{}

	c := Controller{engine.NewEngine(rm)}

	reqRoomJson, err := json.Marshal(room.Room{GUI: "basic", MaxPlayers: 3})
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/room", bytes.NewBuffer(reqRoomJson))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.CreateRoom)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var actualNewRoomResponse room.Room
	err = json.Unmarshal(rr.Body.Bytes(), &actualNewRoomResponse)
	if err != nil {
		t.Fatal(err)
	}

	expectedNewRoom := room.Room{Code: 0, GUI: "basic", MaxPlayers: 3}

	if !reflect.DeepEqual(expectedNewRoom, actualNewRoomResponse) {
		t.Errorf("expected room %v, but got from response %v", expectedNewRoom, actualNewRoomResponse)
	}

	actualNewRoom, err := rm.GetRoom(actualNewRoomResponse.Code)
	if err != nil {
		t.Errorf("expected room in manager, but got nil")
	}

	if !reflect.DeepEqual(expectedNewRoom, actualNewRoom) {
		t.Errorf("expected room %v, but created one is %v", expectedNewRoom, actualNewRoom)
	}
}
