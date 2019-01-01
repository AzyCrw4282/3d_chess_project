package controller

import (
	"encoding/json"
	"fmt"
	"github.com/team142/chessfor4/io/ws"
	"github.com/team142/chessfor4/model/messages"
	"log"
)

func HandleIncoming(client *ws.Client, msg []byte) {
	var message messages.BaseMessage
	if err := json.Unmarshal(msg, &message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
	}

	if message.Msg == messages.InMessageNick {
		HandleTermSearch(client, msg)
	} else {
		log.Println("Unknown route: ", message.Message)
	}

}
