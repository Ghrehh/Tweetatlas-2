package web

import (
	"net/http"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Connection struct {
	orchestrator *ConnectionOrchestrator
	websocketConnection *websocket.Conn
	tweetStream chan *twitter.Tweet
}

func newConnection(co *ConnectionOrchestrator, conn *websocket.Conn) *Connection {
	return &Connection{
		orchestrator: co,
		websocketConnection: conn,
		tweetStream: make(chan *twitter.Tweet),
	}
}

func (c *Connection) writePump() {
	defer c.websocketConnection.Close()

	for {
		tweet := <- c.tweetStream
		err := c.websocketConnection.WriteJSON(tweet)

		if err != nil {
			log.Print("Error sending message to client")
			log.Print(err)
			break
		}
	}
}

func ServeWebsocket(co *ConnectionOrchestrator, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	connection := newConnection(co, conn)
	co.add <- connection

	go connection.writePump()
}
