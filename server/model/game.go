package model

//improt statments
import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/team142/angrychess/io/ws"
	"log"
)

//ListOfGames describes a list of games on the server
type ListOfGames struct {
	Games []map[string]string `json:"games"` //optional string values for the field
}

//CreateGameAndRun creates a game and starts it
func CreateGameAndRun(creator *Player) *Game {
	game := &Game{ //saves the memory address of this short var declaration
		ID:       uuid.NewV4().String(),
		Players:  make(map[int]*Player),
		Boards:   maxSupportedBoards,
		Title:    fmt.Sprintf("%s's game", creator.Profile.Nick),
		Commands: make(chan func(*Game), 256),
		stop:     make(chan bool),
	}
	game.Players[1] = creator
	game.Owner = creator
	creator.SetTeamColorAndBoard(1, game.Boards)
	go game.run()
	return game
}

//Game describes a game on the server
type Game struct {
	ID                 string           `json:"id"`
	Started            bool             `json:"started"`
	Title              string           `json:"title"`
	Owner              *Player          `json:"owner"`
	Players            map[int]*Player  `json:"players"`
	Boards             int              `json:"boards"`
	CanStartBeforeFull bool             `json:"canStartBeforeFull"`
	Commands           chan func(*Game) `json:"-"`
	stop               chan bool        `json:"-"`
}

//DoWork Adds a function of work that must run in the game's go-routine.
func (game *Game) DoWork(f func(*Game)) {
	game.Commands <- f
}

func (game *Game) run() {
	for {
		select {
		case command := <-game.Commands:
			command(game)
		case <-game.stop:
			close(game.Commands)
			close(game.stop)
			log.Println("Stopping game runner")
			return
		}
	}
}

//Stop stops the game from running
func (game *Game) Stop() {
	game.stop <- true
}

func (game *Game) FindSpot() (found bool, spot int) {
	if len(game.Players) >= game.MaxPlayers() {
		return false, 0
	}
	for s := 1; s <= game.MaxPlayers(); s++ {
		if game.Players[s] == nil {
			return true, s
		}
	}
	return false, 0
}

//FindPiece fiends a piece and it's context
func (game *Game) FindPiece(pieceID string) (found bool, piece *Piece, player *Player) {
	for _, player := range game.Players {
		piece, found := player.GetPieceByID(pieceID)
		if found {
			return true, piece, player
		}
	}
	found = false
	return
}

//GetPieceAtPoint finds a piece by board, x and y
func (game *Game) GetPieceAtPoint(board, x, y int) (found bool, piece *Piece) {
	for _, player := range game.Players {
		for _, piece = range player.Pieces {
			if piece.Board == board && piece.X == x && piece.Y == y {
				found = true
				return
			}
		}
	}
	found = false
	return
}

//MaxPlayers determines the max number of players
func (game *Game) MaxPlayers() int {
	return game.Boards * 2
}

//PlayerByClient for easy access
func (game *Game) PlayerByClient(client *ws.Client) (result *Player, seat int, found bool) {
	for seat, p := range game.Players {
		if p.Profile.Client == client {
			return p, seat, true
		}
	}
	return
}

//SetupBoards initializes all boards in the game
func (game *Game) SetupBoards() {
	for _, player := range game.Players {
		player.SetupBoard()
	}

}

//IsReadyToStart checks that the game can start
func (game *Game) IsReadyToStart() (ok bool, message string) {
	if game.CanStartBeforeFull {
		ok = true
		return
	}
	ok = game.MaxPlayers() == len(game.Players)
	if !ok {
		message = "Not enough players"
	}
	return
}

func (game *Game) ChangeSeat(client *ws.Client, seat int) {
	if game.Players[seat] != nil {
		reply := CreateMessageError("Failed to move seats", "Seat taken")
		client.SendObject(reply)
		return
	}
	currentSeat := 0
	var currentPlayer *Player
	for seatN, playerN := range game.Players {
		if playerN.Profile.Client == client {
			currentSeat = seatN
			currentPlayer = playerN
			break
		}
	}
	if currentSeat == 0 {
		log.Fatal("Could not find client in game.. dying")
	}
	delete(game.Players, currentSeat)
	game.Players[seat] = currentPlayer
	currentPlayer.SetTeamColorAndBoard(seat, game.Boards)

}

func (game *Game) RemoveClient(client *ws.Client) {
	pid := 0
	for seat, player := range game.Players {
		if player.Profile.Client == client {
			pid = seat
			break
		}
	}
	if pid != 0 {
		log.Println(">> Removing player with seat ", pid)
		delete(game.Players, pid)
	} else {
		log.Println(">> Could not find player to remove from game")
	}

	fmt.Println("There are now how many players ", len(game.Players))

	if game.Owner.Profile.Client == client {
		game.Owner = nil
		for a := range game.Players {
			log.Println(">> New owner is ", game.Players[a].Profile.Nick)
			game.Owner = game.Players[a]
			break
		}
		if game.Owner == nil && len(game.Players) > 0 {
			log.Println("Error: did not assign new owner.")
		}

	}

}

func (game *Game) ChangeMoveFrom(client *ws.Client) {
	var newPlayer *Player
	player, seat, _ := game.PlayerByClient(client)
	if seat <= game.Boards {
		newPlayer = game.Players[seat+2]
	} else {
		newPlayer = game.Players[seat-2]
	}
	player.MyTurn = false
	newPlayer.MyTurn = true

}
