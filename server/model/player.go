package model

import "github.com/satori/go.uuid"

//Player describes a client in a game
type Player struct {
	Profile *Profile `json:"profile"`
	Color   bool     `json:"color"`
	Team    int      `json:"team"`
	Pieces  []*Piece `json:"pieces"`
	MyTurn  bool     `json:"myTurn"`
	Board   int      `json:"board"`
}

//GetPieceByID for easy access
func (p *Player) GetPieceByID(id string) (piece *Piece, found bool) {
	for _, p := range p.Pieces { //index and struct instances, similar to enum in python
		if p.ID == id {
			piece, found = p, true
			return
		}
	}
	return
}

func (p *Player) shouldGoDown() bool {
	return p.Team == 1
}

//SetTeamAndColor derives the color and team
func (p *Player) SetTeamColorAndBoard(spot int, boards int) {
	if spot <= boards {
		p.Team = 1
		p.Color = (spot+boards)%2 == 0
		p.Board = spot
	} else {
		p.Team = 2
		p.Color = (spot+boards)%2 == 1
		p.Board = spot - boards
	}

}

//SetupBoard initializes the board and players
//adds the required pieces. Mainly straightforward
func (p *Player) SetupBoard() {
	if !p.Color {
		p.MyTurn = true
	}

	//Pawns
	for i := 1; i <= 8; i++ {
		piece := &Piece{
			ID:       uuid.NewV4().String(),
			Color:    p.Color,
			X:        i,
			Identity: identityPawn,
			Cache:    false,
			Board:    p.Board,
		}
		if p.Team == 1 {
			piece.Y = 7 //x -> , y down-facing so that horizontal line of y is that number
		} else {
			piece.Y = 2
		}
		p.Pieces = append(p.Pieces, piece)
	}

	//Two free pawns :D
	for i := 1; i <= 8; i++ {
		piece := &Piece{
			ID:       uuid.NewV4().String(),
			Color:    p.Color,
			X:        0,
			Y:        0,
			Identity: identityPawn,
			Cache:    true,
			Board:    p.Board,
		}
		p.Pieces = append(p.Pieces, piece)
	}

	//Two rooks
	rookLeft := CreateRook(uuid.NewV4().String(), p.Board, p.Color)
	rookRight := CreateRook(uuid.NewV4().String(), p.Board, p.Color)
	if p.Team == 1 {
		rookLeft.Y = 8
		rookRight.Y = 8
	} else {
		rookLeft.Y = 1
		rookRight.Y = 1
	}

	rookLeft.X = 1
	rookRight.X = 8

	p.Pieces = append(p.Pieces, rookLeft)
	p.Pieces = append(p.Pieces, rookRight)

}

//OwnsPiece determines if the piece is owned by a player
func (p *Player) OwnsPiece(ID string) bool {
	for _, p := range p.Pieces {
		if p.ID == ID {
			return true
		}
	}
	return false
}
