package model

import (
	"fmt"
	"github.com/team142/angrychess/io/ws"
)

const (
	maxSupportedBoards = 2
)

//Server holds server state
type Server struct {
	Address            string
	Lobby              map[*ws.Client]*Profile // hashmap interface maps client id to profile
	Games              map[string]*Game        //maps the string to the game
	handler            func(*Server, *ws.Client, *[]byte)
	Todo               chan *item // a channel in send/receive way ->https://gobyexample.com/channels
	CanStartBeforeFull bool
}

type item struct {
	client *ws.Client
	msg    *[]byte
}

//CreateServer starts a new server and returns a new server. Defitinion as defined *server with the pointer
func CreateServer(address string, handler func(*Server, *ws.Client, *[]byte), canStartBeforeFull bool) *Server {
	s := &Server{ //address is returned for the struct including
		Address:            address,
		handler:            handler,
		Lobby:              make(map[*ws.Client]*Profile),
		Games:              make(map[string]*Game),
		Todo:               make(chan *item, 256),
		CanStartBeforeFull: canStartBeforeFull,
	}
	s.run()
	return s
}

// runs the server using a poitner receiver to directly apply byRef change
func (s *Server) run() {
	go func() {
		for i := range s.Todo {
			s.handler(s, i.client, i.msg)
		}
	}()
}

//GameByClientOwner finds a game owned by client. Fings the game using a loop check and
//checking every game and if the profile of the client mathches the value passed by the byRef funcution.
func (s *Server) GameByClientOwner(client *ws.Client) (found bool, game *Game) {
	for _, game := range s.Games {
		if game.Owner.Profile.Client == client {
			return true, game
		}
	}
	return
}

//GameByClientPlaying find any player in a game
func (s *Server) GameByClientPlaying(client *ws.Client) (found bool, game *Game) {
	for _, game := range s.Games {
		for _, player := range game.Players {
			if player.Profile.Client == client { //nested loop to fidn a client, using a game->player->profile method
				return true, game
			}
		}
	}
	return
}

//HandleMessage This message is called by other parts of the system - the interface to the server
func (s *Server) HandleMessage(client *ws.Client, msg *[]byte) {
	i := &item{
		client: client,
		msg:    msg,
	}
	s.Todo <- i // passes item struct into the toodo channel to communicate with multiple goroutines

}

//GetOrCreateProfile creates profiles from a websocket client
//simply checks if a profile exists, else creates a profile and then assigns that client in the lobby with the profile
func (s *Server) GetOrCreateProfile(client *ws.Client) *Profile {
	p := s.Lobby[client]
	if p == nil {
		p = createProfile(client)
		s.Lobby[client] = p
	}
	return p
}

//ListOfGames produces a light struct that describes the games hosted
func (s *Server) CreateListOfGames() *ListOfGames {
	result := ListOfGames{Games: []map[string]string{}} //an array of type games with value of type arrays of strings
	for _, game := range s.Games {
		row := make(map[string]string)
		row["id"] = game.ID
		row["title"] = game.Title
		row["players"] = fmt.Sprint(len(game.Players), "/", game.MaxPlayers())
		result.Games = append(result.Games, row) //appends each games here
	}
	return &result //pointer to the result of the games
}

//CreateMessageListOfGames creates a list of games
func (s *Server) CreateMessageListOfGames() *MessageListOfGames {
	list := s.CreateListOfGames()
	return CreateMessageListOfGames(list)

}
