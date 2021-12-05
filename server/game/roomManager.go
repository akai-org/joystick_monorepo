package game

import (
	"errors"
)

type RoomManager struct {
	rooms map[string]*Room
}

const (
	noSuchRoomMessage = "there is no game with such code"
	roomAlreadyExists = "game with such code already exists"
	roomCodeLength    = 5
)

func NewRoomManager() *RoomManager {
	return &RoomManager{
		make(map[string]*Room),
	}
}

func (manager *RoomManager) appendRoomWithCode(room *Room, code string) error {
	if _, ok := manager.rooms[code]; ok {
		return errors.New(roomAlreadyExists)
	}
	manager.rooms[code] = room
	return nil
}

func (manager *RoomManager) CreateNewRoom(gui string, maxPlayers int) (string, error) {
	room := NewRoom(gui, maxPlayers)
	code := generateCode(roomCodeLength)

	for {
		err := manager.appendRoomWithCode(room, code)
		if err == nil {
			break
		}
		code = generateCode(roomCodeLength)
		continue
	}

	return code, nil
}

func (manager *RoomManager) GetRoom(code string) (*Room, error) {
	if r, ok := manager.rooms[code]; !ok {
		return nil, errors.New(noSuchRoomMessage)
	} else {
		return r, nil
	}
}

func (manager *RoomManager) RemoveRoom(code string) error {
	if _, ok := manager.rooms[code]; !ok {
		return errors.New(noSuchRoomMessage)
	}
	delete(manager.rooms, code)
	return nil
}
