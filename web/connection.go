package web

import (
	"net/http"
	"errors"
	"log"

	"github.com/gorilla/websocket"
)


type Upgrader interface {
	Upgrade(http.ResponseWriter, *http.Request, http.Header) (*websocket.Conn, error)
}

type websocketConnection interface {
	Close() error
	WriteMessage(int, []byte) error
}

type Connection struct {
	orchestrator *ConnectionOrchestrator
	websocketConnection websocketConnection
	dataStream chan *[]byte
}

func newConnection(co *ConnectionOrchestrator, conn websocketConnection) *Connection {
	return &Connection{
		orchestrator: co,
		websocketConnection: conn,
		dataStream: make(chan *[]byte),
	}
}

func (c *Connection) writePump() {
	defer c.websocketConnection.Close()

	for {
		data := <- c.dataStream
		err := c.websocketConnection.WriteMessage(1, *data)

		if err != nil {
			log.Print("Error sending message to client")
			log.Print(err)
			break
		}
	}
}

func ServeWebsocket(co *ConnectionOrchestrator, w http.ResponseWriter, r *http.Request, u Upgrader) (*Connection, error) {
	conn, err := u.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return &Connection{}, errors.New("error opening websocket connection")
	}

	connection := newConnection(co, conn)
	co.add <- connection

	go connection.writePump()

	return connection, nil
}
