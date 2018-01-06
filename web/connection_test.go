package web

import (
	"testing"
)

type fakeUpgrader interface {}

func (u *fakeUpgrader) Upgrade(http.ResponseWriter, *http.Request, http.Header) (*websocket.Conn, error) {
	return &websocket.Conn{}, nil
}

func TestServeWebsocket(t *testing.T) {
	connection := ServeWeboscket(
	)

	
	if {
	}
}
