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
	la.AddSampleTweet(tweet)

	locationAggregateJSON := la.ToJSON()
	
	tweetJSON, _ := json.Marshal(tweet)
	expectedJson := fmt.Sprintf(`{"location_data":{"bar":1,"foo":2,"unknown":1},"sample_tweet":%v,"search_phrases":["bar","foo"]}`, string(tweetJSON))

	if string(locationAggregateJSON) != expectedJson {
		t.Error("expected " + string(locationAggregateJSON) + " to equal " + expectedJson)
	}
}
