package web

import (
	"encoding/json"
	"fmt"
	"os"
)

type LocationAggregater interface {
	AddParsedLocation(string)
	ToJSON() []byte
}

type LocationAggregate struct {
	Data map[string]int
}

func NewLocationAggregate() *LocationAggregate {
	return &LocationAggregate{
		Data: make(map[string]int),
	}
}

func (la *LocationAggregate) AddParsedLocation(location string) {
	if location == "" {
		location = "unknown"
	}

	la.Data[location] +=1
}

func (la *LocationAggregate) ToJSON() []byte {
	jsonString, err := json.Marshal(la.Data)

	if err != nil {
		fmt.Println("Error JSONifying the LocationAggregate")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return jsonString
}
