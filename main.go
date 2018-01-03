package main

import (
	"net/http"
	"os"
	"log"
	"strings"

	"github.com/ghrehh/tweetatlas/twitterstream"
	"github.com/ghrehh/tweetatlas/keys"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/gorilla/websocket"
	"github.com/gorilla/handlers"
  "github.com/gorilla/mux"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var RecentTweet *twitter.Tweet

func tweets(w http.ResponseWriter, r *http.Request) {
	var lastTweetId int64 = 0
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	for {
		if lastTweetId != RecentTweet.ID {
			lastTweetId = RecentTweet.ID

			err = c.WriteJSON(RecentTweet)

			if err != nil {
				log.Print("write:", err)
				break
			}
		}
	}

	defer c.Close()
}

func main() {
	twitterKeys := keys.Parse(keys.Load())
	streamer := twitterstream.NewTweetStreamer(twitterKeys)
	filter := []string{"happy"}
	stream, _ := twitterstream.FilterStream(&streamer, filter)

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		RecentTweet = tweet
	}

	go demux.HandleChan(stream.Messages)

	port := os.Getenv("PORT")
	r := mux.NewRouter()
	r.Headers("Content-Type", "application/json")
	r.HandleFunc("/tweets", tweets)

	if port != "" {
		portStrings := []string{":", port}
		http.ListenAndServe(strings.Join(portStrings, ""), handlers.CORS()(r))
	} else {
		http.ListenAndServe(":5555", handlers.CORS()(r))
	}
	
}
