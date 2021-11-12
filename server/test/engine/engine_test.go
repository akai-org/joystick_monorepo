package engine

import (
	"reflect"
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

func TestCreateRoomAndGetIt(t *testing.T) {
	e := engine.Engine{}
	actualRoomCode := e.CreateRoom("basic", 5)
	actualRoom, _ := e.GetRoom(actualRoomCode)
	expectedRoom := engine.Room{Code: 0, GUI: "basic", MaxPlayers: 5}

	if !reflect.DeepEqual(*actualRoom, expectedRoom) {
		t.Errorf("expected room to be %v, but got %v", expectedRoom, *actualRoom)
	}
}

func TestGetNotExistentRoomAndReturnError(t *testing.T) {
	e := engine.Engine{}
	actualRoom, actualErr := e.GetRoom(0)

	if actualRoom != nil {
		t.Errorf("expected room to be nil, but got %v", *actualRoom)
	}

	if actualErr == nil {
		t.Errorf("expected error to be not nil, but got nil")
	}
}

func TestCreateThreeRoomsAndRemoveSecondAndCheckIfSecondExists(t *testing.T) {
	e := engine.Engine{}
	firstRoomCode := e.CreateRoom("basic", 5)
	secondRoomCode := e.CreateRoom("basic", 5)
	thirdRoomCode := e.CreateRoom("basic", 5)

	e.RemoveRoom(secondRoomCode)

	if _, err := e.GetRoom(firstRoomCode); err != nil {
		t.Errorf("expected first room to exist")
	}

	if _, err := e.GetRoom(secondRoomCode); err == nil {
		t.Errorf("expected second element not to exist")
	}

	if _, err := e.GetRoom(thirdRoomCode); err != nil {
		t.Errorf("expected third room to exist")
	}
}

func TestCreateThreeRoomsAndRemoveSecondAndAddTwoRoomsAndCheckIfExpectedRoomsExistAndTheirCodesAreCorrect(t *testing.T) {
	e := engine.Engine{}
	e.CreateRoom("basic_1", 5)
	e.CreateRoom("basic_2", 5)
	e.CreateRoom("basic_3", 5)

	e.RemoveRoom(1)

	e.CreateRoom("basic_4", 5)
	e.CreateRoom("basic_5", 5)

	expectedFirstRoom := engine.Room{Code: 0, GUI: "basic_1", MaxPlayers: 5}
	expectedSecondRoom := engine.Room{Code: 1, GUI: "basic_4", MaxPlayers: 5}
	expectedThirdRoom := engine.Room{Code: 2, GUI: "basic_3", MaxPlayers: 5}
	expectedFourthRoom := engine.Room{Code: 3, GUI: "basic_5", MaxPlayers: 5}

	if r, err := e.GetRoom(0); err != nil || !reflect.DeepEqual(expectedFirstRoom, *r) {
		t.Errorf("expected desired room %v to exist, but got %v", expectedFirstRoom, *r)
	}

	if r, err := e.GetRoom(1); err != nil || !reflect.DeepEqual(expectedSecondRoom, *r) {
		t.Errorf("expected desired room %v to exist, but got %v", expectedSecondRoom, *r)
	}

	if r, err := e.GetRoom(2); err != nil || !reflect.DeepEqual(expectedThirdRoom, *r) {
		t.Errorf("expected desired room %v to exist, but got %v", expectedThirdRoom, *r)
	}

	if r, err := e.GetRoom(3); err != nil || !reflect.DeepEqual(expectedFourthRoom, *r) {
		t.Errorf("expected desired room %v to exist, but got %v", expectedFourthRoom, *r)
	}
}
