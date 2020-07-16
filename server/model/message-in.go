package model

//includees intital msgs for all and most self-explanatory.
//BaseMessage for all incoming messages
type BaseMessage struct {
	Msg string `json:"msg"`
}

//MessageNick describes a nickname change
type MessageNick struct {
	Nick string `json:"nick"`
}

//MessageJoinGame for joining a game by id
type MessageJoinGame struct {
	ID string `json:"id"`
}

//MessageMove describes a movement
type MessageMove struct {
	PieceID string `json:"pieceId"`
	ToX     int    `json:"tx"`
	ToY     int    `json:"ty"`
	Board   int    `json:"board"`
	Cache   bool   `json:"cache"`
}

//MessagePlace describes placing a piece
type MessagePlace struct {
	ID  string `json:"id"`
	ToX int    `json:"tx"`
	ToY int    `json:"ty"`
}

type MessageChangeSeat struct {
	Seat int `json:"seat"`
}
