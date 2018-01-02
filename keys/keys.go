package keys

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
)

type OauthKeys struct {
	ConsumerKey string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
	AccessToken string `json:"access_token"`
	AccessSecret string `json:"access_secret"`
}

func Load() []byte {
	data, err := ioutil.ReadFile("./twitter_keys.json")

	if err != nil {
		fmt.Println("Error loading 'twitter_keys.json'")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return data
}

func Parse(data []byte) OauthKeys {
	keys := OauthKeys{}
	err := json.Unmarshal(data, &keys)

	if err != nil {
		fmt.Println("Error parsing 'twitter_keys.json'")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return keys
}
