package ws

import (
	"bytes"
	"flag"
	"github.com/fasthttp/websocket"
	"log"
	"net/http"
	"time"
)

//StartWSServer starts a  websocket server- > *handler* is what enable client/server communication via sockets
func StartWSServer(addr *string, handler func(*Client, *[]byte)) {

	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, handler, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveWs(hub *Hub, handler func(*Client, *[]byte), w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//handler is of type HandleMessage function from class Server TODO:check the exact operation here!!!
	client := &Client{Hub: hub, conn: conn, CanSend: true, send: make(chan []byte, 256), handler: handler}
	client.Hub.register <- client

	//separate go-routine for reading and writing
	go client.writePump()
	go client.readPump()
}

/*
A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately,
but the function call is not executed until the surrounding function returns.
*/

func (c *Client) readPump() {
	defer func() { //won't run until the main function has returned
		c.Hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.handleMessage(&message)
	}
}

//a simple method-> uses a ticker to keep connection alive and  c.send() to actively communicate.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for { // runs in a continious loop to write feed of msgs
		select { //runs any channels as long as they can send a value to it
		case message, ok := <-c.send: // channell sends in the data
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The Hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
