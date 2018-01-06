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
	data, err := ioutil.ReadFile("config/stream_params.json")

	if err != nil {
		log.Print("Error opening 'config/stream_params.json'")
		log.Print(err.Error())
		os.Exit(1)
	}

	params := parseStreamParams(data)

	return params.Filter
}
