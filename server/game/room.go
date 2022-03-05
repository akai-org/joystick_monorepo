package game

import (
	"errors"
)

type Room struct {
	Gui           string `json:"gui"`
	players       []*Player
	PlayerChannel chan []byte
}

const (
	noSpaceLeftInRoom = "this game has reached its max capacity of players"
)

func NewRoom(gui string, maxPlayers int) *Room {
	return &Room{
		gui,
		make([]*Player, 0, maxPlayers),
		make(chan []byte),
	}
}

func (r *Room) AddPlayer(player *Player) error {
	if cap(r.players) == len(r.players) {
		return errors.New(noSpaceLeftInRoom)
	}
	player.Room = r
	r.players = append(r.players, player)
	return nil
}
