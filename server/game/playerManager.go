package game

import (
	"errors"
)

type PlayerManager struct {
	players map[string]*player
}

const (
	playerCodeLength          = 32
	playerAlreadyExistMessage = "player with such code already exists"
)

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{
		make(map[string]*player),
	}
}

func (manager *PlayerManager) GetPlayer(id string) (*player, error) {
	if r, ok := manager.players[id]; !ok {
		return nil, errors.New(noSuchRoomMessage)
	} else {
		return r, nil
	}
}

func (manager *PlayerManager) CreateNewPlayer(nickname string) (*player) {
	player := NewPlayer(nickname)
	code := generateCode(playerCodeLength)

	for {
		err := manager.appendPlayerWithCode(player, code)
		if err == nil {
			break
		}
		code = generateCode(roomCodeLength)
		continue
	}

	return player
}

func (manager *PlayerManager) appendPlayerWithCode(player *player, code string) error {
	if _, ok := manager.players[code]; ok {
		return errors.New(playerAlreadyExistMessage)
	}
	manager.players[code] = player
	return nil
}
