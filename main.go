package main

import (
	"net/http"
	"fmt"

	"github.com/ghrehh/tweetatlas/twitterstream"
	"github.com/ghrehh/tweetatlas/keys"
	"github.com/ghrehh/tweetatlas/utils"
	"github.com/ghrehh/tweetatlas/web"

	"github.com/ghrehh/findlocation"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/gorilla/handlers"
  "github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("Tweetatlas Running")

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	locationAggregate := web.NewLocationAggregate()
	locationFinder := findlocation.NewLocationFinder()
	co := web.NewConnectionOrchestrator()

	twitterKeys := keys.GetOauthKeys()
	streamer := twitterstream.NewTweetStream(twitterKeys)
	filter := []string{"happy"}
	stream, err := twitterstream.FilterStream(&streamer, filter)

	if err != nil {
		fmt.Println(err)
	}

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		location := locationFinder.FindLocation(tweet.User.Location)
		locationAggregate.AddParsedLocation(location)
		co.LaStream <- locationAggregate
	}

	demux.Other = func(message interface{}) {
		fmt.Println(message)
	}

	go demux.HandleChan(stream.Messages)
	go co.Run()

	r := mux.NewRouter()
	r.Headers("Content-Type", "application/json")
	r.HandleFunc("/tweets", func(w http.ResponseWriter, r *http.Request) {
		web.ServeWebsocket(co, w, r, &upgrader)
	})

	http.ListenAndServe(utils.GetPort(), handlers.CORS()(r))
}
