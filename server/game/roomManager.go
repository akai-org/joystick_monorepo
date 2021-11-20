package game

import (
	"errors"
	"math/rand"
)

type Manager struct {
	rooms map[string]*room
}

const (
	noSuchRoomMessage = "there is no game with such code"
	roomAlreadyExists = "game with such code already exists"
)

func NewRoomManager() *Manager {
	return &Manager{}
}

func (manager *Manager) generateCode() string {
	random := make([]byte, 5)
	for i := range random {
		random[i] = byte(rand.Intn(25) + 66)
	}
	return string(random)
}

func (manager *Manager) appendRoomWithCode(room *room, code string) error {
	if _, ok := manager.rooms[code]; ok {
		return errors.New(roomAlreadyExists)
	}
	manager.rooms[code] = room
	return nil
}

func (manager *Manager) CreateNewRoom(gui string, maxPlayers int) (code string, err error) {
	room := NewRoom(gui, maxPlayers)
	code = manager.generateCode()

	err = manager.appendRoomWithCode(room, manager.generateCode())
	if err != nil {
		return "", err
	}

	return code, nil
}

func (manager *Manager) GetRoom(code string) (*room, error) {
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
