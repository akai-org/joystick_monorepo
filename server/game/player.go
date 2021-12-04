package game

type player struct {
	Nickname string
	Room *room
}

func NewPlayer(nickname string) *player {
	return &player{
		Nickname: nickname,
	}
}