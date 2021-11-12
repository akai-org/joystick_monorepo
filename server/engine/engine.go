package engine

import (
	"fmt"
)

type Room struct {
	Code       int
	GUI        string
	MaxPlayers int
}

type Engine struct {
	rooms []Room
}

func (e *Engine) CreateRoom(gui string, maxPlayers int) (code int) {
	minAvailableCode := 0
	minAvailableIndex := 0
	for _, r := range e.rooms {
		if r.Code == minAvailableCode {
			minAvailableCode += 1
			minAvailableIndex += 1
		} else if r.Code > minAvailableCode {
			break
		}
	}

	newRoom := Room{minAvailableCode, gui, maxPlayers}
	if len(e.rooms) > 0 && minAvailableIndex < len(e.rooms) {
		e.rooms = append(e.rooms[:minAvailableIndex+1], e.rooms[minAvailableIndex:]...)
		e.rooms[minAvailableIndex] = newRoom
	} else {
		e.rooms = append(e.rooms, newRoom)
	}

	return minAvailableCode
}

func (e *Engine) GetRoom(code int) (room *Room, err error) {
	for _, r := range e.rooms {
		if r.Code == code {
			return &r, nil
		}
	}
	return nil, fmt.Errorf("no room with code %d", code)
}
