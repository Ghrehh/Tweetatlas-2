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
	AddSampleTweet(*twitter.Tweet, string)
	Reset(int)
}

type LocationAggregate struct {
	Data map[string]int `json:"location_data"`
	SampleTweet Tweet `json:"sample_tweet"`
	SearchPhrases []string `json:"search_phrases"`
	SearchPhraseIndex int `json:"search_phrase_index"`
	mutex *sync.Mutex
}

type Tweet struct {
	ParsedLocation string `json:"parsed_location"`
	Data *twitter.Tweet `json:"data"`
}

func NewLocationAggregate(searchPhrases []string) *LocationAggregate {
	return &LocationAggregate{
		Data: make(map[string]int),
		SearchPhrases: searchPhrases,
		mutex: &sync.Mutex{},
	}
}

func (la *LocationAggregate) Reset(searchPhraseIndex int) {
	la.mutex.Lock()
	la.Data = make(map[string]int)
	la.SearchPhraseIndex = searchPhraseIndex
	la.mutex.Unlock()
}

func (la *LocationAggregate) AddParsedLocation(location string) {
	if location == "" {
		location = "unknown"
	}

	la.mutex.Lock()
	la.Data[location] +=1
	la.mutex.Unlock()
}

func (la *LocationAggregate) AddSampleTweet(t *twitter.Tweet, location string) {
	la.mutex.Lock()
	la.SampleTweet = Tweet{ParsedLocation: location, Data: t}
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
