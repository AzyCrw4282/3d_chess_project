package model

/*
Handles pieces, all objects excluding squares are pieces on the board
*/

const (
	identityPawn   = 1
	identityKnight = 2
	identityBishop = 3
	identityRook   = 4
	identityQueen  = 5
	identityKing   = 6
)

//Piece describes a piece on the board
type Piece struct {
	ID       string `json:"id"`
	Identity int    `json:"identity"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	Color    bool   `json:"color"`
	Cache    bool   `json:"cache"`
	Board    int    `json:"board"`
}

//Move moves piece
func (piece *Piece) Move(message *MessageMove) {
	piece.Cache = message.Cache
	piece.Board = message.Board
	if !piece.Cache {
		piece.X = message.ToX
		piece.Y = message.ToY
	}

}

func IsLastTwo(player *Player, y int) bool {
	return (1 == player.Team && 2 >= y) || (y >= 7 && 2 == player.Team)
}

func (piece *Piece) IsEqual(move *MessageMove) bool {
	if piece.Board != move.Board {
		return false
	}
	if piece.Cache != move.Cache {
		return false
	}
	if piece.X != move.ToX {
		return false
	}
	if piece.Y != move.ToY {
		return false
	}
	return true
}

func CreateRook(id string, board int, color bool) *Piece {
	return &Piece{
		ID:       id,
		Board:    board,
		Color:    color,
		Identity: identityRook,
	}
}
