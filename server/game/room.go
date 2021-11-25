package game

import (
	"errors"
)

type room struct {
	Gui     string `json:"gui"`
	players []*player
}

const (
	noSpaceLeftInRoom = "this game has reached its max capacity of players"
)

func NewRoom(gui string, maxPlayers int) *room {
	return &room{gui, make([]*player, 0, maxPlayers)}
}

func (r *room) AddPlayer(player *player) error {
	if cap(r.players) == len(r.players) {
		return errors.New(noSpaceLeftInRoom)
	}
	player.Room = r
	r.players = append(r.players, player)
	return nil
}
