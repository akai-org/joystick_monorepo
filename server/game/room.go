package game

import (
	"akai.org.pl/joystick_server/logger"
	"encoding/json"
	"errors"
	"fmt"
)

type Room struct {
	Gui                  string `json:"gui"`
	players              []*Player
	PlayerChannel        chan []byte
	logger               *logger.Logger
	CommunicationChannel chan string
	occupiedIds          map[uint8]bool
}

const (
	noSpaceLeftInRoom  = "this game has reached its max capacity of players"
	playerAddedEvent   = "player_added"
	playerRemovedEvent = "player_removed"
)

type playerMessage struct {
	EventName string `json:"event_name"`
	Nickname  string `json:"nickname"`
	Id        uint8  `json:"id"`
}

func NewRoom(gui string, maxPlayers int, logger *logger.Logger) *Room {
	return &Room{
		gui,
		make([]*Player, 0, maxPlayers),
		make(chan []byte),
		logger,
		make(chan string),
		make(map[uint8]bool),
	}
}

func (r *Room) AddPlayer(player *Player) error {
	if cap(r.players) == len(r.players) {
		return errors.New(noSpaceLeftInRoom)
	}
	player.Room = r
	r.players = append(r.players, player)
	for i := uint8(0); i < ^uint8(0); i++ {
		if _, ok := r.occupiedIds[i]; !ok {
			r.occupiedIds[i] = true
			player.roomPlayerId = i
			break
		}
	}

	messagePayload := &playerMessage{
		playerAddedEvent,
		player.Nickname,
		player.roomPlayerId,
	}
	message, err := json.Marshal(messagePayload)
	if err != nil {
		r.logger.Debug(fmt.Sprintf("Failed to send notification about player joining"))
	}
	r.CommunicationChannel <- string(message)

	r.logger.Debug(fmt.Sprintf("Player %v has been added to room and got id %v", player.Nickname, player.roomPlayerId))
	return nil
}
