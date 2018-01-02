package main

import (
	"github.com/ghrehh/tweetatlas/twitterstream"
	"github.com/ghrehh/tweetatlas/keys"
	"fmt"
)

func main() {
	twitterKeys := keys.Parse(keys.Load())
	streamer := twitterstream.NewTweetStreamer(twitterKeys)
	filter := []string{"happy"}
	stream, _ := twitterstream.FilterStream(&streamer, filter)

	for {
		tweet := <- stream.Messages
		fmt.Printf("%+v\n", tweet)
	}
}
