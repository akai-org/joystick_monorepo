package room

import (
	"reflect"
	"testing"
)

func TestCreateThreeRoomsWithLastRoomCodeEqualsTwo(t *testing.T) {
	rm := RoomManager{}
	rm.CreateRoom("basic", 5)
	rm.CreateRoom("basic", 5)
	actualRoomNum := rm.CreateRoom("basic", 5)
	expectedRoomNum := 2

	if actualRoomNum != expectedRoomNum {
		t.Errorf("expected code to be %v, but got %v", expectedRoomNum, actualRoomNum)
	}
}

func TestCreateRoomAndGetIt(t *testing.T) {
	rm := RoomManager{}
	actualRoomCode := rm.CreateRoom("basic", 5)
	actualRoom, _ := rm.GetRoom(actualRoomCode)
	expectedRoom := Room{Code: 0, GUI: "basic", MaxPlayers: 5}

	if !reflect.DeepEqual(*actualRoom, expectedRoom) {
		t.Errorf("expected room to be %v, but got %v", expectedRoom, *actualRoom)
	}
}

func TestGetNotExistentRoomAndReturnError(t *testing.T) {
	rm := RoomManager{}
	actualRoom, actualErr := rm.GetRoom(0)

	if actualRoom != nil {
		t.Errorf("expected room to be nil, but got %v", *actualRoom)
	}

	if actualErr == nil {
		t.Errorf("expected error to be not nil, but got nil")
	}
}

func TestCreateThreeRoomsAndRemoveSecondAndCheckIfSecondExists(t *testing.T) {
	rm := RoomManager{}
	firstRoomCode := rm.CreateRoom("basic", 5)
	secondRoomCode := rm.CreateRoom("basic", 5)
	thirdRoomCode := rm.CreateRoom("basic", 5)

	rm.RemoveRoom(secondRoomCode)

	if _, err := rm.GetRoom(firstRoomCode); err != nil {
		t.Errorf("expected first room to exist")
	}

	if _, err := rm.GetRoom(secondRoomCode); err == nil {
		t.Errorf("expected second element not to exist")
	}

	if _, err := rm.GetRoom(thirdRoomCode); err != nil {
		t.Errorf("expected third room to exist")
	}
}

func TestCreateThreeRoomsAndRemoveSecondAndAddTwoRoomsAndCheckIfExpectedRoomsExistAndTheirCodesAreCorrect(t *testing.T) {
	rm := RoomManager{}
	rm.CreateRoom("basic_1", 5)
	rm.CreateRoom("basic_2", 5)
	rm.CreateRoom("basic_3", 5)

	rm.RemoveRoom(1)

	rm.CreateRoom("basic_4", 5)
	rm.CreateRoom("basic_5", 5)

	expectedFirstRoom := Room{Code: 0, GUI: "basic_1", MaxPlayers: 5}
	expectedSecondRoom := Room{Code: 1, GUI: "basic_4", MaxPlayers: 5}
	expectedThirdRoom := Room{Code: 2, GUI: "basic_3", MaxPlayers: 5}
	expectedFourthRoom := Room{Code: 3, GUI: "basic_5", MaxPlayers: 5}

	if r, err := rm.GetRoom(0); err != nil || !reflect.DeepEqual(expectedFirstRoom, *r) {
		t.Errorf("expected desired room %v to exist, but got %v", expectedFirstRoom, *r)
	}

	if r, err := rm.GetRoom(1); err != nil || !reflect.DeepEqual(expectedSecondRoom, *r) {
		t.Errorf("expected desired room %v to exist, but got %v", expectedSecondRoom, *r)
	}

	if r, err := rm.GetRoom(2); err != nil || !reflect.DeepEqual(expectedThirdRoom, *r) {
		t.Errorf("expected desired room %v to exist, but got %v", expectedThirdRoom, *r)
	}

	if r, err := rm.GetRoom(3); err != nil || !reflect.DeepEqual(expectedFourthRoom, *r) {
		t.Errorf("expected desired room %v to exist, but got %v", expectedFourthRoom, *r)
	}
}
