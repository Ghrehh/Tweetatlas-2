package web

import (
	"github.com/dghubble/go-twitter/twitter"
)

type ConnectionOrchestrator struct {
	connections map[*Connection]bool
	tweetStream chan *twitter.Tweet
	add chan *Connection
	remove chan *Connection
}

func NewConnectionOrchestrator(tweetStream chan *twitter.Tweet) *ConnectionOrchestrator {
	return &ConnectionOrchestrator{
		connections: make(map[*Connection]bool),
		tweetStream: tweetStream,
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
				close(connection.tweetStream)
			}
		case tweet := <-co.tweetStream:
			for connection := range co.connections {
				select {
				case connection.tweetStream <- tweet:
				default:
					close(connection.tweetStream)
					delete(co.connections, connection)
				}
			}
		}
	}
}
