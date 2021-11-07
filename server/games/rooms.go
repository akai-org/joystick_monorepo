package games

import (
	"errors"
	"fmt"
)

const noSpaceLeftInRoom = "max player count in room has been reached"

type room struct {
	players []Player
	maxPlayers int
	gui string
}

// dummy values, to be removed
var rooms = map[string]*room {
	"abcd": {
		players: make([]Player, 0),
		maxPlayers: 10,
		gui: "arrows-horizontal",
	},
	"aaaa": {
		players: make([]Player, 0),
		maxPlayers: 10,
		gui: "arrows-horizontal",
	},
	"bbbb": {
		players: make([]Player, 0),
		maxPlayers: 10,
		gui: "arrows-horizontal",
	},
	"cccc": {
		players: make([]Player, 0),
		maxPlayers: 10,
		gui: "arrows-horizontal",
	},
	"dddd": {
		players: make([]Player, 0),
		maxPlayers: 10,
		gui: "arrows-horizontal",
	},
}

func RoomExists(code string) bool {
	_, ok := rooms[code]
	return ok
}

func AddPlayerToRoom(player Player, roomCode string) error {
	room := rooms[roomCode]
	if len(room.players) >= room.maxPlayers {
		return errors.New(noSpaceLeftInRoom)
	}
	room.players = append(rooms[roomCode].players, player)
	fmt.Printf("%d player(s) in room %s\n", len(room.players), roomCode)
	return nil
}