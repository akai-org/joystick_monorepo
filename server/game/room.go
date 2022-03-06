package game

import (
	"akai.org.pl/joystick_server/logger"
	"errors"
	"fmt"
)

type Room struct {
	Gui           string `json:"gui"`
	players       []*Player
	PlayerChannel chan []byte
	logger        *logger.Logger
}

const (
	noSpaceLeftInRoom = "this game has reached its max capacity of players"
)

func NewRoom(gui string, maxPlayers int, logger *logger.Logger) *Room {
	return &Room{
		gui,
		make([]*Player, 0, maxPlayers),
		make(chan []byte),
		logger,
	}
}

func (r *Room) AddPlayer(player *Player) error {
	if cap(r.players) == len(r.players) {
		return errors.New(noSpaceLeftInRoom)
	}
	player.Room = r
	r.players = append(r.players, player)
	player.roomPlayerId = uint8(len(r.players))
	r.logger.Debug(fmt.Sprintf("Player %v has been added to room and got id %v", player.Nickname, player.roomPlayerId))
	return nil
}
