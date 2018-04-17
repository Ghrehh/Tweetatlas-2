package utils

import (
	"log"
	"os"
	"io/ioutil"
	"encoding/json"
)

type StreamParams struct {
	Filter []string `json:"filter"`
}

func parseStreamParams(data []byte) StreamParams {
	params := StreamParams{}
	err := json.Unmarshal(data, &params)

	if err != nil {
		log.Print("Error parsing 'config/stream_params.json'")
		log.Print(err.Error())
		os.Exit(1)
	}

	return params
}

func GetStreamFilter() []string {
	// Attempt to get a search phrase from an environment variable
	// Env var is passed as a JSON string
	var search_phrases []string
	search_phrases_json := os.Getenv("SEARCH_PHRASES")
	err := json.Unmarshal([]byte(search_phrases_json), &search_phrases)

	if search_phrases_json != "" && err == nil {
		return search_phrases
	}

	// Attempt to get a search phrase from config file
	data, err := ioutil.ReadFile("config/stream_params.json")

	if err != nil {
		log.Print("Error opening 'config/stream_params.json'")
		log.Print(err.Error())
		os.Exit(1)
	}

	params := parseStreamParams(data)

	return params.Filter
}
