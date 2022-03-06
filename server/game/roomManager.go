package game

import (
	"akai.org.pl/joystick_server/logger"
	"errors"
	"fmt"
)

type RoomManager struct {
	rooms  map[string]*Room
	logger *logger.Logger
}

const (
	noSuchRoomMessage = "there is no game with such code"
	roomAlreadyExists = "game with such code already exists"
	roomCodeLength    = 5
)

var (
	availableGuis = []string{"ArrowsHorizontal", "ArrowsVertical", "ArrowsVertical1AB", "CrossArrows", "CrossArrows1AB"}
)

func NewRoomManager(logger *logger.Logger) *RoomManager {
	return &RoomManager{
		make(map[string]*Room),
		logger,
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

	if !manager.guiValid(gui) {
		return "", errors.New("the given GUI is not in set of valid GUIs")
	}

	room := NewRoom(gui, maxPlayers, manager.logger)
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

func (manager *RoomManager) guiValid(gui string) bool {
	for _, available := range availableGuis {
		if gui == available {
			manager.logger.Debug(fmt.Sprintf("Given GUI %v is available", gui))
			return true
		}
	}
	manager.logger.Debug(fmt.Sprintf("Given GUI %v not found in list of available GUIs", gui))
	return false
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
