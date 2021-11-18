package room

import (
	"akai.org.pl/joystick_server/games"
	"errors"
)

type Room struct {
	Gui     string `json:"gui"`
	players []*games.Player
}

const (
	noSpaceLeftInRoom = "this room has reached its max capacity of players"
)

func NewRoom(gui string, maxPlayers int) *Room {
	return &Room{gui, make([]*games.Player, 0, maxPlayers)}
}

func (r *Room) AddPlayer(player *games.Player) error {
	if cap(r.players) == len(r.players) {
		return errors.New(noSpaceLeftInRoom)
	}
	r.players = append(r.players, player)
	return nil
}
