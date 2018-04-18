package web

import (
	"log"

	"github.com/ghrehh/tweetatlas/twitterstream"
)

type ConnectionOrchestrator struct {
	connections map[*Connection]bool
	LaStream chan twitterstream.LocationAggregater
	add chan *Connection
	remove chan *Connection
}

func NewConnectionOrchestrator() *ConnectionOrchestrator {
	return &ConnectionOrchestrator{
		connections: make(map[*Connection]bool),
		LaStream: make(chan twitterstream.LocationAggregater),
		add: make(chan *Connection),
		remove: make(chan *Connection),
	}
}

func (co *ConnectionOrchestrator) Run() {
	for {
		select {
		case connection := <- co.add:
			co.connections[connection] = true
		case connection := <- co.remove:
			if _, ok := co.connections[connection]; ok {
				delete(co.connections, connection)
				close(connection.dataStream)
				log.Print("closed connection")
			}
		case la := <- co.LaStream:
			laJSON := la.ToJSON()

			for connection := range co.connections {
				connection.dataStream <- &laJSON
			}
		}
	}
}
