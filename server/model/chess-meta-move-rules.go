package model

import (
	"log"
)

func IsMovePossible(player *Player, piece *Piece, description *MoveDescription) (ok bool, taken *Piece, msg string) {
	if piece.Identity == identityPawn {
		return isMovePossiblePawn(player, description)
	} else if piece.Identity == identityRook {
		return isMovePossibleRook(player, description)
	} else if piece.Identity == identityBishop {
		return isMovePossibleBishop(player, description)
	}
	ok = false
	msg = "Identity not implemented. Cannot check isMovePossible"
	log.Println(msg)
	return
}
