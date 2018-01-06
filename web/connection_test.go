package web

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"errors"

	"github.com/gorilla/websocket"
)

type fakeUpgrader struct {}

func (u *fakeUpgrader) Upgrade(http.ResponseWriter, *http.Request, http.Header) (*websocket.Conn, error) {
	return &websocket.Conn{}, nil
}

type fakeUpgraderError struct {}

func (u *fakeUpgraderError) Upgrade(http.ResponseWriter, *http.Request, http.Header) (*websocket.Conn, error) {
	return &websocket.Conn{}, errors.New("foo")
}

func TestServeWebsocket(t *testing.T) {
	co := NewConnectionOrchestrator()

	connection, err := ServeWebsocket(
		co,
		httptest.NewRecorder(),
		&http.Request{},
		&fakeUpgrader{},
	)

	if err != nil {
		t.Error("error should not be nil")
	}

	coAddChannelConnection := <- co.add

	if coAddChannelConnection != connection {
		t.Error("expected connection to be pushed to the co add channel")
	}
}

func TestServeWebsocketError(t *testing.T) {
	co := NewConnectionOrchestrator()

	_, err := ServeWebsocket(
		co,
		httptest.NewRecorder(),
		&http.Request{},
		&fakeUpgraderError{},
	)

	if err == nil {
		t.Error("error should have been raised")
	}
}
