package room

import "fmt"

type RoomManager struct {
	rooms []Room
}

func (rm *RoomManager) CreateRoom(gui string, maxPlayers int) (code int) {
	minAvailableCode := 0
	minAvailableIndex := 0
	for _, r := range rm.rooms {
		if r.Code == minAvailableCode {
			minAvailableCode += 1
			minAvailableIndex += 1
		} else if r.Code > minAvailableCode {
			break
		}
	}

	newRoom := Room{minAvailableCode, gui, maxPlayers}
	if len(rm.rooms) > 0 && minAvailableIndex < len(rm.rooms) {
		rm.rooms = append(rm.rooms[:minAvailableIndex+1], rm.rooms[minAvailableIndex:]...)
		rm.rooms[minAvailableIndex] = newRoom
	} else {
		rm.rooms = append(rm.rooms, newRoom)
	}

	return minAvailableCode
}

func (rm *RoomManager) GetRoom(code int) (room *Room, err error) {
	for _, r := range rm.rooms {
		if r.Code == code {
			return &r, nil
		}
	}

	return nil, fmt.Errorf("no room with code %d", code)
}

func (rm *RoomManager) RemoveRoom(code int) (err error) {
	for i, r := range rm.rooms {
		if r.Code == code {
			rm.rooms = append(rm.rooms[:i], rm.rooms[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("no room with code %d", code)
}
