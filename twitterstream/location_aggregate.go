package twitterstream

import (
	"encoding/json"
	"sync"
	"log"

	"github.com/dghubble/go-twitter/twitter"
)

type LocationAggregater interface {
	AddParsedLocation(string)
	ToJSON() []byte
	AddSampleTweet(*twitter.Tweet)
}

type LocationAggregate struct {
	Data map[string]int `json:"location_data"`
	SampleTweet *twitter.Tweet `json:"sample_tweet"`
	SearchPhrases []string `json:"search_phrases"`
	mutex *sync.Mutex
}

func NewLocationAggregate(searhPhrases []string) *LocationAggregate {
	return &LocationAggregate{
		Data: make(map[string]int),
		SearchPhrases: searhPhrases,
		mutex: &sync.Mutex{},
	}
}

func (la *LocationAggregate) AddParsedLocation(location string) {
	if location == "" {
		location = "unknown"
	}

	la.mutex.Lock()
	la.Data[location] +=1
	la.mutex.Unlock()
}

func (la *LocationAggregate) AddSampleTweet(tweet *twitter.Tweet) {
	la.mutex.Lock()
	la.SampleTweet = tweet
	la.mutex.Unlock()
}

func (la *LocationAggregate) ToJSON() []byte {
	la.mutex.Lock()
	json, err := json.Marshal(la)
	la.mutex.Unlock()

	if err != nil {
		log.Print("Error JSONifying the LocationAggregate")
		log.Fatal(err.Error())
	}

	return json
}
