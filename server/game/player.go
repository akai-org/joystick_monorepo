package game

import (
	"akai.org.pl/joystick_server/logger"
	"fmt"
)

type Player struct {
	Nickname     string
	Room         *Room
	logger       *logger.Logger
	roomPlayerId uint8
}

func (p *Player) SendMessageToRoom(message []byte) {
	p.logger.Debug(fmt.Sprintf("Player %v sending to channel %v", p.Nickname, p.Room.PlayerChannel))
	commandType := message[0]
	action := message[1]
	p.Room.PlayerChannel <- []byte{p.roomPlayerId, commandType, action}
}

func NewPlayer(nickname string, logger *logger.Logger) *Player {
	return &Player{
		Nickname: nickname,
		logger:   logger,
	}
}
