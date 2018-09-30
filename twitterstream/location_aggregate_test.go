package twitterstream

import (
	"testing"
	"encoding/json"
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
)

func TestLocationAggregate(t *testing.T) {
	la := NewLocationAggregate([]string{"bar","foo"})
	tweet := &twitter.Tweet{}

	la.AddParsedLocation("foo")
	la.AddParsedLocation("foo")
	la.AddParsedLocation("bar")
	la.AddParsedLocation("")
	la.AddSampleTweet(tweet, "baz")

	locationAggregateJSON := la.ToJSON()
	
	tweetJSON, _ := json.Marshal(tweet)
	expectedJson := fmt.Sprintf(`{"location_data":{"bar":1,"foo":2,"unknown":1},"sample_tweet":{"parsed_location":"baz","data":%v},"search_phrases":["bar","foo"],"search_phrase_index":0}`, string(tweetJSON))

	if string(locationAggregateJSON) != expectedJson {
		t.Error("expected " + string(locationAggregateJSON) + " to equal " + expectedJson)
	}
}

func TestReset(t *testing.T) {
	la := NewLocationAggregate([]string{"bar", "foo"})
	la.AddParsedLocation("foo")

	if la.Data["foo"] != 1 {
		t.Error("Expected foo to be 1")
	}

	if la.SearchPhraseIndex != 0 {
		t.Error("Expected search phrase index to be 0")
	}

	la.Reset(1)

	if la.Data["foo"] != 0 {
		t.Error("Expected foo to be nil")
	}

	if la.SearchPhraseIndex != 1 {
		t.Error("Expected search phrase index to be 1")
	}
}
