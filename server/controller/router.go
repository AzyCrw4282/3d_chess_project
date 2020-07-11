package controller

import (
	"encoding/json"
	"fmt"
	"github.com/team142/angrychess/io/ws"
	"github.com/team142/angrychess/model"
	"log"
)

const (
	inMessageNick        = "nick" //nickname - to set it in the game
	inMessageStartGame   = "start-game"
	inMessageCreateGame  = "create-game"
	inMessageJoinGame    = "join-game"
	inMessageListOfGames = "list-games"
	inMessageMove        = "move"
	inMessageChangeSeat  = "seat"
	inMessageDC          = "disconnect"
)

//HandleIncoming handles messages arriving to the websocket server. thses are then handled in `in-message.go` file
func HandleIncoming(server *model.Server, client *ws.Client, msg *[]byte) {
	var message model.BaseMessage
	if err := json.Unmarshal(*msg, &message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
	}

	if message.Msg == inMessageNick {
		handleInMessageNick(server, client, msg)

	} else if message.Msg == inMessageCreateGame {
		handleInMessageCreateGame(server, client)

	} else if message.Msg == inMessageJoinGame {
		handleInMessageJoinGame(server, client, msg)

	} else if message.Msg == inMessageStartGame {
		handleInMessageStartGame(server, client)

	} else if message.Msg == inMessageListOfGames {
		handleInMessageListOfGame(server, client)

	} else if message.Msg == inMessageMove {
		handleInMessageMove(server, client, msg)

	} else if message.Msg == inMessageChangeSeat {
		handleInMessageChangeSeat(server, client, msg)

	} else if message.Msg == inMessageDC {
		handleInMessageDC(server, client)

	} else {
		log.Println("Unknown route: ", message.Msg)
		log.Println(string(*msg))
	}

}
