package model

import (
	"github.com/team142/angrychess/util"
)

//TODO: check all below
type MoveDescription struct {
	XDiff             int
	YDiff             int
	Down              bool //Y is decreasing
	Diagonal          bool
	BeingPlaced       bool
	BeingRemoved      bool
	MovingBoards      bool
	PawnOnSpawn       bool
	LastTwoRows       bool
	OtherBoard        bool
	LandingOnPiece    *Piece
	LandingOnPieceOwn bool
	PiecesBetween     []*Piece
}

func CalcMoveDescription(game *Game, player *Player, piece *Piece, move *MessageMove) *MoveDescription {
	result := &MoveDescription{} //all passed byRef for the benefit that valuec an be directly changed

	//Change in vertical tiles
	result.YDiff = util.Abs(piece.Y - move.ToY)

	//Change in horizontal tiles
	result.XDiff = util.Abs(piece.X - move.ToX)

	//Is the piece moving down the board
	result.Down = piece.Y > move.ToY

	//Is it going from off the board to on the board
	result.BeingPlaced = piece.Cache && !move.Cache

	//Is it going from the board to the cache
	result.BeingRemoved = !piece.Cache && move.Cache // the separate board

	//Is it moving to another board
	result.MovingBoards = piece.Board != move.Board

	//Is the pawn on it's starting row
	result.PawnOnSpawn = player.Team == 1 && piece.Y == 7 || player.Team == 2 && piece.Y == 2

	//Is the placement on the last two of rows for that players direction
	result.LastTwoRows = (1 == player.Team && 2 >= piece.Y) || (piece.Y >= 7 && 2 == player.Team)

	//Is it not the player's board
	result.OtherBoard = player.Board != move.Board

	//Pieces between the from tile and to tile
	CalcPiecesBetween(game, player, piece, move, result)

Outer: // a label in golang
	for _, pl := range game.Players {
		for _, pi := range pl.Pieces {
			if pi.IsEqual(move) {
				result.LandingOnPiece = pi
				result.LandingOnPieceOwn = pl.Profile.ID == player.Profile.ID
				break Outer
			}
		}
	}
	return result

}

//TODO: understand logic and workout uml diagrams.
//main method that checks for pieces in between it.
func CalcPiecesBetween(game *Game, player *Player, piece *Piece, move *MessageMove, result *MoveDescription) {
	//Don't worry about one tile
	if result.XDiff+result.YDiff <= 1 {
		return
	}
	//Knights don't worry about pieces between
	if piece.Identity == identityKnight {
		return
	}

	if result.XDiff <= 1 && result.YDiff <= 1 {
		//Ignore one tile moves
		return
	}

	//Horizontal moves and is also greater than 1 diff
	if result.XDiff > 1 && result.YDiff == 0 {
		lower, higher := util.OrderPoints(move.ToX, piece.X)
		for x := lower + 1; x < higher; x++ {
			found, p := game.GetPieceAtPoint(piece.Board, x, piece.Y)
			if found {
				result.PiecesBetween = append(result.PiecesBetween, p)
			}
		}
	} else if result.XDiff == 0 && result.YDiff > 1 {
		//Vertical moves
		lower, higher := util.OrderPoints(move.ToY, piece.Y)
		for y := lower + 1; y < higher; y++ {
			found, p := game.GetPieceAtPoint(piece.Board, piece.X, y)
			if found {
				result.PiecesBetween = append(result.PiecesBetween, p)
			}
		}
	}

	if result.XDiff == result.YDiff {
		start := &Tile{X: piece.X, Y: piece.Y}
		end := &Tile{X: move.ToX, Y: move.ToY}
		c := start.GetTilesUntil(end)
		for tile := range c {
			found, p := game.GetPieceAtPoint(piece.Board, tile.X, tile.Y)
			if found {
				result.PiecesBetween = append(result.PiecesBetween, p)
			}
		}
	}

	//TODO: do diagonal moves
	//

}
