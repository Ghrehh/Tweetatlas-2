package utils

import (
	"log"
	"os"
	"io/ioutil"
	"encoding/json"
)

type OauthKeys struct {
	ConsumerKey string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
	AccessToken string `json:"access_token"`
	AccessSecret string `json:"access_secret"`
}

func parseOauthKeys(data []byte) OauthKeys {
	keys := OauthKeys{}
	err := json.Unmarshal(data, &keys)

	if err != nil {
		log.Print("Error parsing 'config/twitter_keys.json'")
		log.Fatal(err.Error())
	}

	return keys
}

func GetOauthKeys() OauthKeys {
	data, err := ioutil.ReadFile("config/twitter_keys.json")

	if err == nil {
		return parseOauthKeys(data)
	}

	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")

	if consumerKey == "" {
		log.Fatal("TWITTER_CONSUMER_KEY env variable not set")
	}

	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")

	if consumerSecret == "" {
		log.Fatal("TWITTER_CONSUMER_SECRET env variable not set")
	}

	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")

	if accessToken == "" {
		log.Fatal("TWITTER_ACCESS_TOKEN env variable not set")
	}

	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	if accessSecret == "" {
		log.Fatal("TWITTER_ACCESS_SECRET env variable not set")
	}

	return OauthKeys{
		ConsumerKey: consumerKey,
		ConsumerSecret: consumerSecret,
		AccessToken: accessToken,
		AccessSecret: accessSecret,
	}
}
