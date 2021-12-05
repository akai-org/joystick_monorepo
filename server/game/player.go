package game

type Player struct {
	Nickname string
	Room *Room
}

func (p *Player) SendMessageToRoom(message []byte) {
	p.Room.PlayerChannel <- message
}

func NewPlayer(nickname string) *Player {
	return &Player{
		Nickname: nickname,
	}
}