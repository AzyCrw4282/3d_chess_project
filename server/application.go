package main

import (
	"flag"
	"github.com/team142/angrychess/controller"
	"github.com/team142/angrychess/io/ws"
	"github.com/team142/angrychess/model"
	"log"
)

//sets flag but for cmd line entry but not used parsed in this case.
var addr = flag.String("addr", "127.0.0.1:8000", "http service address")
var publicAddress = flag.String("public", "local", "Server public domain name")
var cantStartBeforeFull = flag.String("emptySeat", "true", "Server public domain name")

func main() {

	log.Println("Oh hai 🚀 Lets Go 🎠")
	server := model.CreateServer(*publicAddress, controller.HandleIncoming, *cantStartBeforeFull == "true")

	//Blocking call
	ws.StartWSServer(addr, server.HandleMessage)

}
