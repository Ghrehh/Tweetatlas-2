package utils

import (
	"fmt"
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
		fmt.Println("Error parsing 'config/twitter_keys.json'")
		fmt.Println(err.Error())
		os.Exit(1)
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
		fmt.Println("TWITTER_CONSUMER_KEY env variable not set")
		os.Exit(1)
	}

	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")

	if consumerSecret == "" {
		fmt.Println("TWITTER_CONSUMER_SECRET env variable not set")
		os.Exit(1)
	}

	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")

	if accessToken == "" {
		fmt.Println("TWITTER_ACCESS_TOKEN env variable not set")
		os.Exit(1)
	}

	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	if accessSecret == "" {
		fmt.Println("TWITTER_ACCESS_SECRET env variable not set")
		os.Exit(1)
	}

	return OauthKeys{
		ConsumerKey: consumerKey,
		ConsumerSecret: consumerSecret,
		AccessToken: accessToken,
		AccessSecret: accessSecret,
	}
}
