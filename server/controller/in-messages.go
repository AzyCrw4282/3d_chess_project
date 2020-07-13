package controller

import (
	"encoding/json"
	"fmt"
	"github.com/team142/angrychess/io/ws"
	"github.com/team142/angrychess/model"
	"log"
)

/*
handles all incoming msgs and passes the msg to the model.
All cmds to the next layer are made here. Initially originating from the `view` point.
Most methods  are straightforward.
*/

func handleInMessageStartGame(server *model.Server, client *ws.Client) {
	startGameByClient(server, client)
}

//sets nickname. Method decodes from json and checks for errors before making the call
func handleInMessageNick(server *model.Server, client *ws.Client, msg *[]byte) {
	var message model.MessageNick
	if err := json.Unmarshal(*msg, &message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
	}
	setNick(server, client, message.Nick)
}

// unmarshalls the given string and calls move option
func handleInMessageMove(server *model.Server, client *ws.Client, msg *[]byte) {
	message := &model.MessageMove{}
	if err := json.Unmarshal(*msg, message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
		return
	}
	move(server, message, client)
}

/*
--> a go routine to handle creating a list of games. Use of routine is for the expensive call, dependingin
on the games.
*/
func handleInMessageListOfGame(server *model.Server, client *ws.Client) {
	go func() {
		reply := server.CreateMessageListOfGames()
		log.Println(">> Sending list of games ")
		client.SendObject(reply)
	}()
}

/*
All methods follow the same applied logic in its each case
*/
func handleInMessageJoinGame(server *model.Server, client *ws.Client, msg *[]byte) {
	var message model.MessageJoinGame
	if err := json.Unmarshal(*msg, &message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
	}
	joinGameByClient(server, message.ID, server.Lobby[client])
	notifyLobby(server)
}

func handleInMessageDC(server *model.Server, client *ws.Client) {
	disconnect(server, client)
}

func handleInMessageCreateGame(server *model.Server, client *ws.Client) {
	createGameByClient(server, client)
	notifyLobby(server)
}

func handleInMessageChangeSeat(server *model.Server, client *ws.Client, msg *[]byte) {
	var message model.MessageChangeSeat
	if err := json.Unmarshal(*msg, &message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
		return
	}
	changeSeatByClient(server, client, message.Seat)
}
