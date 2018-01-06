package web

import (
	"testing"
)

func TestLocationAggregate(t *testing.T) {
	la := NewLocationAggregate()

	la.AddParsedLocation("foo")
	la.AddParsedLocation("foo")
	la.AddParsedLocation("bar")
	la.AddParsedLocation("")

	json := la.ToJSON()
	expectedJson := `{"bar":1,"foo":2,"unknown":1}`

	if string(json) != expectedJson {
		t.Error("did not receive expected json output")
	}
}
