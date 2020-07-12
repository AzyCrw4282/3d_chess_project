package controller

import (
	"github.com/team142/angrychess/io/ws"
	"github.com/team142/angrychess/model"
	"log"
)

//Final class in Controller before requests are directed to the model base

//JoinGame gets a player into a game --finds a spot and sets them to join
func joinGame(game *model.Game, player *model.Player) bool {
	found, spot := game.FindSpot()
	if !found {
		return false
	}
	player.SetTeamColorAndBoard(spot, game.Boards)
	game.Players[spot] = player
	return true
}

//StartGame starts the game for all players-- starts and sets the share state for all player data transmission
func startGame(game *model.Game) {
	ok, msg := game.IsReadyToStart() //short var declare
	if !ok {
		reply := model.CreateMessageError("Failed to start game", msg)
		game.Owner.Profile.Client.SendObject(reply)
		return
	}

	game.SetupBoards()
	game.Started = true
	shareState(game)

}

//ShareState tells all players what is going on -- important function, similar to implementaion in java
func shareState(game *model.Game) {
	reply := model.CreateMessageShareState(game)
	for _, player := range game.Players {
		player.Profile.Client.SendObject(reply)
	}

}

//Move moves a piece -- performs all required checks. Can be simplfiied and overal improved
func Move(game *model.Game, client *ws.Client, message *model.MessageMove) (didMove bool) {
	log.Println(">> Moving ")
	pieceFound, piece, piecePlayer := game.FindPiece(message.PieceID)
	//Can't move a piece that does not exists
	if !pieceFound {
		log.Println("Piece not found, " + message.PieceID)
		return
	}
	//Get the message sending player
	player, _, _ := game.PlayerByClient(client)
	//Can this player move?
	if !player.MyTurn {
		log.Println("Not my turn!, " + message.PieceID)
		return
	}
	//TODO: handle taking pieces off board here
	if piece.IsEqual(message) {
		log.Println("No move")
		return
	}
	//current player same as piece-player?
	if player != piecePlayer {
		log.Println("Player does not own piece, " + message.PieceID)
		return
	}

	//Check for bad state
	if message.Cache == false && message.Board == 0 {
		log.Println("Bad request.  Piece must be on board or in cache, not neither")
		return
	}

	description := model.CalcMoveDescription(game, player, piece, message)

	ok, taken, msg := model.IsMovePossible(player, piece, description)
	if !ok {
		log.Println(msg)
		return
	}

	if taken != nil {
		taken.Cache = true
		//TODO take piece
		//Switch players
		//TakePiece(game, player, taken)
	}

	/*
		TODO: do other checks
	*/
	//if !player.OwnsPiece(move.PieceID) {
	//	err = fmt.Errorf("player doesnt not own piece: %s", move.PieceID)
	//}

	didMove = true
	piece.Move(message)
	return
}

//announce announces something to all players in a given game, i.e. whole game broadcast
func announce(game *model.Game, i interface{}) {
	for _, player := range game.Players {
		player.Profile.Client.SendObject(i)
	}
}
