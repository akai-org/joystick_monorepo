package game

import (
	"errors"
)

type PlayerManager struct {
	players map[string]*Player
}

const (
	playerCodeLength          = 32
	playerAlreadyExistMessage = "Player with such code already exists"
)

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{
		make(map[string]*Player),
	}
}

func (manager *PlayerManager) GetPlayer(id string) (*Player, error) {
	if r, ok := manager.players[id]; !ok {
		return nil, errors.New(noSuchRoomMessage)
	} else {
		return r, nil
	}
}

func (manager *PlayerManager) CreateNewPlayer(nickname string) (*Player, string) {
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

	return player, code
}

func (manager *PlayerManager) appendPlayerWithCode(player *Player, code string) error {
	if _, ok := manager.players[code]; ok {
		return errors.New(playerAlreadyExistMessage)
	}
	manager.players[code] = player
	return nil
}
