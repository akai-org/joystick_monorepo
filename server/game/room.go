package game

import (
	"errors"
)

type room struct {
	Gui     string `json:"gui"`
	players []*Player
}

const (
	noSpaceLeftInRoom = "this game has reached its max capacity of players"
)

func NewRoom(gui string, maxPlayers int) *room {
	return &room{gui, make([]*Player, 0, maxPlayers)}
}

func (r *room) AddPlayer(player *Player) error {
	if cap(r.players) == len(r.players) {
		return errors.New(noSpaceLeftInRoom)
	}
	r.players = append(r.players, player)
	return nil
}
