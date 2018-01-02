package main

import (
	"github.com/ghrehh/tweetatlas/twitterstream"
	"github.com/ghrehh/tweetatlas/keys"
	"fmt"
)

func main() {
	twitterKeys := keys.Parse(keys.Load())
	config, token := twitterstream.NewOauthConfig(twitterKeys)
	twitterClient := twitterstream.NewTwitterClient(config, token)
	streamer := twitterstream.TweetStreamer{twitterClient}
	filter := []string{"happy"}
	stream, _ := twitterstream.FilterStream(&streamer, filter)

	for {
		tweet := <- stream.Messages
		fmt.Printf("%+v\n", tweet)
	}
}
