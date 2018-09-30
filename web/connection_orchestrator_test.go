package web

import (
	"testing"
	"time"

	"github.com/dghubble/go-twitter/twitter"
)

type fakeWebsocketConnection struct{
}

func (f fakeWebsocketConnection) Close() error {
	return nil
}

func (f fakeWebsocketConnection) WriteMessage(int, []byte) error {
	return nil
}

type fakeLocationAggregate struct{
}

func (f fakeLocationAggregate) AddParsedLocation(string) {
}

func (f fakeLocationAggregate) AddSampleTweet(*twitter.Tweet, string) {
}

func (f fakeLocationAggregate) ToJSON() []byte {
	return []byte("foo")
}

func (f fakeLocationAggregate) Reset(searchPhrases []string) {
}

func TestConnectionOrchestratorRun(t *testing.T) {
	la := fakeLocationAggregate{}
	co := NewConnectionOrchestrator()
	c := newConnection(co, fakeWebsocketConnection{})

	go co.Run()

	co.add <- c

	if co.connections[c] != true {
		t.Error("expected connection to be added to connections map")
	}

	co.LaStream <- la
	connectionData := <- c.dataStream

	if string(*connectionData) != "foo" {
		t.Error("expected connection to receive correct byte array")
	}

	co.remove <- c
	time.Sleep(time.Millisecond)

	if co.connections[c] != false {
		t.Error("expected connection to have been removed from the connections map")
	}
}
