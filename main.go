package main

import (
	"net/http"
	"log"

	"github.com/ghrehh/tweetatlas/twitterstream"
	"github.com/ghrehh/tweetatlas/utils"
	"github.com/ghrehh/tweetatlas/web"

	"github.com/ghrehh/findlocation"
	"github.com/dghubble/go-twitter/twitter"
  "github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func main() {
	log.Print("App Started")

	// Get the search phrase(s)
	filter := utils.GetStreamFilter()

	// Location aggregate stores our location results
	locationAggregate := twitterstream.NewLocationAggregate(filter)

	// Location finder is a package that attempts to find a Twitter user's location.
	locationFinder := findlocation.NewLocationFinder()

	// Connection orchestartor manages the websocket connections
	co := web.NewConnectionOrchestrator()

	// Get twitter keys
	twitterKeys := utils.GetOauthKeys()

	// Create tweet stream
	streamer := twitterstream.NewTweetStream(twitterKeys)
	stream, err := twitterstream.FilterStream(&streamer, filter)

	if err != nil {
		log.Print(err)
	}

	// Handle different types returned by the stream
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		location := locationFinder.FindLocation(tweet.User.Location)
		locationAggregate.AddParsedLocation(location)
		co.LaStream <- locationAggregate
	}

	// Log any message in the stream other than a tweet
	demux.Other = func(message interface{}) {
		log.Print(message)
	}

	// Upgrade to a websocket connection, allow all origins
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// Routing
	r := mux.NewRouter()
	r.Headers("Content-Type", "application/json")
	r.HandleFunc("/tweets", func(w http.ResponseWriter, r *http.Request) {
		web.ServeWebsocket(co, w, r, &upgrader)
	})

	// Start the stream with the demux
	go demux.HandleChan(stream.Messages)

	// Start the connection orchestrator
	go co.Run()

	// Serve routes with CORs handler
	http.ListenAndServe(utils.GetPort(), r)
}
