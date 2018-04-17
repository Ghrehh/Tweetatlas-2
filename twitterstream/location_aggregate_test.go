package twitterstream

import (
	"testing"
)

func TestLocationAggregate(t *testing.T) {
	la := NewLocationAggregate([]string{"bar","foo"})

	la.AddParsedLocation("foo")
	la.AddParsedLocation("foo")
	la.AddParsedLocation("bar")
	la.AddParsedLocation("")

	json := la.ToJSON()
	expectedJson := `{"location_data":{"bar":1,"foo":2,"unknown":1},"search_phrases":["bar","foo"]}`

	if string(json) != expectedJson {
		t.Error("expected " + string(json) + " to equal " + expectedJson)
	}
}
