package ws

import (
	"encoding/json"
	"fmt"
	"github.com/fasthttp/websocket"
	"log"
	"net/http"
	"time"
)

/*
Keeps track of all clients using a hub
*/

//defines const values
const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 1024
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//Hub keeps track of all connected clients
//acts as the wrapper for all internal connectons made
type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		register:   make(chan *Client), // make channel of type client-> used for client communication
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

//Client describes a websocket client
type Client struct {
	Hub     *Hub
	conn    *websocket.Conn
	send    chan []byte
	CanSend bool
	handler func(*Client, *[]byte) // anonymous function as a field for the struct
}

func (c *Client) handleMessage(msg *[]byte) {
	go c.handler(c, msg) //TODO: is this right?
}

//msg sent to the channel handling array of bytes
func (c *Client) sendBytes(msg []byte) {
	if c.CanSend {
		c.send <- msg
	}
}

func (c *Client) SendObject(o interface{}) {
	b, err := json.Marshal(o)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.sendBytes(b)
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				client.CanSend = false
				delete(h.clients, client)
				close(client.send)

				data := struct {
					Msg string `json:"msg"`
				}{
					"disconnect",
				}
				b, _ := json.Marshal(data)
				go func() {
					//TODO: Might be the problem
					time.Sleep(1 * time.Second)
					client.handler(client, &b)
				}()
				log.Println("Client unregistered")
			}
		case client := <-h.register:
			h.clients[client] = true
			log.Println("New client registered")
		}
	}
}
