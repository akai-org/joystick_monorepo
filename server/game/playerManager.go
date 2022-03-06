package game

import (
	"akai.org.pl/joystick_server/logger"
	"errors"
	"fmt"
)

type PlayerManager struct {
	players map[string]*Player
	logger  *logger.Logger
}

const (
	playerCodeLength          = 32
	playerAlreadyExistMessage = "Player with such code already exists"
)

func NewPlayerManager(logger *logger.Logger) *PlayerManager {
	return &PlayerManager{
		make(map[string]*Player),
		logger,
	}
}

func (manager *PlayerManager) GetPlayer(id string) (*Player, error) {
	if r, ok := manager.players[id]; !ok {
		message := fmt.Sprintf("Player with id %v has not been found", id)
		manager.logger.Debug(message)
		return nil, errors.New(message)
	} else {
		return r, nil
	}
}

func (manager *PlayerManager) CreateNewPlayer(nickname string) (*Player, string) {
	player := NewPlayer(nickname, manager.logger)
	code := generateCode(playerCodeLength)
	for {
		err := manager.appendPlayerWithCode(player, code)
		if err == nil {
			break
		}
		code = generateCode(roomCodeLength)
		continue
	}
	manager.logger.Debug(fmt.Sprintf("Player with id %v has been successfuly created", code))
	return player, code
}

func (manager *PlayerManager) appendPlayerWithCode(player *Player, code string) error {
	if _, ok := manager.players[code]; ok {
		return errors.New(playerAlreadyExistMessage)
	}
	manager.players[code] = player
	return nil
}
