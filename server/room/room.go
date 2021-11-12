package room

type Room struct {
	Code       int    `json:"code"`
	GUI        string `json:"gui"`
	MaxPlayers int    `json:"maxPlayers"`
}
