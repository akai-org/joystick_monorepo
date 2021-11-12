package engine

import (
	"testing"

	"akai.org.pl/joystick_server/engine"
)

func TestCreateThreeRoomsWithLastRoomCodeEqualsTwo(t *testing.T) {
	e := engine.Engine{}
	e.CreateRoom("basic", 5)
	e.CreateRoom("basic", 5)

	actualRoomNum := e.CreateRoom("basic", 5)
	expectedRoomNum := 2

	if actualRoomNum != expectedRoomNum {
		t.Errorf("expected code to be %v, but got %v", expectedRoomNum, actualRoomNum)
	}
}
