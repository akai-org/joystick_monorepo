package room

import (
	"errors"
	"math/rand"
)

type Manager struct {
	rooms map[string]*Room
}

const (
	noSuchRoomMessage = "there is no room with such code"
	roomAlreadyExists = "room with such code already exists"
)

func New() *Manager {
	return &Manager{}
}

func (manager *Manager) GenerateCode() string {
	random := make([]byte, 5)
	for i := range random {
		random[i] = byte(rand.Intn(25)+66)
	}
	return string(random)
}

func (manager *Manager) AppendRoomWithCode(room *Room, code string) error {
	if _, ok := manager.rooms[code]; ok {
		return errors.New(roomAlreadyExists)
	}
	manager.rooms[code] = room
	return nil
}

func (manager *Manager) GetRoom(code string) (*Room, error) {
	if r, ok := manager.rooms[code]; !ok {
		return nil, errors.New(noSuchRoomMessage)
	} else {
		return r, nil
	}
}

func (manager *Manager) RemoveRoom(code string) error {
	if _, ok := manager.rooms[code]; !ok {
		return errors.New(noSuchRoomMessage)
	}
	delete(manager.rooms, code)
	return nil
}
